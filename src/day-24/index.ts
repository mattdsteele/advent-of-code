/// <reference path="../../typings/node/node.d.ts" />
/// <reference path="../../typings/lodash/lodash.d.ts" />
/// <reference path="../../typings/js-combinatorics/js-combinatorics.d.ts" />

import { flatten } from 'lodash';
let Combinatorics = require('js-combinatorics');

// rosettacode.org/wiki/Power_set#Functional_.28ES_5.29
export let permutations = (input:number[], numbersLeft):number[][] => {
  return Combinatorics.power(input);
};

export let oldPossiblePermutations = (input:number[]) => {
  let firstGroups = permutations(input, 2)
    .filter(x => x.length < input.length / 3);
  console.log(`first group options: ${firstGroups.length}`);

  return flatten(firstGroups.map(first => {
    let rest = input.filter(i => first.indexOf(i) === -1);
    let secondGroups = permutations(rest, 1);
    let secondAndThird = secondGroups.map(second => {
      let third = rest.filter(i => second.indexOf(i) === -1);
      return [second, third];
    });
    return secondAndThird.map(items => {
      let [second, third] = items;
      return [first, second, third];
    });
  }));
};

export let oldEqualGroups = (input:number[]) => {
  let perms = oldPossiblePermutations(input);
  return perms.filter(perm => {
    let [first, second, third] = perm;
    let totalFirst = first.reduce((prev, curr) => prev + curr, 0);
    let totalSecond = second.reduce((prev, curr) => prev + curr, 0);
    let totalThird = third.reduce((prev, curr) => prev + curr, 0);
    return totalFirst === totalSecond && totalFirst === totalThird;
  });
};

export let possiblePermutations = oldPossiblePermutations;
export let equalGroups = oldEqualGroups;

export let oldIdealQe = (input:number[]):number => {
  let groups = oldEqualGroups(input)
    .map(group => group[0])
    .sort((a, b) => a.length - b.length);
  let smallestSize = groups[0].length;
  let possibleOptions = groups.filter(group => group.length === smallestSize);
  return possibleOptions.map(opts => {
    return opts.reduce((prev, curr) => prev * curr, 1);
  })
  .sort((a, b) => a - b)[0];
};

export let idealQe = (input:number[]) => {
  let bestQe = 9999999;
  let permutations = Combinatorics.power(input);
  let optimizedLength = input.length / 3;
  permutations.forEach(first => {
    if (first.length < optimizedLength) {
    }
  });
  return bestQe;
};
