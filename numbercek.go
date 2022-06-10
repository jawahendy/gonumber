package numbercek

import "fmt"

func Cekganjilgenapv1(d int) {

	if d%2 == 0 {
		fmt.Println("genap")
	} else {
		fmt.Println("ganjil")
	}
}

func Cekganjilgenapv2(numbs ...int) {
    for _, number := range numbs {
        if number % 2 == 0 {
            fmt.Print("genap ")
        } else {
            fmt.Print("ganjil ")
        }
    }
}