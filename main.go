package main

import (
	"fmt"
	"os"
	"sort"
)

func main() {
	var str, op string
	var strings []string
	end := false

	for !end {
		fmt.Println("Introduce la cadena")
		fmt.Scan(&str)
		strings = append(strings, str)
		fmt.Print("Quiere introducir otra cadena? (s/n): ")
		fmt.Scan(&op)

		if op == "n" {
			end = true
		}
	}

	file, err := os.Create("asecendente.txt")
	if err != nil {
		println(err)
		return
	}

	sort.Strings(strings)
	for _, s := range strings {
		_, err := file.WriteString(s)
		if err != nil {
			println(err)
			return
		}
		_, err2 := file.WriteString("\n")
		if err2 != nil {
			println(err2)
			return
		}
	}

	file.Close()

	file2, err := os.Create("descendente.txt")
	if err != nil {
		println(err)
		return
	}

	sort.Slice(strings, func(i, j int) bool {
		return strings[i][0] > strings[j][0]
	})

	for _, s := range strings {
		_, err := file2.WriteString(s)
		if err != nil {
			println(err)
			return
		}
		_, err2 := file2.WriteString("\n")
		if err2 != nil {
			println(err2)
			return
		}
	}

	file2.Close()
}