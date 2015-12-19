/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { stepTimes, stepsWithCornersStuck } from '../src/day-18/';
import { expect } from 'chai';

describe('day 18', () => {
  describe('silver', () => {
    let input = [ '.#.#.#',
                  '...##.',
                  '#....#',
                  '..#...',
                  '#.#..#',
                  '####..' ];
    [
      [0, 15],
      [1, 11],
      [2, 8],
      [3, 4],
      [4, 4],
      [5, 4]
    ].forEach(test => {
      let [round, lightsOn] = test;
      it(`shows ${lightsOn} after ${round} rounds`, () => {
        expect(stepTimes(input, round)).to.equal(lightsOn);
      });
    });
  });
  describe('gold', () => {
    let input = [ '##.#.#',
                  '...##.',
                  '#....#',
                  '..#...',
                  '#.#..#',
                  '####.#'];
    [
      [0, 17],
      [1, 18],
      [2, 18],
      [3, 18],
      [4, 14],
      [5, 17]
    ].forEach(test => {
      let [round, lightsOn] = test;
      it(`shows ${lightsOn} after ${round} rounds`, () => {
        expect(stepsWithCornersStuck(input, round)).to.equal(lightsOn);
      });
    });
  });
});
