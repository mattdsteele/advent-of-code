/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { numbersSum, noRed } from '../src/day-12/';
import { expect } from 'chai';

describe('day 12', () => {
  describe('silver', () => {
    [
      ['[1,2,3]', 6],
      ['{"a":2,"b":4}', 6],
      ['[[[3]]]', 3],
      ['{"a":{"b":4},"c":-1}', 3],
      ['{"a":[-1,1]}', 0],
      ['[-1,{"a":1}]', 0],
      ['[]', 0],
      ['{}', 0]
    ].forEach((test:[string, number]) => {
      let [spec, sum] = test;
      it(`sees if ${spec} is valid`, () => {
        expect(numbersSum(spec)).to.equal(sum);
      });
    });
  });
  describe('gold', () => {
    [
      ['[1,2,3]', 6],
      ['[1,{"c":"red","b":2},3]', 4]
    ].forEach((test:[string, number]) => {
      let [spec, sum] = test;
      it(`sees if ${spec} is valid`, () => {
        expect(noRed(spec)).to.equal(sum);
      });
    });
  });
});
