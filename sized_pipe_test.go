package sized_pipe

import (
	"fmt"
	"testing"
)

func TestSizedPipe(t *testing.T) {
	p := NewSizedPipe(3)
	fmt.Println(p.Elements())
	p.Push(1)
	fmt.Println(p.Elements(), p.Arr)
	p.Push(2)
	fmt.Println(p.Elements(), p.Arr)
	p.Push(3)
	fmt.Println(p.Elements(), p.Arr)
	p.Push(4)
	fmt.Println(p.Elements(), p.Arr)
	p.Push(5)
	fmt.Println(p.Elements(), p.Arr)
	p.Push(6)
	fmt.Println(p.Elements(), p.Arr)
	p.Push(7)
	fmt.Println(p.Elements(), p.Arr)
}

func TestSizedPipe2(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	p := NewSizedPipe(0)
	fmt.Println(p.Elements())
	p.Push(1)
	fmt.Println(p.Elements(), p.Arr)
}
