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

export let matchesLady = (sue, items:any):boolean => {
  return Object.keys(items).map(item => {
    return sue[item] === undefined || sue[item] === items[item];
  })
  .filter(e => e === false)
  .length === 0;
};

export let matchingSues = (inputs:string[], items:any):number[] => {
  return inputs.map(i => parseLine(i))
    .map(sue => matchesLady(sue, items) ? sue.name : -1)
    .filter(sue => sue !== -1);
};
