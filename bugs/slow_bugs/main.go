package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	for {
		var lines int
		var bug []byte
		fmt.Fscanln(in, &lines, &bug)
		for i := 0; i < lines; i++ {
			line, err := in.ReadBytes('\n')
			if err != nil {
				panic("Reading line failure")
			}
			if !bytes.Contains(line, bug) {
				out.Write(line)
				continue
			}

			n := make([]byte, len(line))
			for {
				i := bytes.Index(line, bug)
				if i == -1 {
					break
				}
				copy(n[0:], line[:i])
				copy(n[i:], line[i+len(bug):len(line)])
				line = n
				n = make([]byte, len(line))
			}
			out.Write(line)
			out.Flush()
		}

		if in.Buffered() == 0 {
			break
		}
	}
}
