package main

import (
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"reflect"
)

func HexToLittleEndian(s string) *big.Int {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	data = reverseSliceOfByte(data)
	return new(big.Int).SetBytes(data)
}

func HexToBigEndian(s string) *big.Int {
	data, err := hex.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return new(big.Int).SetBytes(data)
}

func BigEndianToHex(i *big.Int) string {
	d := i.Bytes()
	return hex.EncodeToString(d)
}

func LittleEndianToHex(i *big.Int) string {
	d := i.Bytes()
	d = reverseSliceOfByte(d)
	return hex.EncodeToString(d)
}

func reverseSliceOfByte(s []byte) []byte {
	n := reflect.ValueOf(s).Len()
	swap := reflect.Swapper(s)
	for i, j := 0, n-1; i < j; i, j = i+1, j-1 {
		swap(i, j)
	}
	return s
}

func main() {
	s := "ff0000000000"
	fmt.Println(HexToBigEndian(s))
	fmt.Println(HexToLittleEndian(s))
	a, ok := new(big.Int).SetString("280375465082880", 10)
	if !ok {
		panic(errors.New("failed to set string"))
	}
	fmt.Println(BigEndianToHex(a))
	fmt.Println(LittleEndianToHex(a))
}
