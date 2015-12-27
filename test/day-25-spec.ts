/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { orderNumber, code } from '../src/day-25/';
import { expect } from 'chai';

describe('day 25', () => {
  [
    [1, 5, 15],
    [1, 6, 21],
    [2, 1, 2],
    [4, 2, 12],
    [5, 2, 17],
  ].forEach((test:[number, number, number]) => {
    let [row, column, result] = test;
    it(`tests ${row}x${column}`, () => {
      expect(orderNumber(row, column)).to.be.equal(result);
    });
  });

  [
    [1, 1, 20151125],
    [2, 1, 31916031],
    [5, 5, 9250759],
  ].forEach((test:[number, number, number]) => {
    let [row, column, result] = test;
    it(`gets code for ${row}x${column}`, () => {
      expect(code(row, column)).to.be.equal(result);
    });
  });
});
