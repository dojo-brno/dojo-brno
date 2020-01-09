import { expect } from 'chai';

function f(i:number) : number[] {
    if (i == 1) return [];

    for(let j:number = 2; j < i; j++) {
        if ( i % j == 0) {
            return [j].concat(f(i/j));
        }
    }

    return [i];
}

describe('Kata XXX', () => {
    it('Function returns [] for 1', () => {
        expect(f(1)).to.eql([])
    });

    it('Function returns [2] for 2', () => {
        expect(f(2)).to.eql([2])
    });

    it('Function returns [3] for 3', () => {
        expect(f(3)).to.eql([3])
    });

    it('Function returns [2,2] for 4', () => {
        expect(f(4)).to.eql([2,2])
    });

    it('Function returns [2,3] for 6', () => {
        expect(f(6)).to.eql([2,3])
    });

    it('Function returns [3,3] for 9', () => {
        expect(f(9)).to.eql([3,3])
    });

    it('Function returns [2, 2, 3] for 12', () => {
        expect(f(12)).to.eql([2, 2, 3])
    });

    it('Function returns [3,5] for 15', () => {
        expect(f(15)).to.eql([3,5])
    });

    it('Function returns [7,7] for 49', () => {
        expect(f(49)).to.eql([7,7])
    });

    it('Function returns [5,11,13] for {5*11*13} bignumber :)', () => {
        let primes = [5,11,13]
        expect(f(5*11*13)).to.eql(primes)
    });

});
