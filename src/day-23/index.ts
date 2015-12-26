interface Registers {
  a: number,
  b: number
};

let behaviors = {
  inc: register => {
    return (registers) => {
      registers[register]++;
      registers.next++;
      return registers;
    };
  },
  hlf: register => {
    return (registers) => {
      registers[register] /= 2;
      registers.next++;
      return registers;
    }
  },
  tpl: register => {
    return (registers) => {
      registers[register] *= 3;
      registers.next++;
      return registers;
    };
  },
  jmp: nextOne => {
    return registers => {
      registers.next += parseInt(nextOne);
      return registers;
    };
  },
  jio: val => {
    return registers => {
      let regex = /(\w), (.*)/;
      let [_, register, amount] = regex.exec(val);
      if (registers[register] === 1) {
        registers.next += parseInt(amount);
      } else {
        registers.next++;
      }
      return registers;
    }
  },
  jie: val => {
    return registers => {
      let regex = /(\w), (.*)/;
      let [_, register, amount] = regex.exec(val);
      if (registers[register] % 2 === 0) {
        registers.next += parseInt(amount);
      } else {
        registers.next++;
      }
      return registers;
    }
  }
};
export let getRegisters = (instructions:string[], aStart = 0):Registers => {
  let mappedInstructions = instructions.map(instruction => {
    let startVal = instruction.substring(0, 3);
    let register = instruction.substring(4);
    return behaviors[startVal](register);
  });
  let state = {
    a: aStart,
    b: 0,
    next: 0
  };
  while(true) {
    let next = mappedInstructions[state.next];
    if (!next) {
      break;
    }
    state = next(state);
  };
  return state;
};
