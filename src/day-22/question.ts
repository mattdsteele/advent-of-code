import { leastManaSpent } from './';

let startPlayer = {
  hp: 50,
  mana: 500,
  armor: 0,
  effects: [],
  manaSpent: 0,
  history: []
};
let startBoss = {
  hp: 51,
  damage: 9
};

let silver = leastManaSpent(startPlayer, startBoss);
console.log(`silver: ${silver}`);

let gold = leastManaSpent(startPlayer, startBoss, true);
console.log(`gold: ${gold}`);
