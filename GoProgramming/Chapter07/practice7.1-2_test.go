package Chapter01

import (
	"bufio"
	"fmt"
	"io"
	"strings"
	"testing"
)

func Test071(t *testing.T) {
	var c ByteCounter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)
}

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(strings.NewReader(string(p)))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		*c += 1
	}
	return len(p), nil
}

var position int64

type Counting struct {
	W io.Writer
}

func (c *Counting) Write(p []byte) (n int, err error) {
	n, err = c.W.Write(p)
	position += int64(n)
	return n, err
}

func CountingWriter(w io.Writer) (io.Writer, *int64) {
	return &Counting{w}, &position
}

