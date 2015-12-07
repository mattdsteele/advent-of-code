/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />

import { lightHouse, lightHouseGold } from '../src/day-6/';
import { expect } from 'chai';

describe.skip('day 6', () => {
  describe('silver', () => {
    [
      [[
        'turn on 0,0 through 999,999'
      ], 1000000],
      [[
        'toggle 0,0 through 999,0'
      ], 1000],
      [[
        'turn on 0,0 through 999,999',
        'toggle 0,0 through 999,0'
      ], 999000],
      [[
        'turn on 499,499 through 500,500',
      ], 4]
    ].forEach((test: [string[], number]) => {
      let [prefix, litLights] = test;
      it(`tests ${prefix}`, () => {
        expect(lightHouse(prefix)).to.equal(litLights);
      });
    });
  });
  describe('gold', () => {
    [
      [[
        'turn on 0,0 through 999,999'
      ], 1000000],
      [[
        'turn on 0,0 through 999,999',
        'turn on 0,0 through 499,999'
      ], 1500000],
      [[
        'turn off 0,0 through 999,999',
        'turn off 0,0 through 999,999'
      ], 0],
      [[
        'toggle 0,0 through 999,999'
      ], 2000000],
    ].forEach((test: [string[], number]) => {
      let [prefix, litLights] = test;
      it(`tests ${prefix}`, () => {
        expect(lightHouseGold(prefix)).to.equal(litLights);
      });
    });
  });
});
