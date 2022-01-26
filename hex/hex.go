package hex

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/* HexadecimalColor shows how RGBa should be displayed as hexidecimal
   0x  ff ff ff ff    -> to RGBa
       R  G  B  a

Where 'a' is contains bounds of 0.0 to 1.0
*/
type HexadecimalColor uint32

type RGBaColor func() (red, green, blue byte, alpha float64)

const pureWhite HexadecimalColor = 0xffffffff
const pureBlack HexadecimalColor = 0x00000000

func AlphaFromPercent(percent float64) float64 {
	return percent / 100
}

func FromRGBa(r, g, b byte, a float64) HexadecimalColor {
	if a < 0 || a > 1 {
		panic(fmt.Errorf("RGB's alpha must be between a < 0 || a > 1, a=%f", a))
	}
	return HexadecimalColor(r)<<24 + HexadecimalColor(g)<<16 + HexadecimalColor(b)<<8 + HexadecimalColor(math.Ceil(a*255))
}

func (c HexadecimalColor) String() string {
	return fmt.Sprintf("%08x", uint32(c))
}

func FromRGBaColor(color RGBaColor) HexadecimalColor {
	r, g, b, a := color()
	return FromRGBa(r, g, b, a)
}

//func (c RGBaColor) String() string {
//	r, g, b, a := c()
//	return fmt.Sprintf("red: %d, green: %d, blue: %d, alpha: %f", r, g, b, a)
//}

func AsRGBa(hexCode HexadecimalColor) (red, green, blue byte, alpha float64) {
	// aa := float64(hexCode)/255/10
	//fmt.Printf("hex: %d = red: %d, green: %d, blue: %d, alpha: %f = %f -> %f", hexCode, r, g, b, aa, a, aaa)
	return byte(hexCode >> 24), byte(hexCode >> 16), byte(hexCode >> 8), (math.Ceil(float64(hexCode>>1)+.5) / 100) - math.Floor(math.Ceil(float64(hexCode>>1)+.5)/100)
}

func (c HexadecimalColor) ToRGBa() RGBaColor {
	return func() (red, green, blue byte, alpha float64) {
		return AsRGBa(c)
	}
}

func FromString(code string) HexadecimalColor {
	code = strings.TrimPrefix(code, "#")
	code = strings.TrimPrefix(code, "0x")
	if len(code) > 8 {
		panic("must be 8 figures")
	}
	if hex, err := strconv.ParseUint(code, 16, 32); err != nil {
		panic(err)
	} else {
		return HexadecimalColor(hex)
	}
	panic("error converting")

}
