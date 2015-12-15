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

let deerWithPositions = (deer:Reindeer[], timestamp:number) => {
  return deer
    .map(d => { 
      return { deer: d.reindeer, position: distanceAt(d, timestamp) }
    });
};

let fastest = (deer, timestamp) => {
  return deerWithPositions(deer, timestamp)
    .map(deer => deer.position)
    .sort((a, b) => b - a)[0];
};

export let fastestDeer = (lines:string[], timestamp:number):number => {
  return fastest(getDeer(lines), timestamp);
};

export let getDeer = (lines:string[]):Reindeer[] => {
  return lines.map(line => parseLine(line));
};

export let goldState = (deer:Reindeer[], timestamp: number) => {
  let deerWithPoints:any[] = deer.map(d => {
    return { name: d.reindeer, points: 0 };
  });
  for (let i = 1; i <= timestamp; i++) {
    var deerAt = deerWithPositions(deer, i)
      .sort((a, b) => b.position - a.position);
    let winners = deerAt.filter(d => d.position === deerAt[0].position);
    winners.forEach(d => {
      var point = deerWithPoints.filter(deerPoint => deerPoint.name === d.deer)[0];
      point.points = point.points + 1;
    });
  }
  return deerWithPoints;
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
