/// <reference path="../../typings/lodash/lodash.d.ts" />
import { flatten } from 'lodash';

let objectCosts = {
  magic: 53,
  drain: 73,
  shield: 113,
  poison: 173,
  recharge: 229
};

let spendMana = (player, manaCost) => {
  player.mana -= manaCost;
  player.manaSpent += manaCost;
};

let magicAttack = (player, boss) => {
  boss.hp = boss.hp - 4;
  spendMana(player, 53);
};

let drainAttack = (player, boss) => {
  boss.hp -= 2;
  spendMana(player, 73);
  player.hp += 2;
};

class ShieldBehavior {
  durationLeft:number;
  name:string = 'shield';

  constructor(durationLeft = 6) {
    this.durationLeft = durationLeft;
  }
  exec(player, boss) {
    player.armor = 7;
    return new ShieldBehavior(this.durationLeft - 1);
  }
}

class PoisonBehavior {
  durationLeft:number;
  name = 'poison';

  constructor(durationLeft = 6) {
    this.durationLeft = durationLeft;
  }
  exec(player, boss) {
    boss.hp -= 3;
    return new PoisonBehavior(this.durationLeft - 1);
  }
}

class RechargeBehavior {
  durationLeft:number;
  name = 'recharge';

  constructor(durationLeft = 5) {
    this.durationLeft = durationLeft;
  }
  exec(player, boss) {
    player.mana += 101;
    return new RechargeBehavior(this.durationLeft - 1);
  }
}


let shieldAttack = (player, boss) => {
  spendMana(player, 113);
  player.effects.push(new ShieldBehavior());
};

let poisonAttack = (player, boss) => {
  spendMana(player, 173);
  player.effects.push(new PoisonBehavior());
};

let rechargeAttack = (player, boss) => {
  spendMana(player, 229);
  player.effects.push(new RechargeBehavior());
};

let mappings = {
  magic: magicAttack,
  drain: drainAttack,
  shield: shieldAttack,
  poison: poisonAttack,
  recharge: rechargeAttack
};

let applyEffects = (player, boss) => {
  player.effects = player.effects
  .map(effect => effect.exec(player, boss))
  .filter(effect => effect.durationLeft > 0);
};

export let playerTurn = (player, boss, attack) => {

  let newPlayer = {
    hp: player.hp,
    armor: player.armor,
    mana: player.mana,
    effects: player.effects,
    manaSpent: player.manaSpent,
    history: player.history.concat(attack)
  };

  newPlayer.armor = 0;

  let newBoss = {
    hp: boss.hp,
    damage: boss.damage
  };

  applyEffects(newPlayer, newBoss);
  mappings[attack](newPlayer, newBoss);

  let finished = newBoss.hp < 1 || newPlayer.hp < 1;
  let winner = determineWinner(newBoss, newPlayer);

  return {
    player: newPlayer,
    boss: newBoss,
    finished,
    winner
  };
};

let determineWinner = (boss, player) => {
  if (boss.hp <= 0) {
    return 'player';
  } else if (player.hp <= 0) {
    return 'boss';
  }
}

export let bossTurn = (player, boss, newArmor = 0) => {
  let newPlayer = {
    hp: player.hp,
    armor: player.armor,
    mana: player.mana,
    effects: player.effects,
    manaSpent: player.manaSpent,
    history: player.history
  };

  newPlayer.armor = newArmor;

  let newBoss = {
    hp: boss.hp,
    damage: boss.damage
  };

  applyEffects(newPlayer, newBoss);

  if (newBoss.hp > 0) {
    newPlayer.hp -= Math.max(newBoss.damage - newPlayer.armor, 1);
  }

  let finished = newBoss.hp < 1 || newPlayer.hp < 1;
  let winner = determineWinner(newBoss, newPlayer);
  return {
    player: newPlayer,
    boss: newBoss,
    finished,
    winner
  }
};

export let options = (player):string[] => {
  let availableMana = player.mana + player.effects
    .filter(effect => effect.name === 'recharge')
    .map(effect => 101)
    .reduce((prev, current) => prev + current, 0);

  return Object.keys(objectCosts)
    .filter(element => objectCosts[element] <= availableMana)
    .filter(effect => {
      let currentEffects = player.effects
        .filter(e => e.durationLeft > 1)
        .map(e => e.name);
      return currentEffects.indexOf(effect) === -1;
    });
};

let cheapestWinningManaSpent = (states, leastManaSpent):number => {
  let finishedWithWinningPlayers = states
    .filter(state => state.finished)
    .filter(state => state.winner === 'player')
    .map(state => state.player)
    .sort((a, b) => a.manaSpent - b.manaSpent);
  if (finishedWithWinningPlayers.length > 0) {
    let cheapest = finishedWithWinningPlayers[0];
    leastManaSpent = Math.min(leastManaSpent, cheapest.manaSpent);
  }
  return leastManaSpent;
};

export let leastManaSpent = (player, boss, hardMode = false):number => {
  let leastManaSpent = 10000;

  let game = {
    player,
    boss,
    finished: false,
    winner: undefined
  };
  let states = [game];

  while (true) {
    states = flatten(states.map(round => {
      if (hardMode) {
        round.player.hp -= 1;
        if (round.player.hp <= 0) {
          round.finished = true;
          round.winner = 'boss';
          return round;
        }
      }

      let nextAttack = options(round.player);
      return nextAttack.map(attack => playerTurn(round.player, round.boss, attack));
    }));

    leastManaSpent = cheapestWinningManaSpent(states, leastManaSpent);

    states = states
      .filter(state => !state.finished)
      .filter(state => state.player.manaSpent < leastManaSpent);

    if (states.length === 0) {
      break;
    }

    states = states.map(game => bossTurn(game.player, game.boss));
    leastManaSpent = cheapestWinningManaSpent(states, leastManaSpent);

    states = states
      .filter(state => !state.finished)
      .filter(state => state.player.manaSpent < leastManaSpent);

    if (states.length === 0) {
      break;
    }
  };
  return leastManaSpent;
};
