package error

import (
	"errors"
	"fmt"
	"testing"
)

var LessThanTwoError = errors.New("n should bi not less than 2")
var LargerThenHundredError = errors.New("n should bi not larger than 100")

func GetFibonacci(n int) ([]int, error) {
	if n < 2 {
		return nil, LessThanTwoError
	}
	if n > 100 {
		return nil, LargerThenHundredError
	}
	fibList := []int{1, 1}

	for i := 2; i < n; i++ {
		fibList = append(fibList, fibList[i-2]+fibList[i-1])
	}
	return fibList, nil
}

func TestGetFibnoacci(t *testing.T) {
	t.Log(GetFibonacci(5))
	if v, err := GetFibonacci(-10); err != nil {
		if err == LessThanTwoError {
			fmt.Println("it is less")
		}
		t.Error(err)
	} else {
		t.Log(v)
	}
}
