/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { parseLine, distanceAt, getDeer, goldState } from '../src/day-14/';
import { expect } from 'chai';

describe('day 14', () => {
  it('can parse a line', () => {
    let spec = 'Vixen can fly 19 km/s for 7 seconds, but then must rest for 124 seconds.';
    let { reindeer, speed, flyingTime, restingTime } = parseLine(spec);
    expect(reindeer).to.equal('Vixen');
    expect(speed).to.equal(19);
    expect(flyingTime).to.equal(7);
    expect(restingTime).to.equal(124);
  });
  describe('comet', () => {
    let line = 'Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.';
    let comet = parseLine(line);
    [
      [1, 14],
      [10, 140],
      [11, 140],
      [1000, 1120]
    ].forEach(timestamp => {
      it(`tests comet at ${timestamp}`, () => {
        let [seconds, distance] = timestamp;
        expect(distanceAt(comet, seconds)).to.equal(distance);
      });
    });
  });
  describe('dancer', () => {
    let line = 'Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.';
    let dancer = parseLine(line);
    [
      [1, 16],
      [10, 160],
      [11, 176],
      [1000, 1056]
    ].forEach(timestamp => {
      it(`tests dancer at ${timestamp}`, () => {
        let [seconds, distance] = timestamp;
        expect(distanceAt(dancer, seconds)).to.equal(distance);
      });
    });
  });
  describe('gold', () => {
    let lines = [
      'Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.',
      'Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.'
    ];
    let deer = getDeer(lines);
    [
      [1, 1, 0],
      [140, 139, 1],
      [1000, 689, 312]
    ].forEach(spec => {
      it.skip('calculates proper pointage', () => {
        let [ timestamp, dancer, comet ] = spec;
        let gameState = goldState(deer, timestamp);
        expect(gameState.filter(s => s.name === 'Dancer')[0].points).to.equal(dancer);
        expect(gameState.filter(s => s.name === 'Comet')[0].points).to.equal(comet);
      });
    });
  });
});
