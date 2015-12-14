/// <reference path="../../typings/lodash/lodash.d.ts" />
import { uniq, flatten } from 'lodash';

interface Comparison {
  left: string;
  right: string;
  delta: number;
}

export let parseLine = (input:string):Comparison => {
  let regex = /^(\w*) would (\w*) (\d*) happiness units by sitting next to (\w*)\.$/g;
  let [_, left, direction, amount, right] = regex.exec(input);
  let delta = parseInt(amount) * (direction === 'gain' ? 1 : -1);
  return { left, right, delta };
};

export let allNames = (input:Comparison[]):string[] => {
  return uniq(input.map(e => e.left));
};

let setupSeatings = (current:string[], remaining: string[]) => {
  if (remaining.length === 0) {
    return current;
  }
  return remaining.map(person => {
    return setupSeatings(current.concat(person), remaining.filter(i => i !== person));
  });
};

let flattenOut = (val, flatteningsLeft) => {
  if(flatteningsLeft === 0) {
    return val;
  }
  return flattenOut(flatten(val), flatteningsLeft - 1);
};

export let seatingPermutations = (input:string[]) => {
  return flattenOut(input.map(person => {
    return setupSeatings([person], input.filter(i => i !== person));
  }), input.length - 1);
};

let calculateSeating = (total, index, optionsList, comparisons) => {
  let deltaLeft = comparisons.find(comparison => {
    return comparison.left === optionsList[index] && comparison.right === optionsList[(index + 1) % optionsList.length];
  }).delta;
  let deltaRight = comparisons.find(comparison => {
    return comparison.right === optionsList[index] && comparison.left === optionsList[(index + 1) % optionsList.length];
  }).delta;
  if (index === optionsList.length - 1) {
    return total + deltaLeft + deltaRight;
  }
  return calculateSeating(total + deltaLeft + deltaRight, index + 1, optionsList, comparisons);
};

export let optimalFromComparisons = comparisons => {
  let names = allNames(comparisons);
  let seatingOptions = seatingPermutations(names);
  return seatingOptions
    .map(option => calculateSeating(0, 0, option, comparisons))
    .sort((a, b) => b - a)[0]
};

export let addMyself = comparisons => {
  let names = allNames(comparisons);
  names.forEach(name => {
    comparisons.push({
      left: 'Me',
      right: name,
      delta: 0
    });
    comparisons.push({
      left: name,
      right: 'Me',
      delta: 0
    });
  });
  return comparisons;
};

export let optimalHappiness = (input:string[]):number => {
  let comparisons = input.map(i => parseLine(i));
  return optimalFromComparisons(comparisons);
};

