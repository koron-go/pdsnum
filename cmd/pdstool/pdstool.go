package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"regexp"

	"github.com/koron-go/pdsnum"
)

var decodeMode bool

func decode(dst io.Writer, src io.Reader) error {
	// TODO:
	return nil
}

var rxNormalNumber = regexp.MustCompile(`(?:^|\s)([0-9]+)(?:\s|$)`)

func writeLine(w io.Writer, p []byte, lf bool) error {
	if _, err := w.Write(p); err != nil {
		return err
	}
	if lf {
		if _, err := w.Write([]byte("\n")); err != nil {
			return err
		}
	}
	return nil
}

func encode(dst io.Writer, src io.Reader) error {
	scanner := bufio.NewScanner(src)
	for scanner.Scan() {
		line := scanner.Bytes()
		matches := rxNormalNumber.FindAllSubmatchIndex(line, -1)
		if matches == nil {
			err := writeLine(dst, line, true)
			if err != nil {
				return err
			}
			continue
		}
		var curr = 0
		for _, m := range matches {
			start, end := m[2], m[3]
			if curr < start {
				err := writeLine(dst, line[curr:start], false)
				if err != nil {
					return err
				}
			}
			s, err := pdsnum.Encode(string(line[start:end]))
			if err != nil {
				return err
			}
			err = writeLine(dst, []byte(s), false)
			curr = end
		}
		err := writeLine(dst, line[curr:], true)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

func run(dst io.Writer, src io.Reader) error {
	if decodeMode {
		return decode(dst, src)
	}
	return encode(dst, src)
}

func main() {
	flag.BoolVar(&decodeMode, "decode", false, "decode PDS encoded strings")
	flag.Parse()
	err := run(os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}
