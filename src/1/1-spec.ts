import { distance, step, taxiDistance, firstDuplicate } from '.';
import * as chai from 'chai';
const expect = chai.expect;

describe('day 1', () => {
  [
    ['R2', 2],
    ['R2, L3', 5],
    ['R2, R2, R2', 2],
    ['R5, L5, R5, R3', 12],
  ].forEach(([directions, blocksAway]:[string, number]) => {
    it(`works with input ${directions}`, () => {
      expect(distance(directions)).to.equal(blocksAway);
    });
  });
  [
    ['N', 'R1', 'E', [1, 0]],
    ['N', 'R10', 'E', [10, 0]],
    ['N', 'L10', 'W', [-10, 0]],
    ['S', 'R1', 'W', [-1, 0]],
    ['S', 'L5', 'E', [5, 0]],
    ['E', 'R5', 'S', [0, -5]],
    ['W', 'R6', 'N', [0, 6]]
  ]
  .forEach(([startOrientation, direction, endOrientation, endPosition]:[string, string, string, number[]]) => {
    it(`works with ${direction}`, () => {
      const input:{position: [number, number], direction: string} = {
        position: [0, 0],
        direction: startOrientation
      };
      const result = step(input, direction);
      expect(result.position).to.eql(endPosition);
      expect(result.direction).to.equal(endOrientation);
    })
  });
  [
    [5, 5, 10],
    [2, -4, 6]
  ].forEach((tst:[number, number, number]) => {
    const [x, y, distance] = tst;
    it(`calculates distance for [${x},${y}]`, () => {
      expect(taxiDistance([x, y])).to.equal(distance);
    });
  });
});

const input = 'L1, L3, L5, L3, R1, L4, L5, R1, R3, L5, R1, L3, L2, L3, R2, R2, L3, L3, R1, L2, R1, L3, L2, R4, R2, L5, R4, L5, R4, L2, R3, L2, R4, R1, L5, L4, R1, L2, R3, R1, R2, L4, R1, L2, R3, L2, L3, R5, L192, R4, L5, R4, L1, R4, L4, R2, L5, R45, L2, L5, R4, R5, L3, R5, R77, R2, R5, L5, R1, R4, L4, L4, R2, L4, L1, R191, R1, L1, L2, L2, L4, L3, R1, L3, R1, R5, R3, L1, L4, L2, L3, L1, L1, R5, L4, R1, L3, R1, L2, R1, R4, R5, L4, L2, R4, R5, L1, L2, R3, L4, R2, R2, R3, L2, L3, L5, R3, R1, L4, L3, R4, R2, R2, R2, R1, L4, R4, R1, R2, R1, L2, L2, R4, L1, L2, R3, L3, L5, L4, R4, L3, L1, L5, L3, L5, R5, L5, L4, L2, R1, L2, L4, L2, L4, L1, R4, R4, R5, R1, L4, R2, L4, L2, L4, R2, L4, L1, L2, R1, R4, R3, R2, R2, R5, L1, L2';
describe('actuals', () => {
  it('runs silver', () => {
    console.log(distance(input));
  });
  it('runs gold', () => {
    const { position } = firstDuplicate(input);
    console.log(taxiDistance(position));
  })
});

describe('gold', () => {
  describe('first dupe', () => {
    [
      ['R8, R4, R4, R8', [4, 0]],
      ['R8, L1, L8, R1, R7, R1', [7, 1]]
    ].forEach(([directions, endPosition]:[string, [number,number]]) => {
      it(`finds the end position for ${directions}`, () => {
        expect(firstDuplicate(directions).position).to.eql(endPosition);
      });
    });
  });
});
