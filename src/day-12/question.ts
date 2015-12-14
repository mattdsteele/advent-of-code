/// <reference path="../../typings/node/node.d.ts" />

import { numbersSum, noRed } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-12.txt', 'utf8').trim();

console.log(`silver: ${numbersSum(input)}`);
console.log(`gold: ${noRed(input)}`);
