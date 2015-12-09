/// <reference path="../../typings/node/node.d.ts" />

import { numChars, expand } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-8.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let result = strings
  .map(x => x.length - numChars(x))
  .reduce((prev, curr) => prev + curr, 0);

console.log(`silver: ${result}`);

let gold = strings
  .map(x => expand(x) - x.length)
  .reduce((prev, curr) => prev + curr, 0);

console.log(`gold: ${gold}`);
