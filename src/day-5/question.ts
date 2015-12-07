/// <reference path="../../typings/node/node.d.ts" />

import { isNice, isNiceGold } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-5.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let silver = strings
  .map(x => isNice(x))
  .reduce((prev, curr) => curr ? prev + 1 : prev, 0);
console.log(`silver: ${silver}`);

let gold = strings
  .map(x => isNiceGold(x))
  .reduce((prev, curr) => curr ? prev + 1 : prev, 0);
console.log(`gold: ${gold}`);
