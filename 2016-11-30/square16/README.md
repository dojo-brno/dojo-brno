# Exercise

This time we played with a problem brought by Pavel. We had no formal
specification.

The idea is as follows: given a 4x4 grid, find a way to divide the grid into two
parts, as if done by drawing a polygonal line passing through the borders of the
grid cells, such that each part has perimeter 16.

For example, splitting the grid in half is not a solution, because it divides
the original square into two equal rectangles of perimeter 12:

    +-+-+-+-+
    |X|X|X|X|
    +-------+
    |X|X|X|X|
    +-------+
    | | | | |
    +-------+
    | | | | |
    +-+-+-+-+

This more complex shape is a solution:

    +-+-+-+-+
    |X| | | |
    +-------+
    |X| | | |
    +-------+
    |X| |X| |
    +-------+
    |X|X|X| |
    +-+-+-+-+

The polygon in the part marked with X has perimeter:

    4 + 1 + 3 + 1 + 1 + 1 + 2 + 3 = 16

And the polygon in the empty/upper part has perimeter:

    3 + 3 + 4 + 1 + 2 + 1 + 1 + 1 = 16


In order to start coding with tests, we rephrased the problem as: given a
potential solution to the grid problem, write a function that returns whether
the input is a valid solution.

Later, we could write code to generate all potential solutions, and filter them
using the aforementioned function.
