/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { parseLine, bestScore } from '../src/day-15/';
import { expect } from 'chai';

describe('day 15', () => {
  it('can parse a line', () => {
    let spec = 'Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8';
    let { name, capacity, durability, flavor, texture, calories } = parseLine(spec);
    expect(name).to.equal('Butterscotch');
    expect(capacity).to.equal(-1);
    expect(durability).to.equal(-2);
    expect(flavor).to.equal(6);
    expect(texture).to.equal(3);
    expect(calories).to.equal(8);
  });
  it('can calculate the best score', () => {
    let spec = ['Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8',
                'Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3'];
    let best = bestScore(spec);
    expect(best).to.equal(62842880);
  });
});
