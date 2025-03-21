// Code generated by mktypedtests.go. DO NOT EDIT.
//
//go:generate go run mktypedtests.go
package broad_test

import (
	"math/rand/v2"
	"testing"
)

func TestByte0(t *testing.T) {
	data := make([][0]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte1(t *testing.T) {
	data := make([][1]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte2(t *testing.T) {
	data := make([][2]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte3(t *testing.T) {
	data := make([][3]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte4(t *testing.T) {
	data := make([][4]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte5(t *testing.T) {
	data := make([][5]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte6(t *testing.T) {
	data := make([][6]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte7(t *testing.T) {
	data := make([][7]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte8(t *testing.T) {
	data := make([][8]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte9(t *testing.T) {
	data := make([][9]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte10(t *testing.T) {
	data := make([][10]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte11(t *testing.T) {
	data := make([][11]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte12(t *testing.T) {
	data := make([][12]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte13(t *testing.T) {
	data := make([][13]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte14(t *testing.T) {
	data := make([][14]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte15(t *testing.T) {
	data := make([][15]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte16(t *testing.T) {
	data := make([][16]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte17(t *testing.T) {
	data := make([][17]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte18(t *testing.T) {
	data := make([][18]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte19(t *testing.T) {
	data := make([][19]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte20(t *testing.T) {
	data := make([][20]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte21(t *testing.T) {
	data := make([][21]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte22(t *testing.T) {
	data := make([][22]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte23(t *testing.T) {
	data := make([][23]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte24(t *testing.T) {
	data := make([][24]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte25(t *testing.T) {
	data := make([][25]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte26(t *testing.T) {
	data := make([][26]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte27(t *testing.T) {
	data := make([][27]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte28(t *testing.T) {
	data := make([][28]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte29(t *testing.T) {
	data := make([][29]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte30(t *testing.T) {
	data := make([][30]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte31(t *testing.T) {
	data := make([][31]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte32(t *testing.T) {
	data := make([][32]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte33(t *testing.T) {
	data := make([][33]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte34(t *testing.T) {
	data := make([][34]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte35(t *testing.T) {
	data := make([][35]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte36(t *testing.T) {
	data := make([][36]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte37(t *testing.T) {
	data := make([][37]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte38(t *testing.T) {
	data := make([][38]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte39(t *testing.T) {
	data := make([][39]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte40(t *testing.T) {
	data := make([][40]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte41(t *testing.T) {
	data := make([][41]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte42(t *testing.T) {
	data := make([][42]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte43(t *testing.T) {
	data := make([][43]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte44(t *testing.T) {
	data := make([][44]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte45(t *testing.T) {
	data := make([][45]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte46(t *testing.T) {
	data := make([][46]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte47(t *testing.T) {
	data := make([][47]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte48(t *testing.T) {
	data := make([][48]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte49(t *testing.T) {
	data := make([][49]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte50(t *testing.T) {
	data := make([][50]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte51(t *testing.T) {
	data := make([][51]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte52(t *testing.T) {
	data := make([][52]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte53(t *testing.T) {
	data := make([][53]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte54(t *testing.T) {
	data := make([][54]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte55(t *testing.T) {
	data := make([][55]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte56(t *testing.T) {
	data := make([][56]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte57(t *testing.T) {
	data := make([][57]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte58(t *testing.T) {
	data := make([][58]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte59(t *testing.T) {
	data := make([][59]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte60(t *testing.T) {
	data := make([][60]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte61(t *testing.T) {
	data := make([][61]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte62(t *testing.T) {
	data := make([][62]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte63(t *testing.T) {
	data := make([][63]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte64(t *testing.T) {
	data := make([][64]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte65(t *testing.T) {
	data := make([][65]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte66(t *testing.T) {
	data := make([][66]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte67(t *testing.T) {
	data := make([][67]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte68(t *testing.T) {
	data := make([][68]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte69(t *testing.T) {
	data := make([][69]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte70(t *testing.T) {
	data := make([][70]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte71(t *testing.T) {
	data := make([][71]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte72(t *testing.T) {
	data := make([][72]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte73(t *testing.T) {
	data := make([][73]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte74(t *testing.T) {
	data := make([][74]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte75(t *testing.T) {
	data := make([][75]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte76(t *testing.T) {
	data := make([][76]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte77(t *testing.T) {
	data := make([][77]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte78(t *testing.T) {
	data := make([][78]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte79(t *testing.T) {
	data := make([][79]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte80(t *testing.T) {
	data := make([][80]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte81(t *testing.T) {
	data := make([][81]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte82(t *testing.T) {
	data := make([][82]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte83(t *testing.T) {
	data := make([][83]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte84(t *testing.T) {
	data := make([][84]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte85(t *testing.T) {
	data := make([][85]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte86(t *testing.T) {
	data := make([][86]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte87(t *testing.T) {
	data := make([][87]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte88(t *testing.T) {
	data := make([][88]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte89(t *testing.T) {
	data := make([][89]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte90(t *testing.T) {
	data := make([][90]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte91(t *testing.T) {
	data := make([][91]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte92(t *testing.T) {
	data := make([][92]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte93(t *testing.T) {
	data := make([][93]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte94(t *testing.T) {
	data := make([][94]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte95(t *testing.T) {
	data := make([][95]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte96(t *testing.T) {
	data := make([][96]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte97(t *testing.T) {
	data := make([][97]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte98(t *testing.T) {
	data := make([][98]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte99(t *testing.T) {
	data := make([][99]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte100(t *testing.T) {
	data := make([][100]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte101(t *testing.T) {
	data := make([][101]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte102(t *testing.T) {
	data := make([][102]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte103(t *testing.T) {
	data := make([][103]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte104(t *testing.T) {
	data := make([][104]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte105(t *testing.T) {
	data := make([][105]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte106(t *testing.T) {
	data := make([][106]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte107(t *testing.T) {
	data := make([][107]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte108(t *testing.T) {
	data := make([][108]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte109(t *testing.T) {
	data := make([][109]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte110(t *testing.T) {
	data := make([][110]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte111(t *testing.T) {
	data := make([][111]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte112(t *testing.T) {
	data := make([][112]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte113(t *testing.T) {
	data := make([][113]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte114(t *testing.T) {
	data := make([][114]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte115(t *testing.T) {
	data := make([][115]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte116(t *testing.T) {
	data := make([][116]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte117(t *testing.T) {
	data := make([][117]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte118(t *testing.T) {
	data := make([][118]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte119(t *testing.T) {
	data := make([][119]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte120(t *testing.T) {
	data := make([][120]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte121(t *testing.T) {
	data := make([][121]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte122(t *testing.T) {
	data := make([][122]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte123(t *testing.T) {
	data := make([][123]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte124(t *testing.T) {
	data := make([][124]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte125(t *testing.T) {
	data := make([][125]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte126(t *testing.T) {
	data := make([][126]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte127(t *testing.T) {
	data := make([][127]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
func TestByte128(t *testing.T) {
	data := make([][128]byte, 100)
	rng := rand.NewChaCha8([32]byte{})
	for j := 0; j < len(data); j++ {
		rng.Read((&data[j])[:])
	}
	testTyped(t, data)
}
