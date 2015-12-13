/// <reference path="../typings/mocha/mocha.d.ts" />
/// <reference path="../typings/chai/chai.d.ts" />
import { longestRoute, shortestRoute, parseRoutes, permutations, distanceOf } from '../src/day-9/';
import { expect } from 'chai';

interface Spec {
  routes: string[],
  shortest: number,
  paths: number,
  locations: number,
  longest: number,
  sampleShortest: string[]
};

describe('day 9', () => {
  [
    {
      routes: [
        'London to Dublin = 464',
        'London to Belfast = 518',
        'Dublin to Belfast = 141'
        ],
        shortest: 605,
        paths: 6,
        locations: 3,
        longest: 982,
        sampleShortest: ['London', 'Dublin', 'Belfast']
    }
  ].forEach((test: Spec) => {
    let { routes, shortest, longest } = test;
    it(`gets the shortest route for ${routes}`, () => {
      expect(shortestRoute(routes)).to.equal(shortest);
    });
    it(`gets the longest route for ${routes}`, () => {
      expect(longestRoute(routes)).to.equal(longest);
    });

    describe('parsing behavior', () => {
      it('parses into the right number of elements', () => {
        let parsedRoutes = parseRoutes(routes);
        let { locations } = test;
        expect(parsedRoutes.length).to.equal(locations);
      });
      it('creates the right permutations', () => {
        let perms = permutations(parseRoutes(routes));
        let { paths } = test;
        expect(perms.length).to.equal(paths);
      });
      it('can calculate a route', () => {
        let parsedRoutes = parseRoutes(routes);
        let { sampleShortest, shortest } = test;
        expect(distanceOf(sampleShortest, parsedRoutes)).to.equal(shortest);
      });
    });
  });
});
