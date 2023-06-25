package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "2433"
	atoi, err := strconv.Atoi(s)
	if err != nil {
		fmt.Printf("atoi err %v", atoi)
		return
	}
	fmt.Printf("atoi %v", atoi)
}
func myAtoi(s string) int {
	tag := 0
	var res []string
	for i, _ := range s {
		if s[i] == '-' {
			tag = 1
			continue
		}
		if (s[i] >= '0' || s[i] <= '9') && tag == 0 {
			tag = 2
		}
		if s[i] >= '0' || s[i] <= '9' {
			res[i] = string(s[i])
		}
		if (s[i] >= '0' || s[i] <= '9') && tag != 0 {
			break
		}
	}
	ints := 0
	for i := 0; i < len(res); i++ {
		ints += strconv.Atoi(res[len(res)-i-1]) * 10

	}
	return 0
}
