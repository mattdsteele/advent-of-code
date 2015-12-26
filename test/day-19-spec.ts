/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { parseRule, distinctMolecules, calibrate, stepsUntil } from '../src/day-19/';
import { expect } from 'chai';

describe('day 19', () => {
  [
    ['H => HO', 'H', 'HO'],
    ['Ha => H', 'Ha', 'H']
  ].forEach(test => {
    let [input, inLeft, inRight] = test;
    it(`can parse ${input}`, () => {
      let { left, right } = parseRule(input);
      expect(left).to.equal(inLeft);
      expect(right).to.equal(inRight);
    });
  });
  let rules = ['H => HO', 'H => OH', 'O => HH'];
  it('makes distinct molecules', () => {
    let input = 'HOH';
    let molecules = distinctMolecules(rules, input);
    expect(molecules.length).to.equal(4);
    expect(molecules).to.include.members(['HOOH', 'HOHO', 'OHOH', 'HHHH']);
  });
  it('passes another test', () => {
    expect(distinctMolecules(rules, 'HOHOHO').length).to.equal(7);
  });
  describe('multiple sizes', () => {
    let rules = ['Ha => A'];
    expect(distinctMolecules(rules, 'HoHa')).to.include.members(['HoA']);
  });
  it('does proper parsing of an input', () => {
    let input = `H => HO
H => OH
O => HH

HOHOHO`.split('\n');
    expect(calibrate(input).length).to.equal(7);
  });
  describe('gold', () => {
    let rules = `e => xxx
LO => HO
x => LO

HOHOHO`.split('\n');
    it(`can get to the proper steps`, () => {
      expect(stepsUntil(rules)).to.equal(7);
    });
  });
});
