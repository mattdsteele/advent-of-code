/// <reference path="../../typings/node/node.d.ts" />
import {matchingSues} from './';
import * as fs from 'fs';

let input = fs.readFileSync('./inputs/day-16.txt', 'utf8');

let strings = input.split('\n')
  .filter(x => x.trim() !== '');

let scans = {
  children: 3,
  cats: 7,
  samoyeds: 2,
  pomeranians: 3,
  akitas: 0,
  vizslas: 0,
  goldfish: 5,
  trees: 3,
  cars: 2,
  perfumes: 1
}; 

let sues = matchingSues(strings, scans);
console.log(`silver:`, sues);
