package test

import (
	"github.com/Serein-sz/func-util/util"
	"reflect"
	"testing"
)

func TestBaseMap(t *testing.T) {
	pendingTestSlice := make([]int, 0)
	pendingTestSlice = append(pendingTestSlice, 1)
	pendingTestSlice = append(pendingTestSlice, 2)
	pendingTestSlice = append(pendingTestSlice, 3)
	pendingTestSlice = append(pendingTestSlice, 4)
	pendingTestSlice = append(pendingTestSlice, 5)
	pendingTestSlice = append(pendingTestSlice, 6)

	expectResult := make([]int, 0)
	for _, v := range pendingTestSlice {
		expectResult = append(expectResult, v*2)
	}

	realResult := util.Map(pendingTestSlice, func(i int) int {
		return i * 2
	})

	if !reflect.DeepEqual(realResult, expectResult) {
		t.Errorf("expect %v got %v", expectResult, realResult)
	}
}

func TestStructMap(t *testing.T) {
	type Position struct {
		x int
		y int
	}
	pendingTestSlice := make([]Position, 0)
	pendingTestSlice = append(pendingTestSlice, Position{0, 8})
	pendingTestSlice = append(pendingTestSlice, Position{1, 1})
	pendingTestSlice = append(pendingTestSlice, Position{2, 1})
	pendingTestSlice = append(pendingTestSlice, Position{3, 3})
	pendingTestSlice = append(pendingTestSlice, Position{4, 1})

	expectResult := make([]Position, 0)
	for _, p := range pendingTestSlice {
		p.x++
		expectResult = append(expectResult, p)
	}

	realResult := util.Map(pendingTestSlice, func(p Position) Position {
		p.x++
		return p
	})

	if !reflect.DeepEqual(realResult, expectResult) {
		t.Errorf("expect %v got %v", expectResult, realResult)
	}
}
