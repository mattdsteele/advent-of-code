/// <reference path="../../typings/lodash/lodash.d.ts" />
import { flatten, uniq } from 'lodash';

export let parseRule = (input:string):{ left:string, right:string } => {
  let regex = /(\w*) => (\w*)/;
  let [ _, left, right ] = regex.exec(input);
  return { left, right };
};

export let distinctMolecules = (rulesStr:string[], input:string):string[] => {
  let rules = rulesStr.map(r => parseRule(r));
  let inputValues = input.split('');
  return uniq(flatten(rules.map(rule => {
    let leftArr = rule.left.split('');
    let appliedRules = [];
    inputValues.forEach((input, i) => {
      if (leftArr.map((token, j) => inputValues[i + j] === token).every(l => l === true)) {
        let inputClone = inputValues.join('').split('');
        inputClone.splice(i, leftArr.length, rule.right);
        appliedRules.push(inputClone.join(''));
      }
    });
    return appliedRules;
  })));
};

export let calibrate = (input:string[]):string[] => {
  let molecule = input.pop();
  return distinctMolecules(input.filter(x => x.trim() !== ''), molecule);
};

export let stepsUntil = (input:string[]):number => {
  let steps = 0;
  let expected = input.pop();
  let rules = input.filter(r => r.trim() !== '').map(r => parseRule(r));
  let molecule = expected;
  while (molecule !== 'e') {
    let interestingValue = rules.filter(rule => {
      return molecule.indexOf(rule.right) !== -1;
    })[0];
    let moleculeArr = molecule.split('');
      moleculeArr.splice(molecule.indexOf(interestingValue.right), interestingValue.right.length, interestingValue.left);
      molecule = moleculeArr.join('');
    steps++;
  }
  return steps;
};
