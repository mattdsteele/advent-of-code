/// <reference path="../../typings/node/node.d.ts" />
import {combinations, uniqueCombinations} from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-17.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let silver = combinations(strings, 150);
console.log(`silver: ${silver}`);
let gold = uniqueCombinations(strings, 150);
console.log(`gold: ${gold}`);
