function getFeet(dimensions: string): number {
  let [length, width, height]: number[] = dimensions
    .split('x')
    .map(x => parseInt(x));
  let sides = [ length * width, length * height, width * height];
  let slack = Math.min(...sides);
  return sides.map(x => 2 * x)
    .reduce((curr, prev) => curr + prev, 0) + slack;
}

export default getFeet;
