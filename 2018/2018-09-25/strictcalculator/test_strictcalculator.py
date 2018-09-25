import pytest
import string as s

def StringCalculator(numbers):
    sum = 0
    for i in numbers.split(","):
        if not i.strip():
            continue
        sum += float(i)
    return sum

def StringCalculator2(numbers):
    return (sum(float(i) for i in numbers.split(",")))

# def test_StringCalculator():
#     assert StringCalculator("1,2") == 3
#     assert StringCalculator("1,1") == 2
#     assert StringCalculator("1,0") == 1
#     assert StringCalculator("1,2,3") == 6
#     assert StringCalculator("1,2,3,4") == 10
#     assert StringCalculator("1.1, 2") == 3.1


@pytest.mark.parametrize(
    ['vyraz','vysledek'],
    [("1,2", 3),
    ("1,1",2),
    ("1,0",1),
    ("1,2,3",6),
    ("1,2,3,4",10),
    ("1.1,2",3.1),
    ("",0),
    ("1.1,",1.1),
    ("1.1, ",1.1),
    ]
)
def test_StringCalculator(vyraz, vysledek):
    assert StringCalculator(vyraz) == vysledek

@pytest.mark.parametrize(
    ['vyraz','vysledek'],
    [("1,2", 3),
    ("1,1",2),
    ("1,0",1),
    ("1,2,3",6),
    ("1,2,3,4",10),
    ("1.1,2",3.1),
    # ("",0),
    # ("1.1,",1.1),
    # ("1.1, ",1.1),
    ]
)
def test_StringCalculator(vyraz, vysledek):
    assert StringCalculator2(vyraz) == vysledek
