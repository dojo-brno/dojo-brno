# Exercise

We did an unusual exercise of reverse-engineering a given binary.

We got an archive `dcChallenge_01.tar.gz` provided by Raphael, containing an
ELF 32-bit binary, and used some Linux tools to understand what it does and how
we could derive C source code for implementing a similar behavior.

Raphael gave us an overview of the virtual memory layout for the 32-bit Intel
architecture, among other things, and we got some hands-on experience with
(at least) these tools:

- file
- ldd
- strace
- ltrace
- objdump
- gdb

Mid-way, we've also decided to complement the learning session with tmux and
Vim...

## Reproducing the experiment

Start with extracting the binary:

```
tar xzvf dcChallenge_01.tar.gz
```

Then, use the tools listed above to inspect the file `dcChallenge_01`. In
particular, using `gdb` you can gain insight from disassembling the `main`
function:

```console
$ gdb dcChallenge_01
(gdb) disassemble main
Dump of assembler code for function main:
   0x08048384 <+0>:     lea    0x4(%esp),%ecx
   0x08048388 <+4>:     and    $0xfffffff0,%esp
   0x0804838b <+7>:     pushl  -0x4(%ecx)
   0x0804838e <+10>:    push   %ebp
   0x0804838f <+11>:    mov    %esp,%ebp
   0x08048391 <+13>:    push   %edi
   0x08048392 <+14>:    push   %esi
   0x08048393 <+15>:    push   %ebx
   0x08048394 <+16>:    push   %ecx
...
```

**Warning**: I may be writing completely wrong and stupid things here....

What we learned is that that first part is the prolog of the function call,
setting up some state in the stack.

Then, we gained insight about how many bytes are used by local variables,
because of how much gets subtracted from `%esp` (after it has been reset to
`%ebp` in offset main+11):

```
...
   0x08048395 <+17>:    sub    $0x88,%esp
...
```

That's 0x88 == 136 bytes.

And we saw a pointer dereference / local variable access, because of the
negative indexing on `%ebp`:

```
...
  0x0804839b <+23>:    lea    -0x31(%ebp),%eax
...
```

And we saw a `test` and `je`, test and jump that suggests an `if` in C code.

Down further, we saw more fiddling with local variables and a call to `puts`,
closer to where the function returns.

Finishing a reverse-engineered `main.c` is left as an exercise to the reader :-)
