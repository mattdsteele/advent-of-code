/// <reference path="../../typings/rx/rx.all.d.ts" />
/// <reference path="../../typings/rx/rx.testing.d.ts" />
import * as Rx from 'rx';

interface Reindeer {
  reindeer: string;
  speed: number;
  flyingTime: number;
  restingTime: number;
};

export let parseLine = (line:string):Reindeer => {
  let lineRegex = /^(\w*) can fly (\d*) km\/s for (\d*) seconds, but then must rest for (\d*) seconds.$/g;
  let [ _, reindeer, speed, flyingTime, restingTime ] = lineRegex.exec(line);
  return {
    reindeer,
    speed: parseInt(speed),
    flyingTime: parseInt(flyingTime),
    restingTime: parseInt(restingTime)
  };
};

export let fastestDeer = (lines:string[], timestamp:number):number => {
  return lines.map(line => parseLine(line))
    .map(deer => distanceAt(deer, timestamp))
    .sort((a, b) => b - a)[0];
};

export let distanceAt = (reindeer:Reindeer, timestamp:number):number => {
  let currentPosition = 0;
  let isFlying = true;
  let stateLeft = reindeer.flyingTime;

  let scheduler = new Rx.TestScheduler();
  let source = Rx.Observable.interval(1, scheduler);
  let subscription = source.subscribe(x => {
    currentPosition += isFlying ? reindeer.speed : 0;
    stateLeft--;
    if (stateLeft === 0) {
      isFlying = !isFlying;
      stateLeft = isFlying ? reindeer.flyingTime : reindeer.restingTime;
    }
  },
    err => {},
    () => {}
  );
  scheduler.advanceTo(timestamp);
  return currentPosition;
};
