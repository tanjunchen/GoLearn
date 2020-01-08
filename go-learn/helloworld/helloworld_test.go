package helloworld

import "testing"

func TestHelloWorld(t *testing.T) {
	expected := 2
	actual := HelloWorld()

	if actual != expected {
		t.Errorf("Expect 1, but got %d!", actual)
	}
}

func TestSum(t *testing.T) {
	set := []int{17, 23, 100, 76, 55}
	expected := 271
	actual := Sum(set)

	if actual != expected {
		t.Errorf("Expect %d, but got %d!", expected, actual)
	}
}
