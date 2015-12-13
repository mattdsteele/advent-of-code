/// <reference path="../../typings/node/node.d.ts" />

import { shortestRoute, longestRoute } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-9.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

console.log(`silver: ${shortestRoute(strings)}`);
console.log(`gold: ${longestRoute(strings)}`);
