package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	// these are the inputs in the form of "flags" which are the command line arguments that begin with --
	s := flag.Int("number-students", 0, "")
	h := flag.Int("hands-raised", 0, "")
	x := flag.String("students-raised-hands", "", "coma separated")
	flag.Parse()

	// split students-raised-hands (x) by "," turning into an array(or slice, how's known in Go)
	ids := strings.Split(*x, ",")

	// validate the number of hands-raised and students-raised-hands that must be equal
	if len(ids) != *h {
		log.Fatal("number of hands-raised different from number of student ids that raised had")
	}

	// create a map that will store the student id in the map key and the number of times he raised the hands in the map value.
	// Visualization:
	// [2] => 1
	// [3] => 5
	// in this case student 2 raised the hand 1 times and student 3 raised 5 times
	m := map[int]int{}
	for _, v := range ids {
		tmp, _ := strconv.Atoi(v)
		m[tmp] += 1
	}

	// loop from 0 to number of students
	for i := 0; i < *s; i++ {
		// get the number of times studend raised hand from the map using the studend id in the key
		times, _ := m[i+1]
		fmt.Printf("Student %d raised their hand %d times.\n", i+1, times)
	}
}
