package To

import (
    "fmt"
)

type Program struct {
    size int
    instructions []byte
    at int
}

func check(e error) {
    if e != nil {
        fmt.Println(e);
        panic(e);
    }
}

func Execute(code string) string {
    var program = new(Program);
    program.size = 30000;
    program.instructions = make([]byte, program.size, program.size);
    program.at = 0;

	return executeWith(program, code);
}

func executeWith(program *Program, code string) string {
    var loopStart = -1;
    var loopEnd = -1;
    var ignore = 0;
    var skipClosingLoop = 0;
	var text string;

    for pos, char := range code {
        if ignore == 1 {
            if char == '[' {
                skipClosingLoop += 1;
            } else if char == ']' {
                if skipClosingLoop != 0 {
                    skipClosingLoop -= 1;
                    continue;
                }

                loopEnd = pos;
                ignore = 0;
                if loopStart == loopEnd {
                    loopStart = -1;
                    loopEnd = -1;
                    continue;
                }
                loop := code[loopStart:loopEnd];
                for program.instructions[program.at] > 0 {
                    executeWith(program, loop);
                }
            }
            continue;
        }
        switch char {
            case '+':
                program.instructions[program.at] += 1;
            case '-':
                program.instructions[program.at] -= 1;
            case '>':
                if program.at == program.size-1 {
                    program.at = 0;
                } else {
                    program.at += 1;
                }
            case '<':
                if program.at == 0 {
                    program.at = program.size - 1;
                } else {
                    program.at -= 1;
                }
            case '.':
                text = fmt.Sprintf("%v%c", text, rune(program.instructions[program.at]));
            case '[':
                loopStart = pos + 1;
                ignore = 1;
        }
    }
	return text;
}