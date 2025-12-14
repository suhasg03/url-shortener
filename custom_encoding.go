package main

// import (
// 	"fmt"
// 	"strconv"
// )

// func encodeXOR(n int64, key byte) string {
// 	bytes := make([]byte, 8)
// 	for i := 0; i < 8; i++ {
// 		bytes[i] = byte(n>>(i*8)) ^ key
// 	}
// 	return fmt.Sprintf("%08x", bytes) // Hex string
// }

// func decodeXOR(s string, key byte) (int64, error) {
// 	bytes := make([]byte, 8)
// 	for i := 0; i < 8; i++ {
// 		b, _ := strconv.ParseUint(s[i*2:i*2+2], 16, 8)
// 		bytes[i] = byte(b) ^ key
// 	}
// 	var result int64
// 	for i := 0; i < 8; i++ {
// 		result |= int64(bytes[i]) << (i * 8)
// 	}
// 	return result, nil
// }
