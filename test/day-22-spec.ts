/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { playerTurn, bossTurn, options, leastManaSpent } from '../src/day-22/';
import { expect } from 'chai';

let standardStart = () => {
  let startPlayer = {
    hp: 10,
    mana: 250,
    armor: 0,
    effects: [],
    manaSpent: 0,
    history: []
  };
  let startBoss = {
    hp: 13,
    damage: 8
  };
  return { startPlayer, startBoss };
};
describe('day 22', () => {
  describe('player turn', () => {
    describe('basic attacks', () => {
      it('magic missile', function() {
        let { startPlayer, startBoss } = standardStart();
        var { player, boss } = playerTurn(startPlayer, startBoss, 'magic');
        expect(player.hp).to.be.equal(10);
        expect(player.mana).to.be.equal(197);
        expect(player.armor).to.be.equal(0);
        expect(player.effects.length).to.be.equal(0);

        expect(boss.hp).to.be.equal(9);

        expect(startBoss.hp).to.be.equal(13);
      });
      it('drain', function() {
        let { startPlayer, startBoss } = standardStart();
        var { player, boss } = playerTurn(startPlayer, startBoss, 'drain');
        expect(player.hp).to.be.equal(12);
        expect(player.mana).to.be.equal(177);
        expect(player.armor).to.be.equal(0);

        expect(boss.hp).to.be.equal(11);
      });

      it('shield', () => {
        let { startPlayer, startBoss } = standardStart();
        var { player, boss } = playerTurn(startPlayer, startBoss, 'shield');
        expect(player.hp).to.be.equal(10);
        expect(player.mana).to.be.equal(137);
        expect(player.armor).to.be.equal(0);
        expect(player.effects.length).to.be.equal(1);

        let [ shieldEffect ] = player.effects;
        expect(shieldEffect.durationLeft).to.be.equal(6);
        expect(shieldEffect.exec).not.to.be.undefined;
      });

      it('poison', () => {
        let { startPlayer, startBoss } = standardStart();
        var { player, boss } = playerTurn(startPlayer, startBoss, 'poison');
        expect(player.mana).to.be.equal(77);
        expect(player.effects.length).to.be.equal(1);

        let [ effect ] = player.effects;
        expect(effect.name).to.be.equal('poison');
        expect(effect.durationLeft).to.be.equal(6);
      });

      it('recharge', () => {
        let { startPlayer, startBoss } = standardStart();
        var { player, boss } = playerTurn(startPlayer, startBoss, 'recharge');
        expect(player.effects.length).to.be.equal(1);

        let [ effect ] = player.effects;
        expect(effect.name).to.be.equal('recharge');
        expect(player.mana).to.be.equal(21);
        expect(effect.durationLeft).to.be.equal(5);
      });

      describe('effects', () => {
        it('shield', () => {
          let { startPlayer, startBoss } = standardStart();
          var { player, boss } = playerTurn(startPlayer, startBoss, 'shield');
          let [ effect ] = player.effects;
          let newEffect = effect.exec(player, boss);

          expect(effect.name).to.be.equal('shield');
          expect(effect.durationLeft).to.be.equal(6);
          expect(newEffect.durationLeft).to.be.equal(5);

          expect(player.armor).to.be.equal(7);
        });

        it('poison', () => {
          let { startPlayer, startBoss } = standardStart();
          var { player, boss } = playerTurn(startPlayer, startBoss, 'poison');
          let [ effect ] = player.effects;
          let newEffect = effect.exec(player, boss);

          expect(effect.durationLeft).to.be.equal(6);
          expect(newEffect.durationLeft).to.be.equal(5);

          expect(boss.hp).to.be.equal(10);
        });

        it('recharge', () => {
          let { startPlayer, startBoss } = standardStart();
          var { player, boss } = playerTurn(startPlayer, startBoss, 'recharge');
          let [ effect ] = player.effects;
          let currentMana = player.mana;
          let newEffect = effect.exec(player, boss);

          expect(effect.durationLeft).to.be.equal(5);
          expect(newEffect.durationLeft).to.be.equal(4);

          expect(player.mana).to.be.equal(currentMana + 101);
        });

      });
    });
  });
  describe('boss turn', () => {
    it('hits you', () => {
      let { startPlayer, startBoss } = standardStart();
      let { player, boss } = bossTurn(startPlayer, startBoss);
      expect(player.hp).to.be.equal(2);
    });
    it('honors armor', () => {
      let { startPlayer, startBoss } = standardStart();
      let { player, boss } = bossTurn(startPlayer, startBoss, 5);
      expect(player.hp).to.be.equal(7);
    });
    it('takes at least one damage', () => {
      let { startPlayer, startBoss } = standardStart();
      let { player, boss } = bossTurn(startPlayer, startBoss, 9999);
      expect(player.hp).to.be.equal(9);
    });
    it('executes durations', () => {
      let { startPlayer, startBoss } = standardStart();

      var { player, boss } = playerTurn(startPlayer, startBoss, 'shield');
      let afterBoss = bossTurn(player, boss);

      let player2 = afterBoss.player;
      let boss2 = afterBoss.boss;
      let [ effect ] = player2.effects;
      expect(effect.durationLeft).to.be.equal(5);
      expect(player2.armor).to.be.equal(7);
      expect(player2.hp).to.be.equal(9);
    });
  });
  describe('turn 2', () => {
    it('executes on player turns as well', () => {
      let { startPlayer, startBoss } = standardStart();
      var { player, boss } = playerTurn(startPlayer, startBoss, 'poison');
      var { player, boss } = bossTurn(player, boss);
      expect(boss.hp).to.be.equal(10);

      var { player, boss } = playerTurn(player, boss, 'shield');

      expect(boss.hp).to.be.equal(7);
    });

    it('removes effects when finished', () => {
      let { startPlayer, startBoss } = standardStart();
      var { player, boss } = playerTurn(startPlayer, startBoss, 'poison');
      player.effects[0].durationLeft = 1;

      var { player, boss } = bossTurn(player, boss);
      expect(boss.hp).to.be.equal(10);
      expect(player.effects.length).to.be.equal(0);

      var { player, boss } = playerTurn(player, boss, 'magic');

      expect(boss.hp).to.be.equal(6);
      expect(player.effects.length).to.be.equal(0);
    });
  });
  describe('full game', () => {
    it('plays the example game', () => {
      var { startPlayer, startBoss } = standardStart();
      var { player, boss } = playerTurn(startPlayer, startBoss, 'poison');
      var { player, boss } = bossTurn(player, boss);
      var { player, boss, finished } = playerTurn(player, boss, 'magic');
      expect(finished).to.be.false;

      var { player, boss, finished, winner } = bossTurn(player, boss);

      expect(finished).to.be.true;
      expect(winner).to.be.equal('player');
      expect(boss.hp).to.be.equal(0);
      expect(player.hp).to.be.equal(2);
      expect(player.mana).to.be.equal(24);
    });
    it('plays another game', () => {
      var { startPlayer, startBoss } = standardStart();
      startBoss.hp = 14;
      var { player, boss } = playerTurn(startPlayer, startBoss, 'recharge');
      var { player, boss } = bossTurn(player, boss);
      var { player, boss } = playerTurn(player, boss, 'shield');
      var { player, boss } = bossTurn(player, boss);
      var { player, boss } = playerTurn(player, boss, 'drain');
      var { player, boss } = bossTurn(player, boss);
      var { player, boss } = playerTurn(player, boss, 'poison');
      var { player, boss } = bossTurn(player, boss);
      var { player, boss } = playerTurn(player, boss, 'magic');
      var { player, boss, finished, winner } = bossTurn(player, boss);

      expect(finished).to.be.true;
      expect(winner).to.be.equal('player');
      expect(boss.hp).to.be.equal(-1);
      expect(player.hp).to.be.equal(1);
      expect(player.mana).to.be.equal(114);
    });

    it('plays when the boss wins', () => {
      var { startPlayer, startBoss } = standardStart();
      startBoss.hp = 24;
      var { player, boss } = playerTurn(startPlayer, startBoss, 'recharge');
      var { player, boss } = bossTurn(player, boss);
      var { player, boss } = playerTurn(player, boss, 'recharge');
      var { player, boss, finished, winner } = bossTurn(player, boss);

      expect(finished).to.be.true;
      expect(winner).to.be.equal('boss');
      expect(boss.hp).to.be.equal(24);
      expect(player.hp).to.be.equal(-6);
      expect(player.mana).to.be.equal(196);
    });
    it('keeps track of how much mana has been spent', () => {
      var { startPlayer, startBoss } = standardStart();
      var { player, boss } = playerTurn(startPlayer, startBoss, 'poison');
      var { player, boss } = bossTurn(player, boss);
      var { player, boss, finished } = playerTurn(player, boss, 'magic');
      var { player } = bossTurn(player, boss);

      expect(player.manaSpent).to.be.equal(226);
    });
  });
  describe('sequence', () => {
    it('gives all options when mana is plentiful', () => {
      var { startPlayer, startBoss } = standardStart();
      let available = options(startPlayer);
      expect(available).to.include.members(['magic', 'drain', 'shield', 'recharge', 'poison']);
    });
    it('gives fewer options when mana is limited', () => {
      var { startPlayer, startBoss } = standardStart();
      startPlayer.mana = 150;
      let available = options(startPlayer);
      expect(available).to.include.members(['magic', 'drain', 'shield']);
      expect(available).not.to.include.members(['poison', 'recharge']);
    });
    it('takes recharge into consideration', () => {
      var { startPlayer, startBoss } = standardStart();
      var { player, boss } = playerTurn(startPlayer, startBoss, 'recharge');
      player.mana = 100;
      let available = options(player);
      expect(available).to.include.members(['magic', 'drain', 'shield', 'poison']);
      expect(available).not.to.include.members(['recharge']);
    });
    it('does not let you do two poisons at a time', () => {
      var { startPlayer, startBoss } = standardStart();
      var { player, boss } = playerTurn(startPlayer, startBoss, 'poison');
      player.mana = 200;
      let available = options(player);
      expect(available).to.include.members(['magic', 'drain', 'shield']);
      expect(available).not.to.include.members(['poison']);
    });
  });
  describe.skip('actual game', () => {
    it('runs a game', () => {
      var { startPlayer, startBoss } = standardStart();
      let winningCost = leastManaSpent(startPlayer, startBoss);
      expect(winningCost).to.be.equal(226);
    });
    it('runs another game', () => {
      var { startPlayer, startBoss } = standardStart();
      startBoss.hp = 14;
      let winningCost = leastManaSpent(startPlayer, startBoss);
      expect(winningCost).to.be.equal(641);
    });

    it('plays easy mode', () => {
      var { startPlayer, startBoss } = standardStart();
      startPlayer.hp = 50;
      startPlayer.mana = 500;
      startBoss.hp = 51;
      startBoss.damage = 9;
      let winningCost = leastManaSpent(startPlayer, startBoss);
      expect(winningCost).to.be.equal(900);
    });

    it('plays hard mode', () => {
      var { startPlayer, startBoss } = standardStart();
      startPlayer.hp = 50;
      startPlayer.mana = 500;
      startBoss.hp = 51;
      startBoss.damage = 9;
      let winningCost = leastManaSpent(startPlayer, startBoss, true);
      expect(winningCost).to.be.greaterThan(900);
    });
  });
});
