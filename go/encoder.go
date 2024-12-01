package main

import (
	"strconv"
	"strings"
)

type Codec struct {
}

const delim = "#"

func (codec *Codec) Encode(strs []string) string {
	var builder strings.Builder

	for _, s := range strs {
		builder.WriteString(strconv.Itoa(len(s)) + delim)
		builder.WriteString(s)
	}

	return builder.String()
}

func (codec *Codec) Decode(strs string) []string {
	var result []string

	i := 0
	for i < len(strs) {
		delimPos := strings.Index(strs[i:], delim)
		if delimPos == -1 {
			break
		}

		lenStr := strs[i : i+delimPos]
		length, _ := strconv.Atoi(lenStr)

		i += delimPos + 1

		result = append(result, strs[i:i+length])
		i += length
	}

	return result
}
