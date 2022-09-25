# Golang Brainfuck

this is a golang module for converting text to and from brainfuck. 

## Setup

To add this to your project, run:
```
go get https://github.com/mtpedro/brainfrick
```
Note that you have to first initialize you go project:
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

import "https://github.com/mtpedro/brainfrick"

...
```

## How to use

To convert to brainfrick, you can call:
```
brainfrick.FromBrain("+[----->+++<]>+.");
```
this will return the ascii letter H.

you may also convert strings to brainfuck code with:
```
brainfrick.ToBrain("H");
```
this will return the following brainfuck code:
```
>++++++++++[<+++++++>-]++.
```

# THIS PROJECT IS UNFINISHED.