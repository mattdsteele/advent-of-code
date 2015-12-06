/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import getFeet from '../src/day-2/day-2';
import { expect } from 'chai';

describe('day 2', () => {
  [
    ['2x3x4', 58],
    ['1x1x10', 43]
  ].forEach((test: [string, number]) => {
    let [dimensions, output] = test;
    it(`tests ${dimensions}`, () => {
      expect(getFeet(dimensions)).to.equal(output);
    });
  });
});
