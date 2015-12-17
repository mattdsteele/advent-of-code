interface Ingredient {
  name: string;
  capacity: number;
  durability: number;
  flavor: number;
  texture: number;
  calories: number;
}

export let parseLine = (line:string):Ingredient => {
  let regex = /(\w*): capacity ([-\d]*), durability ([-\d]*), flavor ([-\d]*), texture ([-\d]*), calories ([-\d]*)/g;

  let [ _, name, capacity, durability, flavor, texture, calories] = regex.exec(line);
  return {
    name,
    capacity: parseInt(capacity),
    durability: parseInt(durability),
    flavor: parseInt(flavor),
    texture: parseInt(texture),
    calories: parseInt(calories)
  };
};

let addPoints = (currentIngredients, ingredient, numberLeft) => {
    return {
      capacity: (ingredient.capacity * numberLeft) + currentIngredients.capacity,
      durability: (ingredient.durability * numberLeft) + currentIngredients.durability,
      flavor: (ingredient.flavor * numberLeft) + currentIngredients.flavor,
      texture: (ingredient.texture * numberLeft)  + currentIngredients.texture
    };
};

export let scoreFor = (currentIngredients, availableIngredients, numberLeft):number => {
  if (availableIngredients.length === 1) {
    let points = addPoints(currentIngredients, availableIngredients[0], numberLeft);

    if (points.capacity > 0 && points.durability > 0 && points.flavor > 0 && points.texture > 0) {
      return points.capacity *
        points.durability *
        points.flavor *
        points.texture;
    } else {
      return 0;
    }
  }

  let highestScore = 0;
  for (let i = 1; i < numberLeft - availableIngredients.length; i++) {
    let points = addPoints(currentIngredients, availableIngredients[0], i);
    let totalScore = scoreFor(points, availableIngredients.slice(1), numberLeft - i);
    highestScore = Math.max(totalScore, highestScore);
  }
  return highestScore;
};

export let bestScore = (lines:string[]):number => {
  let ingredients = lines.map(l => parseLine(l));
  let max = 100;
  return scoreFor({
    capacity: 0,
    durability: 0,
    flavor: 0,
    texture: 0
  }, ingredients, max);
};
