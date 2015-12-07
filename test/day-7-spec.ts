/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />

import { makeGates, newIdentity } from '../src/day-7/';
import { expect } from 'chai';

describe('day 7', () => {
  describe('silver', () => {
    [
      [[
        '123 -> a'
      ], 123],
      [[
        '456 -> a'
      ], 456],
      [[
        '456 -> a',
        '123 -> ab',
      ], 456],
      [[
        '123 -> h',
        'NOT h -> a'
      ], 65412],
      [[
        '123 -> x',
        '456 -> y',
        'x AND y -> a'
      ], 72],
      [[
        '123 -> x',
        '456 -> y',
        'x LSHIFT 2 -> a'
      ], 492],
      [[
        '123 -> x',
        'x -> z',
        '456 -> y',
        'z LSHIFT 2 -> a'
      ], 492],
      [[
        'z LSHIFT 2 -> a',
        '123 -> x',
        'x -> z',
        '456 -> y'
      ], 492],
      [[
        '123 -> x',
        '456 -> y',
        '123 AND y -> a'
      ], 72],
    ].forEach((test: [string[], number]) => {
      let [steps, aVal] = test;
      it(`tests ${steps}`, () => {
        let gates = makeGates(steps);
        expect(gates['a']()).to.equal(aVal);
      });
    });
  });
  describe('gold', () => {
    [
      [[
        '123 -> x',
        '456 -> y',
        'x AND y -> a'
      ], 67]
    ].forEach((test: [string[], number]) => {
      let [steps, aVal] = test;
      it(`tests ${steps}`, () => {
        let gates = makeGates(steps);
        let newGates = newIdentity(gates, 'y', 455);
        expect(newGates['a']()).to.equal(aVal);
      });
    });
  });
});
