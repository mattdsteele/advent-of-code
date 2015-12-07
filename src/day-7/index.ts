/// <reference path="../../typings/node/node.d.ts" />
var UInt = require('uint-js').UInt;

let memoize = fn => {
  let memoized = function() {
    var cached = memoized['cache'];
    if (!cached) {
      let result = fn.apply(this);
      memoized['cache'] = result;
    }
     return memoized['cache'];
  };
  return memoized;
};

export function gateValues(steps: string[]) {

  let gates = {};

  let notGate = fromGate => {
    return () => {
      let originalGate = gates[fromGate]();
      return new UInt({value: originalGate, bits: 16}).neg().value();
    };
  };

  let identityGate = val => {
    return () => val;
  };

  let passthroughGate = val => {
    return () => gates[val]();
  }

  let nonUnaryGate = (left, right, operator) => {
    return () => {
      let fnize = check => {
        if (isNaN(check)) {
          return passthroughGate(check);
        }
        return identityGate(check);
      };
      let operations = {
        AND: (left, right) => fnize(left)() & fnize(right)(),
        OR: (left, right) => fnize(left)() | fnize(right)(),
        LSHIFT: (left, right) => fnize(left)() << parseInt(right),
        RSHIFT: (left, right) => fnize(left)() >> parseInt(right)
      };
      return operations[operator](left, right);
    };
  };

  for (let step of steps) {
    let assignRegex = /^(\w*) -> (\w*)$/g;
    let result = assignRegex.exec(step);
    if (result) {
      let [_, gateValStr, gate] = result;
      var parsedStringTest = parseInt(gateValStr);
      if (isNaN(parsedStringTest)) {
        gates[gate] = memoize(passthroughGate(gateValStr));
      } else {
        gates[gate] = memoize(identityGate(parsedStringTest));
      }
    } else {
      let notRegex = /^NOT (\w*) -> (\w*)$/g;
      let result = notRegex.exec(step);
      if (result) {
        let [_, fromGate, toGate] = result;
        gates[toGate] = memoize(notGate(fromGate));
      } else {
        let nonUnaryRegex = /^(\w*) (\w*) (\w*) -> (\w*)$/g;
        let result = nonUnaryRegex.exec(step);
        if (!result) {
          throw new Error(`did not get past this step: ${step}`);
        }
        let [_, firstLeft, operator, secondLeft, gate] = result;
        gates[gate] = memoize(nonUnaryGate(firstLeft, secondLeft, operator));
      }
    }
  }
  return gates;
}
