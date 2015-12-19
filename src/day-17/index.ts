/// <reference path="../../typings/lodash/lodash.d.ts" />
import { eq } from 'lodash';

let generatePermutations = (allOptions, currentValue, containers) => {
  if (containers.length === 1) {
    allOptions.push(currentValue);
    allOptions.push(currentValue.concat(containers[0]));
  } else {
    generatePermutations(allOptions, currentValue, containers.slice(1));
    generatePermutations(allOptions, currentValue.concat(containers[0]), containers.slice(1));
  }
};

let addsTo = (permutations, total) => {
  return permutations
    .map(permutation => {
      return permutation
        .reduce((prev, curr) => prev + curr, 0)
    })
    .filter(i => i === total);
};

export let combinations = (containers: any[], total):number => {
  containers = containers.map(i => parseInt(i));
  let permutations = [];
  generatePermutations(permutations, [], containers);
  return addsTo(permutations, total).length;
};

export let uniqueCombinations = (containers: any[], total):number => {
  containers = containers.map(i => parseInt(i));
  let permutations = [];
  generatePermutations(permutations, [], containers);
  
  let matches = permutations.filter(perm => {
    return perm.reduce((prev, curr) => prev + curr, 0) === total;
  }).sort((a, b) => a.length - b.length);
  let minLength = matches[0].length;

  return matches.filter(u => u.length === minLength).length;
};
