/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { equalPerms, idealQe } from '../src/day-24/';
import { expect } from 'chai';

describe('day 24', () => {
  it('makes equal groups', () => {
    let inputs = [1,2,3,4,5,7,8,9,10,11];
    let perms = equalPerms(inputs, 3);
    console.log(perms);
    expect(perms.length).to.be.equal(10);
  });

  it('can give the best quantum entaglement', () => {
    let inputs = [1,2,3,4,5,7,8,9,10,11];
    expect(idealQe(inputs)).to.be.equal(99);
  });

  it('does gold', () => {
    let inputs = [1,2,3,4,5,7,8,9,10,11];
    expect(idealQe(inputs, 4)).to.be.equal(44);
  });
});
