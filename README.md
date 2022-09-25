# Brainfrick - Brainfuck in Go.

this is a golang module for converting text to and from brainfuck. 

## Setup

To add this to your project, run:
```
go get github.com/mtpedro/brainfrick
```
just import the module in your project and you're good to go!
```
import "github.com/mtpedro/brainfrick"
```

## How to use

To convert brainfuck to text, you can call:
```
brainfrick.FromBrain("+[----->+++<]>+.");
```
this will return the ascii letter H.

you may also convert text to brainfuck code with:
```
brainfrick.ToBrain("H");
```
this will return the following brainfuck code:
```
>++++++++++[<+++++++>-]++.
```

Have fun!