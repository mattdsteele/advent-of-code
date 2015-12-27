export let orderNumber = (row:number, column:number):number => {
  if (column === 1 && row === 1) {
    return 1;
  }

  if (column === 1) {
    return orderNumber(1, row - 1) + 1;
  }

  return orderNumber(row, column - 1) + column + row - 1;
};

export let code = (row, column):number => {
  let val = 20151125;
  let iteration = orderNumber(row, column);
  for (let i = 1; i < iteration; i++) {
    val = (val * 252533) % 33554393;
  }
  return val;
};
