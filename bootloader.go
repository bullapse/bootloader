package bootloader

import (
	"os"
	"bufio"
	"fmt"
)

func run(fp string) {
	f, e := os.Open(fp)
	if e != nil {
		fmt.Println("File at " + string(fp) + "cannot be found.")
		fmt.Println(e)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		l := sc.Text()
		fmt.Println(l)
	}
}
