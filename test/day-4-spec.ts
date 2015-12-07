/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { adventCoin } from '../src/day-4/';
import { expect } from 'chai';

describe('day 4', () => {
  describe.skip('silver', () => { //these tests are too slow to always run!
    [
      ['abcdef', 609043],
      ['pqrstuv', 1048970],
    ].forEach((test: [string, number]) => {
      let [prefix, answer] = test;
      it(`tests ${prefix}`, () => {
        expect(adventCoin(prefix)).to.equal(answer);
      });
    });
  });
});
