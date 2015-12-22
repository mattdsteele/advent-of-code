/// <reference path="../../typings/lodash/lodash.d.ts" />
import { flatten } from 'lodash';

interface Stats {
  hp: number;
  armor: number;
  damage: number;
}

interface Equipment {
  cost: number;
  damage: number;
  armor: number;
}

let calculate = (attacker:Stats, defender:Stats):Stats => {
  let damage = Math.max(1, attacker.damage - defender.armor);
  defender.hp = defender.hp - damage;
  return defender;
};

let playRound = (attacker:Stats, defender:Stats, shouldWin:boolean):boolean => {
  defender = calculate(attacker, defender);
  if (defender.hp < 1) {
    return shouldWin;
  }
  return playRound(defender, attacker, !shouldWin);
};

export let doesPlayerWin = (player:Stats, boss:Stats):boolean => {
  return playRound(player, boss, true);
};

export let goShopping = (weapons:Equipment[], armor:Equipment[], rings:Equipment[]):Equipment[] => {
  let ringOptions:Equipment[] = [ { cost: 0, damage: 0, armor: 0 } ];
  ringOptions = ringOptions.concat(rings);
  let doubleRings = flatten(rings.map(ring => {
    let otherRings = rings.filter(r => r !== ring);
    return otherRings.map(r => {
      return {
        cost: r.cost + ring.cost,
        damage: r.damage + ring.damage,
        armor: r.armor + ring.armor
      };
    });
  }));
  ringOptions = ringOptions.concat(doubleRings);

  let weaponsAndEquipmentPerm = flatten(weapons.map(weapon => {
    let withEquipment = flatten(armor.map(equipment => {
      return {
        cost: weapon.cost + equipment.cost,
        damage: weapon.damage + equipment.damage,
        armor: weapon.armor + equipment.armor
      };
    }));
    withEquipment.push(weapon);
    return withEquipment;
  }));

  return flatten(weaponsAndEquipmentPerm.map(wae => {
    return ringOptions.map(ring => {
      return {
        cost: wae.cost + ring.cost,
        damage: wae.damage + ring.damage,
        armor: wae.armor + ring.armor
      };
    });
  }));
};

export let costOfWinners = (player:Stats, boss:Stats, weapons:Equipment[], armor:Equipment[], rings:Equipment[]):any[] => {
  return goShopping(weapons, armor, rings).map((adjustedValue:Equipment) => {
    return {
      cost: adjustedValue.cost,
      wins: doesPlayerWin(
          { hp: player.hp, armor: player.armor + adjustedValue.armor, damage: player.damage + adjustedValue.damage }, 
          { hp: boss.hp, armor: boss.armor, damage: boss.damage })
    }
  })
  .filter(result => result.wins)
  .map(result => result.cost)
  .sort((a, b) => a - b);
};
