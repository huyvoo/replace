package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/hvyavuno/replace/rp"
)

var fname *string
var target *string
var d *string

func main() {
	fname = flag.String("f", "", "file name to be replaced")
	d = flag.String("d", "", "the dir to search for")
	target = flag.String("t", "", "file name that will overwrite")

	flag.Parse()

	
	ct := true 

	if *fname == "" {
		ct = false 

		fmt.Println("fname is empty")
	}

	if *d == "" {
		ct = false
		fmt.Println("d is empty")

	}

	if *target == "" {
		ct = false 

		fmt.Println("target is empty")

	}

	if !ct {
		os.Exit(1)
	}

	s, _ := rp.Exists(*d)

	if !s {
		fmt.Printf("%s does not exists\n", *d)
		os.Exit(1)
	}

	rp.FindAndRename2(*d, *fname, *target)
}
