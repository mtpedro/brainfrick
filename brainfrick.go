package brainfrick

import (
	"github.com/mtpedro/brainfrick/pkg/to"
	"github.com/mtpedro/brainfrick/pkg/from"
)

func FromBrain(code string) string {
	return To.Execute(code);
}

func ToBrain(text string) string {
	return From.Convert(text);
}