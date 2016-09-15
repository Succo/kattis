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
		linesS, _ := in.ReadString(' ')
		lines, _ := strconv.Atoi(linesS)
		bug, _ := in.ReadBytes('\n')
		for i := 0; i < lines; i++ {
			line, err := in.ReadBytes('\n')
			if err != nil {
				panic("Reading line failure")
			}
			if !bytes.Contains(line, bug) {
				out.Write(line)
				continue
			}
			subLines := bytes.Split(line, bug)
			for len(subLines) > 1 {
				var buffer bytes.Buffer
				for _, subLine := range subLines {
					buffer.Write(subLine)
				}
				line, err = buffer.ReadBytes('\n')
				if err != nil {
					panic("Reading line from buffer failure")
				}
				subLines = bytes.Split(line, bug)
			}
			out.Write(subLines[0])
		}

		out.Flush()
		if in.Buffered() == 0 {
			break
		}
	}
}
