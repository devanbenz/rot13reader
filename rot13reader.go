package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (reader *rot13Reader) Read(b []byte) (int, error) {
	r, err := reader.r.Read(b)
	for x := range b[:r] {
		fmt.Println(b[x])
		if b[x] != 32 {
			switch {
			case b[x] > 109:
				a := 123 - b[x]
				b[x] = (13 - a) + 65
			default:
				b[x] = b[x] + 13
			}
		}
	}
	return r, err
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
