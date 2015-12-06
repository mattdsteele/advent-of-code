function gotoFloor(input: string): number {
  return input.split('')
    .map(ch => ch === '(' ? 1 : -1)
    .reduce((prev, curr) => prev + curr, 0);
}

export default gotoFloor;
