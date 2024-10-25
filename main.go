package main

import (
	"fmt"
	"github.com/ifuryst/go-benchstat-demo/pkg/util"
)

func main() {
	mockBiz()
}

func mockBiz() {
	ss := []string{"a", "b", "c"}
	s1 := "a"
	isInclude := util.Includes(ss, s1)
	fmt.Printf("Is %s in %v? %v\n", s1, ss, isInclude)
	s2 := "d"
	isInclude = util.Includes(ss, s2)
	fmt.Printf("Is %s in %v? %v\n", s2, ss, isInclude)
}
