import unittest


def convert(number):
    if number == one:
        return 1
    if number == two:
        return 2
    if number == three:
        return 3
    if number == four:
        return 4
    return 0


class TestBankOCR(unittest.TestCase):
    def test_convert0(self):
        self.assertEqual(convert(zero), 0)

    def test_convert1(self):
        self.assertEqual(convert(one), 1)

    def test_convert2(self):
        self.assertEqual(convert(two), 2)

    def test_convert3(self):
        self.assertEqual(convert(three), 3)

    def test_convert4(self):
        self.assertEqual(convert(four), 4)

zeroAndOne =


zero = """\
 _
| |
|_|"""

one = """\

  |
  |"""

two = """\
 _
 _|
|_ """

three = """\
 _
 _|
 _|"""

four = """\

|_|
  |"""


dict = {zero : 0, one : 1, two : 2, three : 3, four : 4}
