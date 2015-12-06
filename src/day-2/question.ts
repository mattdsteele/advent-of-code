/// <reference path="../../typings/node/node.d.ts" />

import { getFeet, calcRibbon } from './day-2';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-2.txt', 'utf8');

let boxes = input.split('\n')
  .filter(x => x.trim() !== '');

let gold = boxes
  .map(x => getFeet(x))
  .reduce((prev, curr) => prev + curr, 0);
console.log(`gold: ${gold}`);

let silver = boxes
  .map(x => calcRibbon(x))
  .reduce((prev, curr) => prev + curr, 0);
console.log(`silver: ${silver}`);
