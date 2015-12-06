/// <reference path="../../typings/node/node.d.ts" />

import { gotoFloor, firstBasement } from './day-1';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-1.txt', 'utf8');
console.log(`gold: ${gotoFloor(input)}`);
console.log(`silver: ${firstBasement(input)}`);
