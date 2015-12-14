import { nextPassword } from './';

let currentPassword = 'cqjxjnds';

let silverPassword = nextPassword(currentPassword);
console.log(`silver: ${silverPassword}`);
console.log(`gold: ${nextPassword(silverPassword)}`);
