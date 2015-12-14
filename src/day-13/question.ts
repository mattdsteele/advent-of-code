/// <reference path="../../typings/node/node.d.ts" />

import { optimalFromComparisons, optimalHappiness, parseLine, addMyself } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-13.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

console.log(`silver: ${optimalHappiness(strings)}`);

let comparisons = strings.map(i => parseLine(i));
let withMyself = addMyself(comparisons);
console.log(`gold: ${optimalFromComparisons(withMyself)}`);
