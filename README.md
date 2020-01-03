# Auld Lang

Auld Lang is [I believe] a turing complete, case insensitive language with basic 
syntax based on the lyrics of Auld Lang Syne. Why? I blame 
[Chris Oldwood](https://twitter.com/chrisoldwood) and this 
[tweet](https://twitter.com/chrisoldwood/status/1211960165496508416).

Influenced by 
[Rho&#8242;&#8242;](https://en.wikipedia.org/wiki/P%E2%80%B2%E2%80%B2), among
[others](https://en.wikipedia.org/wiki/Brainfuck) [NSFW], it's not a blind copy,
due mainly to getting the original "program" to do something. Notably is 
supports arguments, and it's looping structure is very different.

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

A `?` in the terminator will read a string from `stdin` and subtract the
character count of that string to the current memory cell. Then move the `ptr`
to the right.

In pseudo-code:

```
text = readline()
*ptr = *ptr - len(text)
ptr++
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

Auld Lang understands the following 10 instructions. If it does not recognise an
instruction a syntax error is thrown. Blank lines are ignored.

### Happy

The `Happy` allocates memory for the program. It takes an argument which defines
the size of the memory in cells. Zero is an invalid memory size so an argument 
must be provided. `Happy` is normally the first line in any Auld script as, by
default, the defined memory only 1 cell in size. `Happy` can be called at any
time and is a destructive operation, throwing away the previous memory contents
and setting the `ptr` back to 0.

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
value of the current cell. The exact character will be `abs(value) mod 127`.
Characters 0-31 will be rendered as a `?`. If an argument is provided the
pointer is moved to the left by the value of the argument after the character is
output.

### Sin auld lang syne

The instruction `sin auld lang syne` will output the character specified by the
value of the current cell. The exact character will be `abs(value) mod 127`.
Characters 0-31 will be rendered as a `?`. If an argument is provided the
pointer is moved to the right by the value of the argument after the character is
output.

### We'll

The `We'll` instruction takes an argument and decrements the value of the 
current cell by the amount specified in the argument.

### And

The `And` instruction takes an argument and increments the value of the current
cell by the amount specified in the argument.

### Frea

The `Frea` instruction moves the pointer to the right by the value specified in
the argument.

### We

The `We` instruction takes an argument and jumps to the line after the next
`But` instruction if the value of the current cell is less than the value
passed in the argument. If no `But` exists, the program terminates.

In pseudo-code:

```
while (*ptr >= cell) { ... }
```

### But

The `But` instruction takes an argument and jumps to the previous `We`
instruction if the value of the current cell is less than the value passed in
the argument. If no `We` exists, the program jumps to the first instruction
after `Happy`.

```
do { ... } while (*ptr >= cell)
```

### Kevlin

The `Kevlin` keyword turns on debug.

Kevlin takes no arguments (and no prisoners). Invoking Kevlin will dump the
contents of memory after every operation. Kevlin should never be invoked in
production, although there are bigger issues if this is being used in production
code.

## Example Program

The following is the original Auld Lang program. It is unknown if there is a
canonical "Happy" line, and sadly if there is that has been lost to time. 

```
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

When this is run as is the program requires a 15 character input to continue and 
then outputs: `[1][2][21]<QQ('([17]b`

This appears to be sending an empty header, then a NAK (possibly indicating
it is not ready to reply to a poll). The string `<QQ('(` is then output before
Device Control 1 is selected and the character `b` sent. One possible conclusion
is the program is meant to be run on a machine connected to a network and other
devices.

That said, without a "Happy" line terminating in `.` or `,` the line:

```
and never brought to mind?
```

Is never run which would imply a "Happy" line is required.

## Hello, World!

That's lovely, can we make it do real work? Can we do "Hello, World!"? Why yes,
yes we can!

```
Happy 2020 Everyone

And so it is that we can sing
And keep the tune the same
And sort of because it's hard!
And bugger that for a game of soldiers, it's late I want to get on with some other stuff and I need very !
And specific sentence lengths, and ooh, look at that, now I've thrown grammar out of the window too, this really!
And is cheating. Honestly, if I needed to, at this point I'd just split words over lines so that I could get the!
And correct line lengths, but as it is I've only needed to sneak in a wayward space, however I'm bored now and this!
And gets a lot easier if I just use spaces      !
And                                 !
And I'll do a better version later, I promise, but for now, spaces!                        !
And                                                                                                                !
And                                                                                                                   !
And                                                                                                             !
And                                                                                                     !
And then we can just wrap the buffer!!
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
sin auld lang syne o
```

```
Hello, World!
```

## Auld Lang Sine

Of course we want to do something a lot more interesting that that, so here is
a program that will output: 

```
        an        
      l    g      
a    d      s    e
  ul          in  
```

```
Happy New Year To You All!

Frae here we need to be
And load data!
And load this cell with ASCII spaces
We loop o'er the rest
Frae o
And load the cells with ASCII spaces
But we exit when we get to our control cell
We'll drop the data located in this memory cell!

Frae to here!
And once here we need to bump this cell up
And so it is now the letter "a"!
And similarly this cell needs to be bumped up even more
And so it is now the letter "n"!
Frae to there!

Should auld acquaintance be forgot
Sin auld lang syne o
And print "\n"
Sin auld lang syne o

Happy New Year To You All!
Frae here we need to be
And load data!
And load this cell with ASCII spaces
We loop o'er the rest
Frae o
And load the cells with ASCII spaces
But we exit when we get to our control cell
We'll drop the data located in this memory cell!

Frae there!
And we need to add a load more numbers to this
And so it is now become the letter "l"!
Frae here
And make this much bigger than it was before
And so now making it the letter "g"!
Frae hither!

Should auld acquaintance be forgot
Sin auld lang syne o
And print "\n"
Sin auld lang syne o

Happy New Year To You All!
Frae here we need to be
And load data!
And load this cell with ASCII spaces
We loop o'er the rest
Frae o
And load the cells with ASCII spaces
But we exit when we get to our control cell
We'll drop the data located in this memory cell!

And this new cell needs to get incremented
And so it is now the letter "a"
Frae here!
And we need to add a load more to this
And so it is now become the letter "d"!
Frae herein
And this needs to get really large
And I mean really really large
And so it is now the letter "s"
Frae here!
And make this here a tiny bit bigger than before
And so it is now the letter e
Frae so

Should auld acquaintance be forgot
Sin auld lang syne o
And print "\n"
Sin auld lang syne o

Happy New Year To You All!
Frae here we need to be
And load data!
And load this cell with ASCII spaces
We loop o'er the rest
Frae o
And load the cells with ASCII spaces
But we exit when we get to our control cell
We'll drop the data located in this memory cell

Frae so!
And by putting a bunch of stuff into this cell
And incrementing lots
And so it becomes the letter u!
And this cell needed to become the letter "l"
And which was done before if you recall
Frae over there!
And let's make this the letter "i" indeed
And we load in the stuff needed for that!
And finally this needs to be the letter "n"
And I've also seen that before so it's easy!
Frae so!

Should auld acquaintance be forgot
Sin auld lang syne o
And print "\n"
Sin auld lang syne o
```

## License

Public Domain
