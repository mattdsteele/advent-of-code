import BathroomCode from './';
import { expect } from 'chai';
import { readFileSync } from 'fs';

const input = `ULL
RRDDD
LURDL
UUUUD
`;
const realInput = readFileSync('./src/2/input.txt').toString();
describe('day 2', () => {
  const bc = new BathroomCode();
  describe('behaviors', () => {
    it('goes u/d/l/r', () => {
      expect(bc.direction('U', '5')).to.eql('2');
      expect(bc.direction('D', '5')).to.eql('8');
      expect(bc.direction('L', '5')).to.eql('4');
      expect(bc.direction('R', '5')).to.eql('6');
    });
    it('hits walls', () => {
      expect(bc.direction('U', '1')).to.eql('1');
      expect(bc.direction('D', '7')).to.eql('7');
      expect(bc.direction('L', '4')).to.eql('4');
      expect(bc.direction('R', '9')).to.eql('9');
    });
    it('runs a spectrum', () => {
      expect(bc.line('ULL', '5')).to.eql('1');
      expect(bc.line('RRDDD', '1')).to.eql('9');
      expect(bc.line('LURDL', '9')).to.eql('8');
      expect(bc.line('UUUUD', '8')).to.eql('5');
    });

    it('produces a code', () => {
      expect(bc.codeFor(input)).to.eql('1985');
    });
  });

  describe('silver', () => {
    it('produces silver', () => {
      console.log(bc.codeFor(realInput));
    });
  });
});

describe('day 2 gold', () => {
  const bc = new BathroomCode(true);
  it('can calculate gold', () => {
    expect(bc.codeFor(input)).to.eql('5DB3');
  });
  it('produces gold', () => {
    console.log(bc.codeFor(realInput));
  });
});
