# Coding Dojo Brno

This repository is the home of the source code produced during the Coding Dojo
Brno sessions.

The code has no value in itself, and it should be considered simply as a log of
our activities.


## Mailing List

We discuss before and after sessions via a [mailing
list](https://groups.google.com/forum/#!forum/dojo-brno).


## Coding Dojo session step-by-step

This section describes how we usually setup a Coding Dojo session.


### 1. Setup programming language environment

We've been coding in Go most of the time. Follow the [install
instructions](https://golang.org/doc/install) and make sure to setup the
`GOPATH` environment variable and clone this repository under
`$GOPATH/src/github.com/dojo-brno/dojo-brno`.

If you have a single path in `GOPATH`, downloading the repository to the
appropriate location can be achieved by the command:

```
go get github.com/dojo-brno/dojo-brno
```

It's also useful to have `$GOPATH/bin` in your `PATH`, so that you can easily
run installed Go programs.


### 2. Install `redgreen`

This step is optional.

Since we'll be doing TDD, it is helpful to have some form of visual feedback on
the current state of tests.

We should really be using [dojotimer](https://github.com/juanplopes/dojotimer),
but for now we've been using the
[redgreen](https://github.com/rhcarvalho/redgreen) Go program. To install it,
run:

```
go get github.com/rhcarvalho/redgreen
```

Instructions on how we use it will follow in another subsection.


### 3. Pick a kata (programming task)

While there are many resources out there, we've been mostly working on tasks
from:

* http://codingdojo.org/cgi-bin/index.pl?KataCatalogue


### 4. Create working directory

This repository is organized by date and kata name. To help verify that you have
a working Go environment and facilitate following the convention, there's a "Go
script" to create the directory structure for a new session/kata.

Run it like this:

```console
go run new.go fizzbuzz 'http://codingdojo.org/cgi-bin/index.pl?KataFizzBuzz'
```

Replace `fizzbuzz` with a package name that refers to the kata that will be
worked on. The top-level directory with the current date will be created
automatically. Avoid using spaces, dashes, underscores and uppercase letters in
the name, for it might cause trouble with the Go tooling.

The URL argument is optional, and is used to populate a `README.md` file in the
new package directory.


### 5. Start `redgreen`

This step is optional. It requires `redgreen` to be installed, according to the
previous instructions.

We use `redgreen` with `tmux` to have two terminal panes, one running
`redgreen`, and another free for us to do whatever we may need, from running
tests to consulting documentation with `godoc`.

For example, the command bellow will create a new session, turn off the status
bar (to save space in the screen), split the window horizontally giving 20% of
the terminal width to `redgreen`

```
tmux new-session \; set status off \; split-window -hp 20 redgreen \; last-pane
```

Make sure that you run `tmux` and `redgreen` from the working directory of the
programing task, so that tests are run correctly.

To have the red/green state always visible, we set the window to be "Always on
Top" (on GNOME).


### 6. Open your favorite or less favorite editor

You may have your preferred editor in which you're proficient at. Keep in mind
that not everyone may be able to use it as well as you do, though. For that
reason, try to keep it simple and use the simplest text editor that may get the
job done.

Make sure everybody can easily create and save files, and type in code.


### 7. Hack on

Remember:

* Pair programming with a time-box (5 min)
* Write a test
* See the test fail (red)
* Write code to make the test pass, and nothing more (green)
* Maybe refactor
* Write another test and see it fail (red)
* ...
* At the end of the session, arrange time to do a retrospective. Try to
  understand what went well and what can be improved and how. Discuss topics
  that people are interested in, ask and answer questions, etc.

If you do the retrospective with a whiteboard, you may take a picture of it at
the end, and share it in the mailing-list in the next day.


### 8. Afterwards

After the session is over, add an `AUTHORS` file with the names of the
participants, for future reference. See the existing files.

We also keep a link to the problem source / description in a `README.md` file
next to the implementation for reference.

Commit & push the code, so that we can keep track of our history.

Send an email with a retrospective review, you may foster further discussion
during the week before the next meeting.
