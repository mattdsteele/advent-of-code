let goldPresentsFor = (houseToCheck:number):number => {
    let totalPresents = 0;
    let items = [];
    for (let i = 0; i < houseToCheck; i++) {
      items.push(i);
    }
    return items
      .map((_, i) => i + 1)
      .filter(elf => houseToCheck / elf < 50)
      .filter(elf => houseToCheck % elf === 0)
      .reduce((prev, curr) => prev + (curr * 11), 0);
};

let presentsFor = (houseToCheck:number):number => {
    let totalPresents = 0;
    let items = [];
    for (let i = 0; i < houseToCheck; i++) {
      items.push(i);
    }
    return items
      .map((_, i) => i + 1)
      .filter(elf => houseToCheck % elf === 0)
      .reduce((prev, curr) => prev + (curr * 10), 0);
};

let calculate = (minimum, algorithm):number => {
  let houseNumber = 0;
  while (true) {
    let presents = algorithm(houseNumber);
    console.log(houseNumber, presents);
    if (presents >= minimum) {
      return houseNumber;
    }
    houseNumber++;
    //houseNumber += 420; this is faster but breaks the tests
  }
};

export let minimumPresents = (minPresents:number):number => {
  return calculate(minPresents, presentsFor);
};

export let only50 = (minPresents:number):number => {
  return calculate(minPresents, goldPresentsFor);
}
