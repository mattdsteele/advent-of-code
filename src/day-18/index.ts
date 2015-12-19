class Conway {
  private cornersOn;
  isAlive = cell => cell === '#';
  constructor(cornersOn) {
    this.cornersOn = cornersOn;
  }

  isCorner(x, y, grid):boolean {
    let [row, width] = [grid.length - 1, grid[0].length - 1];
    return (x === 0 || x === row) && (y === 0 || y === width);
  }

  getNeighbors (x, y, grid):string[] {
      let prevRow = grid[x - 1];
      let nextRow = grid[x + 1];
      let neighbors = [
        grid[x][y - 1],
        grid[x][y + 1]
      ];
      if (prevRow) {
        neighbors = neighbors.concat(prevRow[y - 1], prevRow[y], prevRow[y + 1]);
      }
      if (nextRow) {
        neighbors = neighbors.concat(nextRow[y - 1], nextRow[y], nextRow[y + 1]);
      }
      return neighbors;
  }
  step (grid:string[][], stepsRemaining):number {
    if (stepsRemaining === 0) {
      return grid.reduce((prev, curr) => prev + curr.filter(this.isAlive).length, 0);
    }
    let newGrid = grid.map((row, rowNum, arr) => {
      return row.map((cell, cellNum) => {
        if (this.isCorner(rowNum, cellNum, grid) && this.cornersOn) {
          return '#';
        }
        let neighbors = this.getNeighbors(rowNum, cellNum, grid);
        let neighborCount = neighbors.filter(this.isAlive).length;
        if (this.isAlive(cell)) {
          if (neighborCount === 2 || neighborCount === 3) {
            return '#';
          } else {
            return '.';
          }
        } else {
          return neighborCount === 3 ? '#' : '.';
        }
      });
    });
    return this.step(newGrid, stepsRemaining - 1);
  }
}


export let stepTimes = (rows:string[], timesToStep:number):number => {
  let grid:string[][] = rows.map(row => row.split(''));
  return new Conway(false).step(grid, timesToStep);
};
export let stepsWithCornersStuck = (rows:string[], timesToStep:number):number => {
  let grid:string[][] = rows.map(row => row.split(''));
  return new Conway(true).step(grid, timesToStep);
};
