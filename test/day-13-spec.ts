/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { parseLine, allNames, seatingPermutations, optimalHappiness } from '../src/day-13/';
import { expect } from 'chai';

describe('day 13', () => {
  describe('silver', () => {
    [
      ['Alice would gain 54 happiness units by sitting next to Bob.', { left: 'Alice', right: 'Bob', delta: 54}],
      ['Bob would lose 4 happiness units by sitting next to Carol.', { left: 'Bob', right: 'Carol', delta: -4}],
    ].forEach((test:[string, any]) => {
      let [spec, parsedSpec] = test;
      it(`parses ${spec}`, () => {
        let { left, right, delta } = parsedSpec;
        let parsed = parseLine(spec);
        expect(parsed.left).to.equal(left);
        expect(parsed.right).to.equal(right);
        expect(parsed.delta).to.equal(delta);
      });
    });
    [
      [[
        {left: 'Alice', right: 'Bob', delta: 1},
        {left: 'Bob', right: 'Carol', delta: 1},
        {left: 'Carol', right: 'Alice', delta: 1}
      ], ['Alice', 'Bob', 'Carol']]
    ].forEach((test:[any[], string[]]) => {
      let [parsedLines, expected] = test;
      it(`gets names right for ${expected}`, () => {
        let distinctNames = allNames(parsedLines);
        expect(distinctNames).to.have.members(expected);
      });
      it('creates the right number of permutations', () => {
        let permutations = seatingPermutations(allNames(parsedLines));
        expect(permutations.length).to.equal(6);
      });
    });
    it('tests a real scenario', () => {
      let test = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`
      let lines = test.split('\n');
      let bestHappiness = optimalHappiness(lines);
      expect(bestHappiness).to.equal(330);
    });
  });
});
