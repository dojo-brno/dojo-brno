# Exercise

Write a program that prints the hexadecimal representation of each byte in the
input, like this:

```console
$ echo "Hello, world!" | ./proj1 -x
48656c6c6f0a
```

Later, alter the program to implement other ways of printing:

```console
$ echo "Hello, world! Ahoj svete!" | ./proj1
00000000  48 65 6c 6c 6f 2c 20 77  6f 72 6c 64 21 20 41 68  |Hello, world! Ah|
00000010  6f 6a 20 73 76 65 74 65  21 0a                    |oj svete!.      |
$ echo "Hello, world! Ahoj svete!" | ./proj1 -s 14 -n 5
0000000e  41 68 6f 6a 20                                    |Ahoj            |
$ echo "Hello, world!" | ./proj1 -x
48656c6c6f0a
$ printf 'Hello, world!\0Ahoj svete!\n\0AP\nABCD\n' | ./proj1 -S 3
Hello, world!
Ahoj svete!
ABCD
$ echo "48 65 6c6c6f a" | ./proj1 -r
Hello
```

# About this

This was a different type of meeting, where Honza (@jprukner) and I
(@rhcarvalho) tackled the exercise individually, each of us using a mix of no
particular methodology and TDD, then compared what we've got.
