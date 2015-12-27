/// <reference path="../../typings/node/node.d.ts" />
/// <reference path="../../typings/lodash/lodash.d.ts" />
/// <reference path="../../typings/js-combinatorics/js-combinatorics.d.ts" />

import { flatten } from 'lodash';
let Combinatorics = require('js-combinatorics');

export let equalPerms = (inputs, groups):number[][] => {
  let groupings = inputs.reduce((prev, curr) => prev + curr, 0) / groups;
  let maxSize = inputs.length / groups;
  let perms = [];
  for (let i = 1; i < maxSize; i++) {
    perms = perms.concat(Combinatorics.combination(inputs, i).filter(i => {
      return i.reduce((prev, curr) => prev + curr, 0) === groupings;
    }));
  }
  return perms;
};

export let idealQe = (inputs, numberGroups = 3) => {
  let perms = equalPerms(inputs, numberGroups);

  let shortest = perms.sort((a, b) => {
    return a.length - b.length;
  })[0].length;

  let possiblePerms = perms.filter(i => i.length === shortest);

  return possiblePerms.map(i => {
    return i.reduce((prev, curr) => prev * curr, 1);
  })
  .sort((a, b) => a - b)[0];
};
