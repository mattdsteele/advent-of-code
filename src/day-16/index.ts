export let parseLine = (input:string):any => {
  let regex = /Sue (\d*): (\w*): (\d*), (\w*): (\d*), (\w*): (\d*)/g;
  let [_, name, item1, value1, item2, value2, item3, value3] = regex.exec(input);
  let results = {
    name
  };
  results[item1] = parseInt(value1);
  results[item2] = parseInt(value2);
  results[item3] = parseInt(value3);
  return results;
};

let greaterThan = (sue, scanned) => sue > scanned;
let fewerThan = (sue, scanned) => sue < scanned;

let strategy = {
  cats: greaterThan,
  trees: greaterThan,
  pomeranians: fewerThan,
  goldfish: fewerThan
};

let equalityStrategy = (sue, scanned, item) => {
    return sue === undefined || sue === scanned;
};

let scanningStrategy = (sue, scanned, item) => {
  if (sue === undefined) {
    return true;
  }
  let uniqueStrategy = strategy[item];
  if (uniqueStrategy) {
    return uniqueStrategy(sue, scanned);
  }
  return sue === scanned;
};

export let matchesLady = (sue, items:any, strategy=equalityStrategy):boolean => {
  return Object.keys(items).map(item => {
    return strategy(sue[item], items[item], item);
  })
  .filter(e => e === false)
  .length === 0;
};

export let matchingSues = (inputs:string[], items:any):number[] => {
  return inputs.map(i => parseLine(i))
    .map(sue => matchesLady(sue, items, equalityStrategy) ? sue.name : -1)
    .filter(sue => sue !== -1);
};

export let goldScenario = (inputs:string[], items:any):number[] => {
  return inputs.map(i => parseLine(i))
    .map(sue => matchesLady(sue, items, scanningStrategy) ? sue.name : -1)
    .filter(sue => sue !== -1);
};
