/// <reference path="../typings/mocha/mocha.d.ts" />
import app from '../src/day-1';

describe('Calculator', () => {
  var subject : app;

  beforeEach(function () {
    subject = new app();
  });

  describe('#add', () => {
    it('should add two numbers together', () => {
      const result : number = subject.add(2, 3);
      if (result !== 5) {
        throw new Error('Expected 2 + 3 = 5 but was ' + result);
      }
    });
  });
});
