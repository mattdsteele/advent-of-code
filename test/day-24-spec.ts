/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { equalGroups, permutations, possiblePermutations, idealQe } from '../src/day-24/';
import { expect } from 'chai';

describe('day 24', () => {
  it('can produce a permutation', () => {
    let inputs = [1,2,3,4];
    let perms = permutations(inputs, 2);
    expect(perms.length).to.be.equal(10);
  });
  it('does another permutation', () => {
    let inputs = [1,2,3];
    let perms = permutations(inputs, 2);
    expect(perms.length).to.be.equal(3);
  });
  it('yet another', () => {
    let perms = permutations([1,2,3,4,5], 2);
    expect(perms.length).to.be.equal(25);
  });
  it.skip('makes smaller groups', () => {
    let inputs = [1,2,3,4];
    let groupings = possiblePermutations(inputs);
    expect(groupings.length).to.be.equal(36);
  });
  it.skip('can make equal groupings', () => {
    let inputs = [1,2,3,4,5,7,8,9,10,11];
    let groupings = equalGroups(inputs);
    expect(groupings.length).to.be.equal(96);
  });
  it('can give the best quantum entaglement', () => {
    let inputs = [1,2,3,4,5,7,8,9,10,11];
    expect(idealQe(inputs)).to.be.equal(99);
  });
});
