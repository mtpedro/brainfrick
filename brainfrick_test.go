package brainfrick

import (
	"testing"
	"fmt"
)

func error_msg_generator[R any, D any](res R, desired_res D) string {
	return fmt.Sprintf("wanted: %v, got %v", desired_res, res);
}

func TestFromBrainBasic(t *testing.T) {
	res := FromBrain("+[----->+++<]>+.");
	if res != "h" {
		t.Fatal(error_msg_generator(res, "h"));
	}
}

func TestFromBrainIntermediate(t *testing.T) {
	res := FromBrain(">++++++++++[<+++++++++>-]<+++++++.><.>>");
	if res != "aa" {
		t.Fatal(error_msg_generator(res, "aa"));
	}
}

func TestFromBrainAdvanced(t *testing.T) {
	res := FromBrain(">++++++++++[<+++++++>-]<++.>>++++++++++[<++++++++++>-]<+.>>++++++++++[<++++++++++>-]<++++++++.><.>>>++++++++++[<+++++++++++>-]<+.>>++++++++++[<+++>-]<++.>>++++++++++[<++++++++>-]<+++++++.><<<.>>>>>++++++++++[<+++++++++++>-]<++++.><<<<<<<.>>>>>>>>>++++++++++[<++++++++++>-]<.>>++++++++++[<+++>-]<+++.>");
	if res != "Hello World!" {
		t.Fatal(error_msg_generator(res, "Hello World!"));
	}
}

func TestToBrainBasic(t *testing.T) {
	res := ToBrain("H");
	if res != ">++++++++++[<+++++++>-]<++.>" {
		t.Fatal(error_msg_generator(res, ">++++++++++[<+++++++>-]<++.>"))
	}
}

func TestToBrainIntermediate(t *testing.T) {
	res := ToBrain("aa");
	if res != ">++++++++++[<+++++++++>-]<+++++++.><.>>" {
		t.Fatal(error_msg_generator(res, ">++++++++++[<+++++++++>-]<+++++++.><.>>"))
	}
}

func TestToBrainAdvanced(t *testing.T) {
	res := ToBrain("Hello World!");
	if res != ">++++++++++[<+++++++>-]<++.>>++++++++++[<++++++++++>-]<+.>>++++++++++[<++++++++++>-]<++++++++.><.>>>++++++++++[<+++++++++++>-]<+.>>++++++++++[<+++>-]<++.>>++++++++++[<++++++++>-]<+++++++.><<<.>>>>>++++++++++[<+++++++++++>-]<++++.><<<<<<<.>>>>>>>>>++++++++++[<++++++++++>-]<.>>++++++++++[<+++>-]<+++.>" {
		t.Fatal(error_msg_generator(res, ">++++++++++[<+++++++>-]<++.>>++++++++++[<++++++++++>-]<+.>>++++++++++[<++++++++++>-]<++++++++.><.>>>++++++++++[<+++++++++++>-]<+.>>++++++++++[<+++>-]<++.>>++++++++++[<++++++++>-]<+++++++.><<<.>>>>>++++++++++[<+++++++++++>-]<++++.><<<<<<<.>>>>>>>>>++++++++++[<++++++++++>-]<.>>++++++++++[<+++>-]<+++.>"))
	}
}
