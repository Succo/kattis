package main

import (
	"bufio"
	"bytes"
	"os"
	"strconv"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		linesDescr, _ := in.ReadString(' ')
		lines, _ := strconv.Atoi(linesDescr[:len(linesDescr)-1])
		bug, _ := in.ReadBytes('\n')
		bug = bug[:len(bug)-1]
		var old bytes.Buffer
		for i := 0; i < lines; i++ {
			line, err := in.ReadBytes('\n')
			if err != nil {
				panic("Reading line failure")
			}
			old.Write(line)
		}

		// old and new []byte
		o := old.Bytes()
		n := make([]byte, len(o))
		for {
			bugs := 0
			// index where we are writing
			w := 0
			// index where we are starting to read
			r := 0
			for {
				i := bytes.Index(o[r:], bug)
				if i == -1 {
					break
				}
				bugs++
				w += copy(n[w:], o[r:r+i])
				r += i + len(bug)
			}
			copy(n[w:], o[r:len(o)])
			o = n
			if bugs == 0 {
				break
			}
			n = make([]byte, len(o)-bugs)
		}
		out.Write(o)
		out.Flush()
		if in.Buffered() == 0 {
			break
		}
	}
}
