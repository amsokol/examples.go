package main

import (
	"errors"
	"fmt"
	"log"
)

type Data struct {
	value string
	flag  bool
}

//go:noinline
func paramsByVal(dv Data) {
	vv := len(dv.value)
	if dv.flag {
		vv += 15 - len(dv.value)
	}

	log.Printf("value: %d", vv)
}

//go:noinline
func paramsByRef(dr *Data) {
	vr := len(dr.value)
	if dr.flag {
		vr += 15 - len(dr.value)
		dr.value = "dddd"
	}

	log.Printf("value: %d", vr)
}

//go:noinline
func returnErr(drr *Data) error {
	err := errors.New("some error")
	if err != nil {
		return fmt.Errorf("another one error %s: %w", drr.value, err)
	}

	return nil
}

func main() {
	d1 := Data{"", true}
	d2 := Data{"", false}

	if d1.flag {
		d1.value = "date1"
	}

	if !d1.flag {
		d1.value = "date2"
	}

	paramsByVal(d1)
	paramsByRef(&d2)

	_ = returnErr(&d2)
}

// go build -gcflags "-m -m"
