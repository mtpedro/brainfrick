# Brainfrick - Brainfuck in Go.

this is a golang module for converting text to and from brainfuck. 

## Setup

To add this to your project, run:
```
go get github.com/mtpedro/brainfrick
```
Note that you have to first initialize your golang project:
```
go mod init <your project home>
```
Next, run:
```
go mod tidy
```

just import the module in your project and you're good to go!
```
//main.go
package main

import "github.com/mtpedro/brainfrick"

...
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