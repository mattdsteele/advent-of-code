function stringify(x, y): string {
  return [x, y].join(',');
}

function updatePosition(currX, currY, position): number[] {
  switch (position) {
    case '^':
      currY++;
      break;
    case 'v':
      currY--;
      break;
    case '<':
      currX--;
      break;
    case '>':
      currX++;
      break;
  }
  return [currX, currY];
}

function deliverPresents(routeStops: string[], houses) {
  let [currX, currY] = [0, 0];
  for (let position of routeStops) {
    [currX, currY] = updatePosition(currX, currY, position);
    houses[stringify(currX, currY)] = true;
  }
}

export function housesDelivered(route: string): number {
  let houses = {};

  houses[stringify(0, 0)] = true;

  let routeStops = route.split('');
  deliverPresents(routeStops, houses);

  return Object.keys(houses).length;
}

export function roboSanta(route: string): number {
  let houses = {};
  houses[stringify(0, 0)] = true;
  let routeStops = route.split('');
  deliverPresents(routeStops.filter((x, i) => i % 2 === 0), houses);
  deliverPresents(routeStops.filter((x, i) => i % 2 === 1), houses);

  return Object.keys(houses).length;
}
