/// <reference path="../../typings/node/node.d.ts" />
import { fastestDeer } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-14.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let fastest = fastestDeer(strings, 2503);
console.log(`silver: ${fastest}`);
