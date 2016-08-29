package bootloader

import (
	"os"
	"bufio"
	"fmt"
	"log"
)

func Run(fp string) {
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

func RunWithLog(fp string, l log.Logger) {
	f, e := os.Open(fp)
	if e != nil {
		l.Println("File at " + string(fp) + "cannot be found.")
		l.Println(e)
	}
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		li := sc.Text()
		l.Println(li)
	}
}