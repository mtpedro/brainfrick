package brainfrick

import (
	"testing"
	"fmt"
)

func error_msg_generator[R any, D any](res R, desired_res D) string {
	return fmt.Sprintf("wanted: %v, got %v", desired_res, res);
}

func TestToBrain(t *testing.T) {
	res := ToBrain("H");
	if res != ">++++++++++[<+++++++>-]++." {
		t.Fatal(error_msg_generator(res, ">++++++++++[<+++++++>-]++."))
	}
}


func TestFromBrain(t *testing.T) {
	res := FromBrain("+[----->+++<]>+.");
	if res != "h" {
		t.Fatal(error_msg_generator(res, "h"));
	}
}