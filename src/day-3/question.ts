/// <reference path="../../typings/node/node.d.ts" />

import { housesDelivered, roboSanta } from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-3.txt', 'utf8');
console.log(`silver: ${housesDelivered(input)}`);
console.log(`gold: ${roboSanta(input)}`);
