/// <reference path="../../typings/node/node.d.ts" />
import {stepTimes, stepsWithCornersStuck} from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-18.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let silver = stepTimes(strings, 100);
console.log(`silver: ${silver}`);

let gold = stepsWithCornersStuck(strings, 100);
console.log(`gold: ${gold}`);
