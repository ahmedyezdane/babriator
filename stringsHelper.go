package main

import (
	"strconv"
	"strings"
)

func JoinBytes(bs []byte, separator string) string {
	
	var sb strings.Builder

	for i := 0; i < len(bs); i++ {

		sb.WriteString(strconv.Itoa(int(bs[i])))

		if i < (len(bs)-1) {
			sb.WriteString("-")
		}
	}

	return sb.String()
}