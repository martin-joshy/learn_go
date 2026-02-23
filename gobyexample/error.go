package main

import (
	"errors"
	"fmt"
)

func handleError(v int) (int, error) {
	if v == 30 {
		return -1, errors.New("I hate 30")
	} else {
		return v + 3, nil
	}
}

var (
	errPower    = errors.New("Can't boil the water")
	errOutOfTea = errors.New("We are out of tea")
)

func makeTea(t int) error {
	if t == 1 {
		fmt.Println("We made the Tea ..!!")
	} else if t == 4 {
		return fmt.Errorf("Please buy tea bags %w", errOutOfTea)
	} else {
		return errPower
	}
	return nil
}

func main() {
	res := make([]int, 3)
	fmt.Println(res)
	// for _, i := range []int{30, 20} {
	// 	if res, err := handleError(i); err != nil {
	// 		fmt.Println("I dont like this number", i, err)
	// 	} else {
	// 		fmt.Println("I love this number", i, "Here is the result: ", res)
	// 	}
	// }
	//
	// for _, i := range []int{1, 4, 9} {
	// 	if err := makeTea(i); err != nil {
	// 		if errors.Is(err, errOutOfTea) {
	// 			fmt.Println("Okay I am going to buy tea as ran out")
	// 		} else if errors.Is(err, errPower) {
	// 			fmt.Println("Okay I am going to pay the gas bill")
	// 		} else {
	// 			fmt.Printf("Unknow error %s/n", err)
	// 		}
	// 		fmt.Println("The tea is ready")
	// 	}
	// }
}
