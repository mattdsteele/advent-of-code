/// <reference path="../../typings/node/node.d.ts" />

import getFeet from './day-2';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-2.txt', 'utf8');
let result = input.split('\n')
  .filter(x => x.trim() !== '')
  .map(x => getFeet(x))
  .reduce((prev, curr) => prev + curr, 0);
console.log(result);
