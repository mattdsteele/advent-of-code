/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { isNice, isNiceGold } from '../src/day-5/';
import { expect } from 'chai';

describe('day 5', () => {
  describe('silver', () => { //these tests are too slow to always run!
    [
      ['ugknbfddgicrmopn', true],
      ['aaa', true],
      ['jchzalrnumimnmhp', false],
      ['haegwjzuvuyypxyu', false],
      ['dvszwmarrgswjxmb', false]
    ].forEach((test: [string, boolean]) => {
      let [prefix, answer] = test;
      it(`tests ${prefix}`, () => {
        expect(isNice(prefix)).to.equal(answer);
      });
    });
  });
  describe('gold', () => { //these tests are too slow to always run!
    [
      ['qjhvhtzxzqqjkmpb', true],
      ['xxyxx', true],
      ['uurcxstgmygtbstg', false],
      ['ieodomkazucvgmuy', false],
      ['aaa', false],
      ['aaaa', true]
    ].forEach((test: [string, boolean]) => {
      let [prefix, answer] = test;
      it(`tests ${prefix}`, () => {
        expect(isNiceGold(prefix)).to.equal(answer);
      });
    });
  });
});
