/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { parseLine, matchesLady, matchingSues } from '../src/day-16/';
import { expect } from 'chai';

describe('day 15', () => {
  it('can parse a line', () => {
    let spec = 'Sue 1: goldfish: 6, trees: 9, akitas: 0';
    let item = parseLine(spec);
    expect(item.name).to.equal('1');
    expect(item.goldfish).to.equal(6);
    expect(item.trees).to.equal(9);
    expect(item.akitas).to.equal(0);
    expect(item.akitas).to.equal(0);
    expect(item.rando).to.be.undefined;
  });
  it('identifies with complete information', () => {
    let items = {
      goldfish: 6,
      trees: 9,
      akitas: 0
    };
    let sue = parseLine('Sue 1: goldfish: 6, trees: 9, akitas: 0');
    expect(matchesLady(sue, items)).to.be.true;
  });
  it('identifies with incomplete information', () => {
    let items = {
      goldfish: 6,
      trees: 9,
      akitas: 0
    };
    let sue = parseLine('Sue 1: goldfish: 6, trees: 9, cats: 2');
    expect(matchesLady(sue, items)).to.be.true;
  });
  it('excludes based on bad info', () => {
    let items = {
      goldfish: 7,
      trees: 9,
      akitas: 0
    };
    let sue = parseLine('Sue 1: goldfish: 6, trees: 9, akitas: 0');
    expect(matchesLady(sue, items)).to.be.false;
  });
  it('finds all the matching sues', () => {
    let items = {
      goldfish: 7,
      trees: 9,
      akitas: 0,
    };
    let sues = matchingSues([
      'Sue 1: goldfish: 7, trees: 9, akitas: 0',
      'Sue 2: goldfish: 6, trees: 10, akitas: 0',
      'Sue 3: cats: 9, trees: 9, akitas: 0',
    ], items);
    expect(sues).to.include.members(['1', '3']);
  });
});
