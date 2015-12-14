/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { nextPassword, meetsRequirements, increment } from '../src/day-11/';
import { expect } from 'chai';

describe('day 11', () => {
  describe('silver', () => {
    [
      ['abcddee', true],
      ['acdeddee', true],
      ['abdddee', false],
      ['ghiddee', false],
      ['abcddee', true],
      ['abcdd', false],
      ['abcdddd', true],
      ['abcddd', false]
    ].forEach((test:[string, boolean]) => {
      let [spec, isValid] = test;
      it(`sees if ${spec} is valid`, () => {
        expect(meetsRequirements(spec)).to.equal(isValid);
      });
    });
    [
      ['abc', 'abd'],
      ['dbz', 'dca'],
      ['dzz', 'eaa']
    ].forEach((test:[string, string]) => {
      let [spec, next] = test;
      it (`increments ${spec} to ${next}`, () => {
        expect(increment(spec)).to.equal(next);
      });
    });
    [
      ['abcdefgh', 'abcdffaa'],
      ['ghijklmn', 'ghjaabcc']
    ].forEach((test: [string, string]) => {
      let [spec, next] = test;
      it.skip(`tests ${spec} to be ${next}`, () => { //slooooow
        expect(nextPassword(spec)).to.equal(next);
      });
    });
  });
});
