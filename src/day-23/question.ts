/// <reference path="../../typings/node/node.d.ts" />
import {getRegisters} from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-23.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let registers = getRegisters(strings);
console.log(`silver: ${registers.b}`);

let gold = getRegisters(strings, 1);
console.log(`gold: ${gold.b}`);
