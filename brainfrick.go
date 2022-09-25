package brainfrick

import (
	"github.com/mtpedro/brainfrick/lib/to"
	"github.com/mtpedro/brainfrick/lib/from"
)

func FromBrain(code string) string {
	return To.Execute(code);
}

func ToBrain(text string) string {
	return From.Convert(text);
}