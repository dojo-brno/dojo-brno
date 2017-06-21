# Exercise

This week Boris gave us a lesson on how the Java type system works, including
_Generics_, how it applies to _Collections_ and including things like `<?
extends ...>` and `<? super ...>`.

We didn't produce any code.

There were no slides and nothing could possibly reproduce the live coding and
explanations that Boris gave. As a reference, here's a Wikipedia page about
the [Java syntax](https://en.wikipedia.org/wiki/Java_syntax).

After all the talking and revisiting [last week's
code](../../2017-06-14/chooseexportlocation/boris/org/bbelovic/kata/exportlocation/ExportLocationSelector.java),
we finally had a better understanding of how the neat few lines of code work. As
a comparison, we looked at how [sorting in Go is not
generic](https://golang.org/pkg/sort/#example__sortKeys), requiring a lot more
boilerplate to sort a list of values by a certain attribute. After that, Rodolfo
showed a bit of Python in a live interpreter session and how sorting works in
that language, in particular how one can sort a list containing arbitrary types
without any extra code.
