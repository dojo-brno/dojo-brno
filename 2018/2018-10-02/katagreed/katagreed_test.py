class Greed:
    def score(dice):
        sorted_list = list(sorted(dice))
        if sorted_list == [1,2,3,4,5,6]:
            return 1200
        if Greed.__is_three_pairs(sorted_list):
            return 800
        values_of_triples = [1000,200,300,400,500,600]
        result = 0
        for i in range(6):
            count = dice.count(i+1)
            if count > 2:
                result += 2**(count-3) * values_of_triples[i]
        if dice.count(1) is 1:
            result += 100
        if dice.count(5) is 1:
            result += 50
        return result

    def __is_three_pairs(sorted_list):
        return ( len(sorted_list) == 6
            and sorted_list[0] == sorted_list[1]
            and sorted_list[2] == sorted_list[3]
            and sorted_list[4] == sorted_list[5]
            and sorted_list[1] != sorted_list[2]
            and sorted_list[3] != sorted_list[4])


def test_empty():
    assert Greed.score([]) == 0

def test_one():
    assert Greed.score([1]) == 100

def test_one_and_two():
    assert Greed.score([1,2]) == 100

def test_five_and_two():
    assert Greed.score([5,2]) == 50

def test_straight():
    assert Greed.score([1,3,5,2,6,4]) == 1200

def test_triple_ones():
    assert Greed.score([1,1,1]) == 1000

def test_four_ones():
    assert Greed.score([1,1,1,1]) == 2000

def test_five_ones():
    assert Greed.score([1,1,1,1,1]) == 4000

def test_six_ones():
    assert Greed.score([1,1,1,1,1,1]) == 8000

def test_triple_twos():
    assert Greed.score([2,2,2]) == 200

def test_four_twos():
    assert Greed.score([2,2,2,2]) == 400

def test_five_twos():
    assert Greed.score([2,2,2,2,2]) == 800

def test_six_twos():
    assert Greed.score([2,2,2,2,2,2]) == 1600

def test_triple_threes():
    assert Greed.score([3,3,3]) == 300

def test_four_threes():
    assert Greed.score([3,3,3,3]) == 600

def test_five_threes():
    assert Greed.score([3,3,3,3,3]) == 1200

def test_six_threes():
    assert Greed.score([3,3,3,3,3,3]) == 2400

def test_two_triples():
    assert Greed.score([1,1,1,2,2,2]) == 1200

def test_two_triples_ones_fives():
    assert Greed.score([1,1,1,5,5,5]) == 1500

def test_one_and_five():
    assert Greed.score([1,5]) == 150

def test_one_and_five():
    assert Greed.score([1,1,5,5]) == 0

def test_pairs():
    assert Greed.score([1,1,3,3,6,6]) == 800
