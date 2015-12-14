/// <reference path="../../typings/lodash/lodash.d.ts" />
import { repeat } from 'lodash';

export let nextPassword = (input:string):string => {
  while (true) {
    let next = increment(input);
    if (meetsRequirements(next)) {
      return next;
    }
    input = next;
  }
};

let incrementingChars = (input:string[]):boolean => {
  return input.filter((el, i, arr) => {
    if (arr[i+1] && arr[i+2]) {
      let currentCharCode = el.charCodeAt(0);
      return String.fromCharCode(currentCharCode + 1) === arr[i+1] && String.fromCharCode(currentCharCode + 2) === arr[i+2];
    } else {
      return false;
    }
  }).length > 0;
};

let noForbiddenLetters = (input:string[]):boolean => {
  let forbiddenLetters = ['i', 'o', 'l'];
  return input
    .filter(letter => forbiddenLetters.indexOf(letter) !== -1)
    .length === 0;
};

let twoRepeatingElements = (input:string[]):boolean => {
  return input
    .map((ch, i) => input[i+1] === ch )
    .filter((ch, i, arr) => ch && !(arr[i-1] && !arr[i-2]))
    .length > 1;

};

export let meetsRequirements = (input:string):boolean => {
  let inputElements = input.split('');
  return [
    incrementingChars,
    noForbiddenLetters,
    twoRepeatingElements
  ]
  .map(e => e(inputElements))
  .every(e => e);
};

let incrementedChars = (input, totalWrapped:number):number => {
  let lastElement = input[input.length - totalWrapped];
  if (lastElement === 'z') {
    return incrementedChars(input, totalWrapped + 1);
  }
  return totalWrapped;
};

export let increment = (input:string):string => {
  let inputElements = input.split('');

  let elementsToReplace = incrementedChars(inputElements, 1);
  let newLastChar = String.fromCharCode(inputElements[inputElements.length - elementsToReplace].charCodeAt(0) + 1);

  inputElements.splice(inputElements.length - elementsToReplace, elementsToReplace, newLastChar, repeat('a', elementsToReplace - 1));

  return inputElements.join('');
};
