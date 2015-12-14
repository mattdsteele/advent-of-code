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

let lookAndSayRecursive = (input:string):string => startRecursion(input.split(''));
let lookAndSayProcedural = (inStr:string):string => {
  let result = '';
  let input:string[] = inStr.split('');
  let [ currentValue ] = input.splice(0, 1);
  let count = 1;

  while(input.length > 0) {
    let [ nextValue ] = input.splice(0, 1);
    if (nextValue !== currentValue) {
      result += count;
      result += currentValue;
      currentValue = nextValue;
      count = 1;
    } else {
      count++;
    }
  }
  result += count;
  result += currentValue;
  return result;
};

export let lookAndSay = lookAndSayProcedural;
