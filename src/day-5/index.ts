interface NiceCheck {
  (input: string): boolean
}

let minVowels:NiceCheck = input => {
  return input.split('')
    .filter(c => 'aeiou'.indexOf(c) !== -1)
    .length >= 3;
};

let twiceRow:NiceCheck = input => {
  return input.split('')
    .filter((e, i, arr) => {
      return arr[i+1] === e;
    })
    .length > 0;
};

let badStrings:NiceCheck = input => {
  let forbiddenStrings = ['ab', 'cd', 'pq', 'xy'];
  return input.split('')
    .filter((e, i, arr) => {
      let val = `${e}${arr[i+1]}`;
      return forbiddenStrings.indexOf(val) !== -1;
    })
    .length === 0;
};

let makeCompositeCheck = (rules: NiceCheck[]) => {
  let compositeFn: NiceCheck;
  compositeFn = input => {
    return rules
      .map(test => test.call(this, input))
      .filter(result => !result)
      .length === 0;
  };
  return compositeFn;
};

let sandwichedLetter:NiceCheck = input => {
  return input.split('')
    .filter((e, i, arr) => {
      return arr[i+2] === e;
    })
    .length > 0;
};

let repeatingPairs:NiceCheck = input => {
  return input.split('')
    .map((c, i, arr) => {
      if (arr[i+1]) {
        let returned: [string, number];
        returned = [`${c}${arr[i+1]}`, i];
        return returned;
      }
    })
    .filter(i => i !== undefined)
    .filter((e, i, arr) => {
      let [specCh, specIndex] = e;
      return arr.filter(check => {
        let [ch, index] = check;
        return specCh === ch && Math.abs(specIndex - index) > 1;
      })
      .length > 0;
    })
    .length > 0;
};

let isNice:NiceCheck = function(input) {
  return makeCompositeCheck([minVowels, twiceRow, badStrings])(input);
};

let isNiceGold:NiceCheck = input => {
  return makeCompositeCheck([repeatingPairs, sandwichedLetter])(input);
};

export { isNice, isNiceGold };
