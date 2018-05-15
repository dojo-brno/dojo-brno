function getFizzBuzz(value) {
  if(isNaN(value) || (value =='')) {
    return -1;
  }
  if(value >= 100) {
    return 'Alert system is overloaded'``
  }
  if(value == 33) {
    return 'Ahoj';
  }
  if(value % 3 == 0 && value % 5 == 0) {
    return 'FizzBuzz';
  }
  if(value % 3 == 0) {
    return 'Fizz';
  };
  if(value % 5 == 0) {
    return 'Buzz'``
  };
  return value;
}

describe('getFizzBuzz', function() {
  it('should return numbers for same regular numbers', function() {
    expect(getFizzBuzz(1)).toBe(1);
    expect(getFizzBuzz(2)).toBe(2);
  });
  it('should return Fizz for number divisible by 3 but not by 5', function() {
    expect(getFizzBuzz(3)).toBe('Fizz');
    expect(getFizzBuzz(6)).toBe('Fizz');
  });
  it('should return Buzz for buzz numbers', function() {
    expect(getFizzBuzz(5)).toBe('Buzz');
  });
  it('should return FizzBuzz for numbers divisible by 3 and also 5', function() {
    expect(getFizzBuzz(15)).toBe('FizzBuzz');
  });
  it('should return Ahoj for number 33', function() {
    expect(getFizzBuzz(33)).toBe('Ahoj');
  });
  it('should return -1 for non number', function() {
    expect(getFizzBuzz('text')).toBe(-1);
    expect(getFizzBuzz('windows')).toBe(-1);
    expect(getFizzBuzz('')).toBe(-1);
  });
  it('should return for numbers 100 and higer - Allert message for system overloaded', function() {
      expect(getFizzBuzz(100)).toBe('Alert system is overloaded');
      expect(getFizzBuzz(101)).toBe('Alert system is overloaded');
  });
});
