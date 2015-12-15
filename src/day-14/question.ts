/// <reference path="../../typings/node/node.d.ts" />
import { fastestDeer, goldState, getDeer } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-14.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let fastest = fastestDeer(strings, 2503);
console.log(`silver: ${fastest}`);

let gold = goldState(getDeer(strings), 2503)
  .sort((a, b) => b.points - a.points)[0];
console.log(`gold: ${gold.points}`);
