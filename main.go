package brainfrick

import (
	"github.com/mtpedro/brainfrick/lib/to"
)

func FromBrain(code string) string {
	return To.Execute(code);
}

func ToBrain(text string) string {
	return "hi";
}