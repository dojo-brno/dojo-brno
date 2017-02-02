# Coding Dojo session step-by-step

This section roughly describes how we setup our Coding Dojo sessions. It shall
not be considered as a prescription, just a collection of ideas if anyone may
need them.


## 1. Setup the environment

### Programming language

Coding dojos are language agnostic. Pick your favorite language and make sure
you have installed its compiler/interpreter and you have a testing framework, if
the language doesn't come with one.

#### Install Go

We've been coding in [Go](https://golang.org) quite often. Follow the
[installation instructions](https://golang.org/doc/install) for your operating
system and make sure to setup the `GOPATH` environment variable and [test your installation](https://golang.org/doc/install#testing).

Now you can use the `go` tool to download this repository to the appropriate
location under your `GOPATH` with this command:

```
go get -d github.com/dojo-brno/dojo-brno
```

The repository can now be found in `$GOPATH/src/github.com/dojo-brno/dojo-brno`.

#### Install Python

If you are using Linux or macOS, you already have Python installed.
For more guidance, follow the official documentation: https://www.python.org/about/gettingstarted/.

### Tools

We use some tools to help us track test results, time, etc.

*All tools are optional.*

#### Install redgreen

Since we practice TDD, it is helpful to have some form of visual feedback on the
current state of tests.

When not using [dojotimer](https://github.com/juanplopes/dojotimer), we've used
the [redgreen](https://github.com/rhcarvalho/redgreen) command line program. To
install it, run:

```
go get github.com/rhcarvalho/redgreen
```

Instructions on how we use it will follow in another subsection.

---

Note: if you haven't done so when installing Go, add `$GOPATH/bin` to your
`PATH`, so that you can easily run installed Go programs, like for example
`redgreen`. See https://golang.org/doc/code.html#GOPATH.

---

#### Install DojoTimer

DojoTimer is a program created by one of the [Dojo Rio](https://dojorio.org/)
members. It's the *Swiss Army knife* of coding dojoers.

It is written in C#, and can run on Linux using Mono.

Download a pre-compiled [`dojotimer.exe`](http://www.juanlopes.net/dojotimer/),
then make sure you have the Mono runtime and the *gtk-sharp2* library
installed.

Run `mono dojotimer.exe`, or use a [launcher
script](https://gist.github.com/rhcarvalho/4ad4f90b6d60b96dc225183e773fbf9e).

DojoTimer can track the time, show red/green status, show tests results, save a
participant list, automate actions, and even control a physical semaphore
plugged in a serial port!

## 2. Pick a kata (programming task)

While there are many resources out there, we've been mostly working on tasks
from:

* http://codingdojo.org/cgi-bin/index.pl?KataCatalogue

Note: if you have problems accessing it (unfortunately, sometimes the server
is down), you can use the [Way Back Machine](https://web.archive.org/web/http://codingdojo.org/cgi-bin/index.pl?KataCatalogue).

## 3. Create a working directory

This repository is organized by date and kata name. To help verify that you have
a working Go environment and facilitate following the convention, there's a "Go
script" to create the directory structure for a new session/kata.

Run it like this:

```
go run new.go fizzbuzz 'http://codingdojo.org/cgi-bin/index.pl?KataFizzBuzz'
```

Replace `fizzbuzz` with a package name that refers to the kata that will be
worked on. A directory hierarchy based on the current date and the kata name
will be created for you. Avoid using spaces, dashes, underscores and
uppercase letters in the name, for it might cause trouble with the Go tooling.

The URL argument is optional, and is used to populate a `README.md` file in the
new working directory.


## 4. Start `redgreen` or DojoTimer

*This step is optional and requires `redgreen` or DojoTimer to be installed,
according to the [instructions above](#tools).*

### `redgreen`

We use `redgreen` with [`tmux`](https://tmux.github.io/) to have two terminal
panes, one running `redgreen`, and another free for us to do whatever we may
need, from running tests to consulting documentation with `godoc`.

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

### DojoTimer

Run `dojotimer.exe` (Windows) or `mono dojotimer.exe` (others). Explore the
buttons in the interface to learn what you can do and to configure DojoTimer.

## 5. Open a text editor

It's time to open your favorite or less favorite editor.

You may have a love relationship with an editor that you use daily and in which
you feel proficient and comfortable. Please keep in mind that not everyone may
be able to use it as well as you do, though. For that reason, try to keep it
simple and use the simplest text editor that may get the job done.

Make sure everybody can easily create and save files, and type in code.


## 7. Hack on

Remember:

* Pair programming with a time-box (we use 5 min)
* Write a test
* See the test fail (red)  
  Take a good look at the error message, and make sure you have the output that
  you expected.
* Write code to make the test pass, and nothing more (green)
* Judiciously refactor  
  Remember that you start refactoring on green, and should finish again on
  green. Use baby steps, keep everybody in the loop, don't be afraid to step
  back if needed.
* Write another test and see it fail (red)
* Repeat...
* Reserve time to do a retrospective at the end of the session. Try to
  understand what went well and what can be improved and how. Discuss topics
  that people are interested in, ask and answer questions, etc.

If you do the retrospective with a whiteboard, you may take a picture of it at
the end, and share it in the mailing-list in the next day.


## 8. Afterwards

After the session is over, add an `AUTHORS` file with the names of the
participants, for future reference. See the existing files.

We also keep a link to the problem source / description in a `README.md` file
next to the implementation for reference.

Both files should already exist if you used the [`new.go`
program](#3-create-a-working-directory).

Commit & push the code, so that we can keep track of our history. If you don't
have push access, contact one of the repository maintainers and you can become
one.

Send an email with a retrospective review. You may foster further discussion
during the week before the next meeting.
