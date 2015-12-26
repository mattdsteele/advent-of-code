/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { getRegisters } from '../src/day-23/';
import { expect } from 'chai';

describe('day 23', () => {
  [
    [['inc a', 'inc a'], { a: 2, b: 0 }],
    [['inc a', 'inc a', 'inc a'], { a: 3, b: 0 }],
    [['inc a', 'inc a', 'hlf a'], { a: 1, b: 0 }],
    [['inc a', 'tpl a', 'inc a'], { a: 4, b: 0 }],
    [['inc a', 'inc b'], { a: 1, b: 1 }],
    [['inc a', 'jmp +2', 'tpl a', 'inc a'], { a: 2, b: 0 }],
    [['inc a', 'jio a, +2', 'tpl a', 'inc a'], { a: 2, b: 0 }],
    [['inc a', 'inc a', 'inc a', 'jio a, +2', 'tpl a', 'inc a'], { a: 10, b: 0 }],
    [['inc a', 'jmp +2', 'jmp +2', 'jmp -1'], { a: 1, b: 0 }],
    [['inc a', 'inc a', 'jio a, +2', 'tpl a', 'inc a'], { a: 7, b: 0 }],
    [['inc a', 'jie a, +2', 'inc a'], { a: 2, b: 0 }],
    [['inc a', 'inc a', 'jie a, +2', 'inc a'], { a: 2, b: 0 }],
    [['inc b', 'inc b', 'jie b, +2', 'inc a'], { a: 0, b: 2 }],
  ].forEach((tst:[string[], any]) => {
    let [instructions, registers] = tst;
    let { a, b } = registers;
    it(`runs ${instructions}`, () => {
      let regs = getRegisters(instructions);
      expect(regs.a).to.be.equal(a);
      expect(regs.b).to.be.equal(b);
    });
  });
});
