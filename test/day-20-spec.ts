/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { minimumPresents } from '../src/day-20/';
import { expect } from 'chai';

describe('day 20', () => {
  [
    [120, 6],
    [145, 8]
  ].forEach(test => {
    let [minPresents, houseNumber] = test;
    expect(minimumPresents(minPresents)).to.equal(houseNumber);
  });
});
