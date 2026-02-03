// Package pdsnum provides Positional Description Scheme (PDS) tailored for
// digit sequences.
//
// See https://arxiv.org/abs/2408.12430v1 for details.
package pdsnum

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

// Encode applies PDS conversion to a digit sequence.
func Encode(s string) (string, error) {
	// FIXME: consider decimal points
	n := len(s)
	bb := &bytes.Buffer{}
	bb.WriteRune('_')
	for i, d := range s {
		if d == '0' {
			continue
		}
		if d < '1' || d > '9' {
			return "", fmt.Errorf("not a digit: found '%c' at %d", d, i)
		}
		fmt.Fprintf(bb, " %c 0%d", d, n-i)
	}
	bb.WriteString(" _")
	return bb.String(), nil
}

func Decode(s string) (string, error) {
	elements := strings.Split(strings.TrimSpace(s), " ")
	// Number of elements should be even.
	if len(elements)%2 != 0 {
		return "", errors.New("odd number of elements")
	}
	// First and last should be "_"
	if elements[0] != "_" {
		return "", errors.New("first is not underbar")
	}
	if elements[len(elements)-1] != "_" {
		return "", errors.New("last is not underbar")
	}
	// Trim leading and trailing "_"
	elements = elements[1 : len(elements)-1]
	num := 0
	for i := 0; i < len(elements); i += 2 {
		// Decode a digit
		dig, pos := elements[i], elements[i+1]
		if len(dig) != 1 || dig[0] < '0' || dig[0] > '9' {
			return "", fmt.Errorf("not a digit: %s at %s", dig, pos)
		}
		dn := int(dig[0] - '0')
		// Decode a position
		if len(pos) < 2 || pos[0] != '0' {
			return "", fmt.Errorf("not a pos: %s", pos)
		}
		pn, err := strconv.Atoi(pos[1:])
		if err != nil {
			return "", fmt.Errorf("not a pos: %w", err)
		}
		num += dn * int(math.Pow10(pn-1))
	}
	return strconv.Itoa(num), nil
}
