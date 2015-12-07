/// <reference path="../../typings/node/node.d.ts" />

import { makeGates, newIdentity } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-7.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let results = makeGates(strings);
let aVal = results['a']();
console.log(`silver: ${aVal}`);

let newGates = newIdentity(results, 'b', aVal);
console.log(`gold: ${newGates['a']()}`);
