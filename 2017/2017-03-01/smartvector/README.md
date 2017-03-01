# Exercise

Write a "class" Vector, that stores up to N integers. It should have a method
Set for setting values and Get for reading values at a certain index.
Given that adjacent items are commonly repeated, design Vector such that it
consumes as little memory as possible.

Example:

```
v = new Vector(N)
v.Set(0, 7)
assert v.Get(0) == 7
```

A typical input would be something like:

```
7 7 7 7 7 7 7 -1 -1 8 6 5 5 5 5 5 5 5
```

We can think that we are reading values from a temperature sensor every second,
and the temperature changes gradually, so that most adjacent readings have the
same value.
