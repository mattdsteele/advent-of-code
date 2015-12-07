/// <reference path="../../typings/node/node.d.ts" />

import { gateValues } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-7.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let results = gateValues(strings);
console.log(`silver: ${results['a']()}`);
