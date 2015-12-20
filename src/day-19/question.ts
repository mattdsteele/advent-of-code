/// <reference path="../../typings/node/node.d.ts" />
import {calibrate, stepsUntil} from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-19.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let silver = calibrate(strings);
console.log(`silver: ${silver.length}`);

let goldStr = input.split('\n')
  .filter(x => x.trim() !== '');
let gold = stepsUntil(goldStr);
console.log(`gold: ${gold}`);
