/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { lookAndSay } from '../src/day-10/';
import { expect } from 'chai';

describe('day 10', () => {
  describe('silver', () => {
    [
      ['1', '11'],
      ['11', '21'],
      ['21', '1211'],
      ['1211', '111221'],
      ['111221', '312211'],
    ].forEach((test: [string, string]) => {
      let [spec, next] = test;
      it(`tests ${spec} to be ${next}`, () => {
        expect(lookAndSay(spec)).to.equal(next);
      });
    });
  });
});
