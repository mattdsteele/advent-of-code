export function gotoFloor(input: string): number {
  return input.split('')
    .map(ch => ch === '(' ? 1 : -1)
    .reduce((prev, curr) => prev + curr, 0);
};

export function firstBasement(input: string): number {
  let foundBasement = false,
    basementFloor = undefined;
  input.split('')
    .map(ch => ch === '(' ? 1 : -1)
    .reduce((prev, curr, index) => {
      let newFloor = prev + curr;
      if (newFloor === -1 && !foundBasement) {
        basementFloor = index + 1;
        foundBasement = true;
      }
      return newFloor;
    }, 0);
  return basementFloor;
};
