/// <reference path="../../typings/node/node.d.ts" />

import gotoFloor from './day-1';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-1.txt', 'utf8');
console.log(gotoFloor(input));
