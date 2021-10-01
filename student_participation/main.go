package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	s := flag.Int("number-students", 0, "")
	h := flag.Int("hands-raised", 0, "")
	x := flag.String("students-raised-hands", "", "coma separated")
	flag.Parse()
	ids := strings.Split(*x, ",")
	if len(ids) != *h {
		log.Fatal("number of hands-raised different from number of student ids that raised had")
	}

	m := map[int]int{}
	for _, v := range ids {
		tmp, _ := strconv.Atoi(v)
		m[tmp] += 1
	}
	for i := 0; i < *s; i++ {
		times, _ := m[i+1]
		fmt.Printf("Student %d raised their hand %d times.\n", i+1, times)
	}
}
