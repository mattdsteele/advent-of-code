/// <reference path="../../typings/lodash/lodash.d.ts" />
import { uniq, flatten } from 'lodash';

let calculatedDistances = (distances:string[]):number[] => {
  let parsedRoutes = parseRoutes(distances);
  return permutations(parsedRoutes)
    .map(path => distanceOf(path, parsedRoutes))
    .sort((a, b) => a - b);
};

export let longestRoute = (distances:string[]):number => {
  let ordered = calculatedDistances(distances);
  return ordered[ordered.length - 1];
};

export let shortestRoute = (distances:string[]):number => {
  return calculatedDistances(distances)[0];
};

interface Segment {
  first: string,
  second: string,
  distance: number
};

interface Route {
  route: string[],
  distance: number
};

export let parseRoutes = (routes: string[]): Segment[] => {
 let r = routes.map(route => {
    let routeRegex = /^(\w*) to (\w*) = (\d*)$/;
    let [_, first, second, distance] = routeRegex.exec(route);
    return {
      first, 
      second,
      distance: parseInt(distance)
    };
  })
  return r;
};

let makeTravelTree = (loc, nextLoc) => {
  let nextRoutes = nextLoc
    .filter(x => x !== loc)
    .map((x, i, arr) => makeTravelTree(x, arr));
  return {
    start: loc,
    routes: nextRoutes
  };
};

let flattenTree = (route, tree) => {
  let nextRouteLocation = route.concat(tree.start);
  if (tree.routes.length === 0) {
    return route;
  }
  return tree.routes.map(e => {
    return flattenTree(nextRouteLocation, e);
  });
};

export let permutations = (parsedRoutes: Segment[]) => {
  let locations = uniq(parsedRoutes.reduce((acc, route) => {
    acc.push(route.first);
    acc.push(route.second);
    return acc;
  }, []));

  let startTravelTrees = locations.map(e => makeTravelTree(e, locations));

  let totalDistance = startTravelTrees.map(e => makePaths([e.start], e));
  return flattenOut(totalDistance, locations.length - 1);
};

let flattenOut = (val, flatteningsLeft) => {
  if (flatteningsLeft === 0) {
    return val;
  }
  return flattenOut(flatten(val), flatteningsLeft - 1);
}

let makePaths = (acc, leaf) => {
  if (leaf.routes.length === 0) {
    return acc;
  }
  return leaf.routes.map(r => makePaths(acc.concat(r.start), r));
};

export let distanceOf = (route:string[], parsedRoutes:Segment[]):number => {
  let allSegments = parsedRoutes.reduce((acc, route) => {
    acc.push(route);
    acc.push({
      first: route.second,
      second: route.first,
      distance: route.distance
    });
    return acc;
  }, []);

  return calculateDistance(0, route, allSegments);
};

let calculateDistance = (total, restOfRoute, allSegments) => {
  let newTotal = total + allSegments.find(seg => {
    return seg.first === restOfRoute[0] && seg.second == restOfRoute[1];
  }).distance;
  if (restOfRoute.length === 2) {
    return newTotal;
  }
  return calculateDistance(newTotal, restOfRoute.slice(1), allSegments);
};
