package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"strings"
)

func main() {
	key := "pqrstuv"
	fmt.Printf("part 1: %d\n", getFirstCoinIndex(key, 5))
	fmt.Printf("part 2: %d\n", getFirstCoinIndex(key, 6))
}

func getFirstCoinIndex(key string, prefixLen int) int {
	prefix := strings.Repeat("0", prefixLen)
	i := 1
	for {
		data := fmt.Appendf(nil, "%s%d", key, i)
		bytes := md5.Sum(data)
		str := hex.EncodeToString(bytes[:])
		if strings.HasPrefix(str, prefix) {
			return i
		}
		i += 1
	}
}
