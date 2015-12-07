/// <reference path="../../typings/node/node.d.ts" />

import { createHash } from 'crypto';

export function adventCoin(prefix: string, numberZeros = 5): number {
  let test = 0;
  while(true) {
    let input = `${prefix}${test}`;
    let hash = createHash('md5').update(input).digest('hex');
    let val = hash.substring(0, numberZeros);
    if (val === new Array(numberZeros + 1).join('0')) {
      return test;
    }
    test++;
  }
}
