/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { numChars, replaced, expand } from '../src/day-8/';
import { expect } from 'chai';

describe('day 8', () => {
  describe('silver', () => {
    [
      [`""`, 0],
      [`"abc"`, 3],
      [`"aaa\\"aaa"`, 7],
      [`"\\x27"`, 1],
      [`"\\\\x27"`, 4],
    ].forEach((test: [string, number]) => {
      let [spec, length] = test;
      it(`tests ${spec} to be ${length}`, () => {
        expect(numChars(spec)).to.equal(length);
      });
    });
  });
  describe('gold', () => {
    [
      [`""`, 6],
      [`"abc"`, 9],
      [`"aaa\\"aaa"`, 16],
      [`"\\x27"`, 11],
    ].forEach((test: [string, number]) => {
      let [spec, length] = test;
      it(`expands ${spec} to be ${length}`, () => {
        expect(expand(spec)).to.equal(length);
      });
    });
  });
});
