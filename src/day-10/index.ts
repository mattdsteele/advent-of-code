let startRecursion = numbers => {
  return recurse('', numbers[0], 1, numbers.slice(1));
};

let recurse = (results:string, currentValue:string, count:number, restOfString):string => {
  if (restOfString.length == 0) {
    return results + count + currentValue;
  }
  let nextValue = restOfString[0];
  if (nextValue !== currentValue) {
    return recurse(results + count + currentValue, nextValue, 1, restOfString.slice(1));
  }
  return recurse(results, nextValue, count + 1, restOfString.slice(1));
};

export let lookAndSay = (input:string):string => startRecursion(input.split(''));
