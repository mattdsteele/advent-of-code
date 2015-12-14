import { lookAndSay } from './';

let input = '3113322113';

for (let x = 0; x < 40; x++) {
  console.log(x);
  input = lookAndSay(input);
}

console.log(`silver: ${input.split('').length}`);
