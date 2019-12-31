# Auld Lang

Auld Lang is [I believe] a turing complete, case insensitive language with basic 
syntax based on the lyrics of Auld Lang Syne. Why? I blame this 
[tweet](https://twitter.com/chrisoldwood/status/1211960165496508416).

Influenced by 
[Rho&#8242;&#8242;](https://en.wikipedia.org/wiki/P%E2%80%B2%E2%80%B2), among
[others](https://en.wikipedia.org/wiki/Brainfuck) [NSFW], it's not a blind copy,
due mainly to getting the original "program" to do something.

## Syntax

Auld Lang files contain a series of instructions in the format: 
`instruction[.!?,;]?\n`.

The optional character before the `\n` is the _terminator_. Terminators have 
significance (see Terminators below).

## Memory

Auld Lang is based on a looped `n` cell memory, each holding a 64 bit int. The 
pointer `(ptr)` starts at `0`. The following logic is applied:

```
while (ptr >= n) { ptr = ptr - n }
while (pty < 0) { prt = prt + n }
```

## Arguments

Some instructions in Auld Lang take an integer argument. These always take the
form `instruction (.*?)[.!?,;]?\n`. The argument is the count of characters in 
the capture group
 

## Terminators

Terminators are special characters at the end of an instruction that can perform
a second instruction.

### ?

A `?` in the terminator will read a string from `stdin` and add the character
count of that string to the current memory cell.

In pseudo-code:

```
text = readline()
*ptr = *ptr + len(text)
```

### !

A `!` in the terminator will move the pointer to the right (`ptr++`).

### ;

A `;` in the terminator will move the pointer to the left (`ptr--`).

### .

A `.` in the terminator will decrement the value of the current cell (`*ptr--`).

### ,

A `,` in the terminator will increment the value of the current cell (`*ptr--`).

### Instructions

Auld Lang understands the following 7 instructions. If it does not recognise an
instruction a syntax error is thrown. Blank lines are ignored.

### Happy

The first line in any Auld Lang is the `Happy` instruction which takes an
argument. The argument defines the memory allocation for the program. Zero is an
invalid memory size so an argument must be provided.

For example

```
Happy New Year!
```

Will set the memory size to 8 and move the pointer to cell `1` due to the `!`
terminator.

### Should auld acquaintance be forgot

The instruction `Should auld acquaintance be forgot` will repeat the next 
instruction until the memory cell currently being pointed to is equal to 0.

In pseudo-code:

```
while (*ptr) { <instruction> }
```

### For auld lang syne

The instruction `For auld lang syne` will output the character specified by the
value of the current cell. If an argument is provided the pointer is moved to 
the left by the value of the argument.

### Sin auld lang syne

The instruction `sin auld lang syne` will output the character specified by the
value of the current cell. If an argument is provided the pointer is moved to 
the right by the value of the argument.

### We'll

The `We'll` instruction takes an argument and decrements the value of the 
current cell by the amount specified in the argument.

### And

The `And` instruction takes an argument and increments the value of the current
cell by the amount specified in the argument.

### We

The `We` instruction takes an argument and jumps to the line after the next
`But` instruction if the value of the current cell is greater than the value
passed in the argument. If no `But` exists, the program terminates.

### But

The `But` instruction takes an argument and jumps to the previous `We`
instruction if the value of the current cell is less than the value passed in
the argument. If no `We` exists, the program jumps to the first instruction
after `Happy`.

## Example Program

The following is an example Auld Lang program. I've not yet worked out what it
will do. The language spec _will_ change if this program turns out to not be
valid, or always gets stuck in an infinite loop.

```
Happy New Year!

Should auld acquaintance be forgot,
and never brought to mind?
Should auld acquaintance be forgot,
and auld lang syne?
For auld lang syne, my jo,
for auld lang syne,
we'll tak a cup o' kindness yet,
for auld lang syne.
And surely ye'll be your pint-stowp!
and surely I'll be mine!
And we'll tak a cup o' kindness yet,
for auld lang syne.
We twa hae run about the braes,
and pu'd the gowans fine;
But we've wander'd mony a weary foot,
sin auld lang syne.
We twa hae paidl'd i' the burn,
frae morning sun till dine;
But seas between us braid hae roar'd
sin auld lang syne.
And there's a hand, my trusty fiere!
and gie's a hand o' thine!
And we'll tak a right gude-willy waught,
for auld lang syne.
For auld lang syne, my jo,
for auld lang syne,
we'll tak a cup o' kindness yet,
for auld lang syne.
And surely ye'll be your pint-stowp!
and surely I'll be mine!
And we'll tak a cup o' kindness yet,
for auld lang syne.
```

## Todo 

Write an interpreter and make this puppy work for real!

## License

Public Domain
