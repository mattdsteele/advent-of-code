class BathroomCode {
  private keypad;
  private mapping;
  private silverKeypad = [
    ['1', '2', '3'],
    ['4', '5', '6'],
    ['7', '8', '9']
  ];

  private goldKeypad = [
    [' ', ' ', '1', ' ', ' '],
    [' ', '2', '3', '4', ' '],
    ['5', '6', '7', '8', '9'],
    [' ', 'A', 'B', 'C', ' '],
    [' ', ' ', 'D', ' ', ' '],
  ];

  private silverMapping = {
    '1': [0, 0],
    '2': [0, 1],
    '3': [0, 2],
    '4': [1, 0],
    '5': [1, 1],
    '6': [1, 2],
    '7': [2, 0],
    '8': [2, 1],
    '9': [2, 2]
  };

  private goldMapping = {
    '1': [0, 2],
    '2': [1, 1],
    '3': [1, 2],
    '4': [1, 3],
    '5': [2, 0],
    '6': [2, 1],
    '7': [2, 2],
    '8': [2, 3],
    '9': [2, 4],
    'A': [3, 1],
    'B': [3, 2],
    'C': [3, 3],
    'D': [4, 2],
  };

  constructor(useGoldKeypad = false) {
    this.keypad = useGoldKeypad ? this.goldKeypad : this.silverKeypad;
    this.mapping = useGoldKeypad ? this.goldMapping : this.silverMapping;
  }

  indexFor(num:string):[number, number] {
    return this.mapping[num];
  }

  line(instructions:string, startingPosition:string):string {
    return instructions
    .split('')
    .reduce((position, instruction) => this.direction(instruction, position), startingPosition);
  }

  buildCode(instructions:string[], code:string[], currentPosition:string):string {
    const [nextInstructions, ...remainingInstructions] = instructions;
    const nextPosition = this.line(nextInstructions, currentPosition);
    code.push(nextPosition);
    if (remainingInstructions.length > 0) {
      return this.buildCode(remainingInstructions, code, nextPosition);
    }
    return code.join('');
  }

  codeFor(instructionSet:string):string {
    return this.buildCode(instructionSet.split('\n').filter(x => x !== ''), [], '5');
  }

  numberFor([x, y]:[number, number]):string {
    if (x < 0) x = 0;
    if (x >= this.keypad.length) x = this.keypad.length - 1;
    if (y < 0) y = 0;
    if (y >= this.keypad[0].length) y = this.keypad[0].length - 1;
    return this.keypad[x][y];
  }

  direction(dir:string, currentPosition:string):string {
    let [x, y] = this.indexFor(currentPosition);
    switch(dir) {
      case 'U': 
        x--;
        break;
      case 'D':
        x++;
        break;
      case 'L':
        y--;
        break;
      case 'R':
        y++;
        break;
    }
    const nextNumber = this.numberFor([x, y]);
    if (nextNumber === ' ') {
      return currentPosition;
    }
    return nextNumber;
  }
}

export default BathroomCode;
