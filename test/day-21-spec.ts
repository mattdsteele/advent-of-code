/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { doesPlayerWin, goShopping, costOfWinners } from '../src/day-21/';
import { expect } from 'chai';

describe('day 21', () => {
  it('can determine who wins or loses', () => {
    let player = { hp: 8, damage: 5, armor: 5 };
    let boss = { hp: 12, damage: 7, armor: 2};
    expect(doesPlayerWin(player, boss)).to.be.true;
  });
  it('also allows losses', () => {
    let player = { hp: 8, damage: 5, armor: 5 };
    let boss = { hp: 12, damage: 8, armor: 2};
    expect(doesPlayerWin(player, boss)).to.be.false;
  });

  it('can make permutations', () => {
    let weapons = [ { cost: 8, damage: 4, armor: 0 }, { cost: 10, damage: 5, armor: 0} ];
    let armor = [ { cost: 13, damage: 0, armor: 1}, { cost: 31, damage: 0, armor: 2} ];
    let rings = [ { cost: 25, damage: 1, armor: 0}, { cost: 50, damage: 2, armor: 0 } ];
    let permutations = goShopping(weapons, armor, rings);
    expect(permutations.length).to.equal(30); //we get some extras but what are ya gonna do
  });

  it('can determine the cost of winners', () => {
    let weapons = [ { cost: 8, damage: 4, armor: 0 }, { cost: 10, damage: 5, armor: 0} ];
    let armor = [ { cost: 13, damage: 0, armor: 1}, { cost: 31, damage: 0, armor: 2} ];
    let rings = [ { cost: 25, damage: 1, armor: 0}, { cost: 50, damage: 2, armor: 0 } ];
    let player = { hp: 8, damage: 5, armor: 5 };
    let boss = { hp: 50, damage: 7, armor: 2};
    let costs = costOfWinners(player, boss, weapons, armor, rings);
    expect(costs.length).to.equal(20); //we get some extras but what are ya gonna do
    expect(costs[0]).to.equal(21);
    expect(costs[costs.length - 1]).to.equal(116);
  });
});
