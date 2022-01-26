package hex

import (
	"fmt"
	"math"
	"math/big"
	"strings"
	"testing"
)

var test RGBaColor = func() (red, green, blue byte, alpha float64) {
	return 67, 255, 100, 0.85
}

var testBad RGBaColor = func() (red, green, blue byte, alpha float64) {
	return 67, 255, 100, 0.86
}

var test2 RGBaColor = func() (red, green, blue byte, alpha float64) {
	return 67, 255, 100, AlphaFromPercent(85)
}

func similarWithTolerance(a, b, tolerance float64) bool {
	result := big.NewFloat(a).Cmp(big.NewFloat(b))
	if result == -1 {
		fmt.Printf("big.NewFloat().Cmp(): %f < %f", a, b)
	} else if result == 1 {
		fmt.Printf("big.NewFloat().Cmp(): %f > %f", a, b)
	} else if result == 0 {
		fmt.Printf("big.NewFloat().Cmp(): %f > %f", a, b)
	}
	fmt.Println()
	if diff := math.Abs(a - b); diff < tolerance {
		fmt.Printf("When a=%f and b =%f => Nearly same by tolerance\n", a, b)
		return true
	} else {
		fmt.Printf("When a=%f and b=%f => Not same Even by Tolerance\n", a, b)
		return false
	}
}

func TestFromRGBa(t *testing.T) {
	color := FromRGBa(67, 255, 100, 0.85)
	t.Log(color)
	t.Logf("int: %d", color)
	if color.String() != "43ff64d9" {
		t.Errorf("hex: %s", color)
		t.Fatalf("int: %d", color)
	}
}

func TestToRGBa(t *testing.T) {
	color := HexadecimalColor(0x43ff64d9)
	rgba := color.ToRGBa()
	r, g, b, a := rgba()
	matchR, matchG, matchB, matchA := test()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if !similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}
	matchR, matchG, matchB, matchA = test2()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if !similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}
	matchR, matchG, matchB, matchA = testBad()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}
}
func TestAsRGBa(t *testing.T) {
	color := HexadecimalColor(0x43ff64d9)
	r, g, b, a := AsRGBa(color)
	matchR, matchG, matchB, matchA := test()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if !similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}
	matchR, matchG, matchB, matchA = test2()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if !similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}
	matchR, matchG, matchB, matchA = testBad()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}

}

func TestFromString(t *testing.T) {
	colorString := "43ff64d9"
	colorString2 := "#43ff64d9"
	if hex := FromString(colorString); hex.String() != colorString {
		t.Fatalf("error: hex: %s != sample: %s", hex, colorString)
	} else if hex := FromString(colorString2); hex.String() != strings.TrimPrefix(colorString2, "#") {
		t.Fatalf("error: hex: %s != sample: %s", hex, colorString2)
	}
	hex := FromString(colorString)
	r, g, b, a := AsRGBa(hex)
	matchR, matchG, matchB, matchA := test()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if !similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}
	hex = FromString(colorString2)
	r, g, b, a = AsRGBa(hex)
	matchR, matchG, matchB, matchA = test()
	if r != matchR {
		t.Fatalf("r %d no match %d", r, matchR)
	} else if g != matchG {
		t.Fatalf("g %d no match %d", g, matchG)
	} else if b != matchB {
		t.Fatalf("b %d no match %d", b, matchB)
	} else if !similarWithTolerance(matchA, a, 0.0001) {
		t.Fatalf("a %f no match %f", a, matchA)
	}
}
