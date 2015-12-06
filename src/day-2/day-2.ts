function sides(dimensions: string): number[] {
  return dimensions
    .split('x')
    .map(x => parseInt(x));
}

export function getFeet(dimensions: string): number {
  let [length, width, height] = sides(dimensions);
  let surface = [ length * width, length * height, width * height];
  let slack = Math.min(...surface);
  return surface.map(x => 2 * x)
    .reduce((curr, prev) => curr + prev, 0) + slack;
};

export function calcRibbon(dimensions: string): number {
  let dims = sides(dimensions);
  let maxSide = Math.max(...dims);
  let [length, width, height] = dims;
  let extra = length * width * height;
  let around = dims
    .sort((a, b) => a - b)
    .filter((el, index) => index + 1 !== dims.length)
    .map(x => x * 2)
    .reduce((prev, curr) => prev + curr, 0);
  return extra + around;
};
