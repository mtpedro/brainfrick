package From

import "fmt"

func Convert(text string) string {
	var Brainfrick string;	
	var NumericValues = []byte(text);
	var History []byte;

	outer: for _, v := range NumericValues {

		for k, b := range History {
			if v == b {
				goback := len(History) - k;

				for i:=0;i<int(goback);i++ { Brainfrick = fmt.Sprintf("%v<", Brainfrick) }
				
				Brainfrick = fmt.Sprintf("%v.>", Brainfrick);

				for i:=0;i<int(goback);i++ { Brainfrick = fmt.Sprintf("%v>", Brainfrick) }

				History = append(History, byte(255));

				continue outer;
			}
		}

		tail := v % 10;
		arm := (v - tail)/10;

		Brainfrick = fmt.Sprintf("%v>++++++++++[<", Brainfrick)

		for i:=0;i<int(arm);i++{ Brainfrick = fmt.Sprintf("%v+", Brainfrick) }
		Brainfrick = fmt.Sprintf("%v>-]<", Brainfrick);
		for i:=0;i<int(tail);i++{ Brainfrick = fmt.Sprintf("%v+", Brainfrick) }
		Brainfrick = fmt.Sprintf("%v.>", Brainfrick);

		History = append(History, v);
	}

	return Brainfrick;
}