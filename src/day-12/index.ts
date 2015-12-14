let objectTotal = (input, objectBehavior):number => {
  return Object.keys(input)
    .map(key => valueOf(input[key], objectBehavior))
    .reduce((prev, curr) => prev + curr, 0);
};

let noObjectTotal = (input):number => {
  let keys = Object.keys(input);
  if (keys.every(e => input[e] !== 'red')) {
    return objectTotal(input, noObjectTotal);
  }
  return 0;
};

let valueOf = (e, objectBehavior):number => {
  if (Array.isArray(e)) {
    return arrayTotal(e, objectBehavior);
  }
  if (typeof e === 'object') {
    return objectBehavior(e, objectBehavior);
  }
  if (typeof e === 'string') {
    return 0;
  }
  return e;
};

let arrayTotal = (input:any[], objectBehavior):number => {
  return input
  .map(e => valueOf(e, objectBehavior))
  .reduce((prev, curr) => prev + curr, 0);
};

export let numbersSum = (input:string):number => {
  let parsedInput = JSON.parse(input);
  return valueOf(parsedInput, objectTotal)
};

export let noRed = (input:string):number => {
  let parsedInput = JSON.parse(input);
  return valueOf(parsedInput, noObjectTotal)
};
