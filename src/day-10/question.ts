import { lookAndSay } from './';

let input = '3113322113';

for (let x = 0; x < 40; x++) {
  input = lookAndSay(input);
}

console.log(`silver: ${input.split('').length}`);

for (let x = 0; x < 10; x++) {
  input = lookAndSay(input);
}

console.log(`gold: ${input.split('').length}`);
