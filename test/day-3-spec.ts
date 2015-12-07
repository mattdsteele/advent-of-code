/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { housesDelivered, roboSanta } from '../src/day-3/';
import { expect } from 'chai';

describe('day 3', () => {
  describe('silver', () => {
    [
      ['>', 2],
      ['^>v<', 4],
      ['^v^v^v^v^v', 2]
    ].forEach((test: [string, number]) => {
      let [route, houses] = test;
      it(`tests ${route}`, () => {
        expect(housesDelivered(route)).to.equal(houses);
      });
    });
  });
  describe('gold', () => {
    [
      ['^v', 3],
      ['^>v<', 3],
      ['^v^v^v^v^v', 11]
    ].forEach((test: [string, number]) => {
      let [route, houses] = test;
      it(`tests ${route}`, () => {
        expect(roboSanta(route)).to.equal(houses);
      });
    });
  });
});
