import { adventCoin } from './';

let prefix = 'ckczppom';
let result = adventCoin(prefix);
console.log(`silver: ${result}`);

let goldResult = adventCoin(prefix, 6);
console.log(`gold: ${goldResult}`);
