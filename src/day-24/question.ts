/// <reference path="../../typings/node/node.d.ts" />
import { idealQe } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-24.txt', 'utf8');

let numbers = input.split('\n')
  .filter(x => x.trim() !== '')
  .map(x => parseInt(x));

console.log(`silver: ${idealQe(numbers)}`);
console.log(`gold: ${idealQe(numbers, 4)}`);
