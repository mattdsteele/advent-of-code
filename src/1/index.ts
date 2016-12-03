import { flatMap, fill } from 'lodash';

export interface Status {
  position: [number, number],
  direction: string
}

const applySteps = (instructions:string[]):Status => {
  const startingPosition:Status = {
    direction: 'N',
    position: [0, 0]
  };
  return instructions.reduce((prev, current) => step(prev, current), startingPosition);
};

const distanceToSingleSteps = (route: string) => {
  const steps = route.split(',').map(x => x.trim());
  const individualSteps = flatMap(steps, (step) => {
    const [rotation, ...amountStr] = step.split('');
    const numberSteps = parseInt(amountStr.join(''));
    const transformed = fill(Array(numberSteps), `S1`);
    transformed[0] = `${rotation}1`;
    return transformed;
  });
  return individualSteps;
};

const distance = (route: string): number => {
  const endStatus = applySteps(distanceToSingleSteps(route));
  return taxiDistance(endStatus.position);
};

const applyMapping = (mappings:{[dir:string]: number}, direction:string, amount:number) => {
  const multiplier:number = mappings[direction] || 0;
  return amount * multiplier;
};

const xDelta = (direction:string, amount:number) => applyMapping({ E: 1, W: -1 }, direction, amount);
const yDelta = (direction:string, amount:number) => applyMapping({ N: 1, S: -1 }, direction, amount);

const calculateRotationIndex = (rotation: string) => {
  switch(rotation) {
    case 'R':
      return 1;
    case 'L':
      return -1;
    default:
      return 0;
  }
}

const step = (current:Status, instruction:string):Status => {
  const [rotation, ...amountStr] = instruction.split('');
  const amount = parseInt(amountStr.join(''));

  const directions = ['N', 'E', 'S', 'W'];
  const currentDirectionIndex = directions.indexOf(current.direction);
  const newDirectionIndex = (currentDirectionIndex + calculateRotationIndex(rotation) + 4) % 4;
  const newDirection = directions[newDirectionIndex];


  const newX = current.position[0] + xDelta(newDirection, amount);
  const newY = current.position[1] + yDelta(newDirection, amount);

  return {
    direction: newDirection,
    position: [newX, newY]
  };
};

const taxiDistance = (endPosition:[number, number]):number => {
  return endPosition
  .map(x => Math.abs(x))
  .reduce((prev, curr) => prev + curr, 0);
};

const hasEntry = (visitedLocations:Set<[number,number]>, entry:[number,number]):boolean => {
  let foundValue = false;
  visitedLocations.forEach(visited => {
    if (visited[0] == entry[0] && visited[1] == entry[1]) {
      foundValue = true;
      return;
    }
  });
  return foundValue;
};

const firstDuplicate = (directions:string):Status => {
  const steps = distanceToSingleSteps(directions);
  let position:Status = {
    direction: 'N',
    position: [0, 0]
  };
  const visitedLocations = new Set();
  let foundPosition:Status;
  for(const nextStep of steps) {
    position = step(position, nextStep);
    if (hasEntry(visitedLocations, position.position)) {
      foundPosition = position;
      return foundPosition;
    }
    visitedLocations.add(position.position);
  }
  return foundPosition;
};
export { distance, step, taxiDistance, firstDuplicate };
