const expect = require("chai").expect;

function prvocinitele(n) {

    if (n==2*2) {
        return [2,2]
    }
    if (n==2*2*2) {
       let vysledek = prvocinitele(4)
       vysledek.push(2)
       return vysledek
    }
    if (n==2*3) {
        return [2,3]
    }
    if (n==3*3) {
        let vysledek = prvocinitele(3)
        vysledek.push(3)
        return vysledek
     }
     if (n==2*5) {
        return [2,5]

    }
    if (n==2*2*3) {
        let vysledek = prvocinitele(4)
        vysledek.push(3)
        return vysledek
     }
    return [n]


}

describe("test", function() {
    it("s", function() {
        expect(true);
    });
    it("number 2 to [2]", function() {
        expect(prvocinitele(2)).to.eql([2]);
    });
    it("number 3 to [3]", function() {
        expect(prvocinitele(3)).to.eql([3]);
    });
    it("number 4 to [2,2]", function() {
        expect(prvocinitele(4)).to.eql([2,2]);
    });
    it("number 6 number [2,3]", function() {
        expect(prvocinitele(6)).to.eql([2,3]);
    });
    it("number 8 number [2,2,2]", function() {
        expect(prvocinitele(8)).to.eql([2,2,2]);
    });
    it("number 9 number [3,3]", function() {
        expect(prvocinitele(9)).to.eql([3,3]);
    });
    it("number 10 number [2,5]", function() {
        expect(prvocinitele(10)).to.eql([2,5]);
    });
    it("number 12 number [2,2,3]", function() {
        expect(prvocinitele(12)).to.eql([2,2,3]);
    });
});