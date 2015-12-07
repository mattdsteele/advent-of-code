class House {
  len: number;
  width: number;
  houseArr: boolean[];
  steps: any[];
  constructor(steps) {
    this.len = 1000;
    this.width = 1000;
    this.houseArr = new Array(this.len * this.width);
    this.turnOff(0, 0, this.len - 1, this.width - 1);
    this.steps = steps.map(step => this.setupStep(step));

    for (let [fn, fromX, fromY, toX, toY] of this.steps) {
      fn.call(this, fromX, fromY, toX, toY);
    }
  }

  setupStep(stepStr) {
    let regex = /(turn off|turn on|toggle) (\d*),(\d*) through (\d*),(\d*)/g;
    let results = regex.exec(stepStr);
    let fnMapping = {
      'turn on': this.turnOn,
      'turn off': this.turnOff,
      'toggle': this.toggle
    };

    let fn = fnMapping[results[1]];
    return [fn, parseInt(results[2]), parseInt(results[3]), parseInt(results[4]), parseInt(results[5])];
  }

  turnOff(fromX: number, fromY: number, toX: number, toY: number) {
    for(let x = fromX; x <= toX; x++) {
      for (let y = fromY; y <= toY; y++) {
        this.houseArr[(x * this.len) + y] = false;
      }
    }
  }

  turnOn(fromX: number, fromY: number, toX: number, toY: number) {
    for(let x = fromX; x <= toX; x++) {
      for (let y = fromY; y <= toY; y++) {
        this.houseArr[(x * this.len) + y] = true;
      }
    }
  }
  
  toggle(fromX, fromY, toX, toY) {
    for(let x = fromX; x <= toX; x++) {
      for (let y = fromY; y <= toY; y++) {
        this.houseArr[(x * this.len) + y] = !this.houseArr[(x * this.len) + y];
      }
    }
  }

  litLights() {
    return this.houseArr.reduce((prev, curr) => {
      if (curr) {
        return prev + 1;
      }
      return prev;
    }, 0);
  }
}

export function lightHouse(steps: string[]): number {
  let house = new House(steps);
  return house.litLights();
}

export function lightHouseGold(steps: string[]): number {
  return 0;
}
