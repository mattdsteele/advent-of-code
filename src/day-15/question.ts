/// <reference path="../../typings/node/node.d.ts" />
import { bestScore } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-15.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let best = bestScore(strings);
console.log(`silver: ${best}`);
