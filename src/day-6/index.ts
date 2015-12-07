class House {
  len: number;
  width: number;
  mappings;
  houseArr: boolean[];
  steps: any[];
  constructor(steps, mappings) {
    this.len = 1000;
    this.width = 1000;
    this.mappings = mappings;
    this.houseArr = new Array(this.len * this.width);
    this.handle(0, 0, this.len - 1, this.width - 1, mappings.init);
    this.steps = steps.map(step => this.setupStep(step));

    for (let [fn, fromX, fromY, toX, toY] of this.steps) {
      this.handle(fromX, fromY, toX, toY, fn);
    }
  }

  setupStep(stepStr) {
    let regex = /(turn off|turn on|toggle) (\d*),(\d*) through (\d*),(\d*)/g;
    let results = regex.exec(stepStr);
    let fnMapping = {
      'turn on': this.mappings.turnOn,
      'turn off': this.mappings.turnOff,
      'toggle': this.mappings.toggle
    };

    let fn = fnMapping[results[1]];
    return [fn, parseInt(results[2]), parseInt(results[3]), parseInt(results[4]), parseInt(results[5])];
  }

  handle(fromX: number, fromY: number, toX: number, toY: number, fn) {
    for(let x = fromX; x <= toX; x++) {
      for (let y = fromY; y <= toY; y++) {
        this.houseArr[(x * this.len) + y] = fn(this.houseArr[(x * this.len) + y]);
      }
    }
  }

  litLights() {
    return this.mappings.numberLit(this.houseArr);
  }
}

export function lightHouse(steps: string[]): number {
  let house = new House(steps, {
    init: val => false,
    turnOff: val => false,
    turnOn: val => true,
    toggle: val => !val,
    numberLit: h => {
      return h.reduce((prev, curr) => {
        if (curr) {
          return prev + 1;
        }
        return prev;
      }, 0);
    }
  });
  return house.litLights();
}

export function lightHouseGold(steps: string[]): number {
  let house = new House(steps, {
    init: val => 0,
    turnOff: val => {
      if (val > 0) {
        return val-1;
      }
      return val;
    },
    turnOn: val => val+1,
    toggle: val => val+2,
    numberLit: h => h.reduce((prev, curr) => prev + curr, 0)
  });
  return house.litLights();
}
