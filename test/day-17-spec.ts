/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { combinations, uniqueCombinations } from '../src/day-17/';
import { expect } from 'chai';

describe('day 17', () => {
  let spec = [20, 15, 10, 5, 5].map(e => e.toString());
  it('can identify all the different ways', () => {
    expect(combinations(spec, 25)).to.equal(4);
  });
  it('can get uniques', () => {
    expect(uniqueCombinations(spec, 25)).to.equal(3);
  });
});
