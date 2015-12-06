/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import gotoFloor from '../src/day-1/day-1';
import { expect } from 'chai';

describe('day 1', () => {
  [
    ['(())', 0],
    ['()()', 0],
    ['(((', 3],
    ['(()(()(', 3],
    ['))(((((', 3],
    ['())', -1],
    ['))(', -1],
    [')))', -3],
    [')())())', -3]
  ].forEach((test: [string, number]) => {
    let [spec, floor] = test;
    it(`tests ${spec} to go to ${floor}`, () => {
      expect(gotoFloor(spec)).to.equal(floor);
    });
  });
});
