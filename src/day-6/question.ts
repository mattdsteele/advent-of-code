/// <reference path="../../typings/node/node.d.ts" />

import { lightHouse, lightHouseGold } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-6.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

console.log(`silver: ${lightHouse(strings)}`);
console.log(`gold: ${lightHouseGold(strings)}`);
