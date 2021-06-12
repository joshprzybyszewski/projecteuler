package easy

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/joshprzybyszewski/projecteuler/mathUtils"
)

const (
	problem67TriangleFileName = `easy/p067_triangle.txt`
)

func SolveProblem67() string {
	/*
		Find the maximum total from top to bottom in
		triangle.txt (right click and 'Save Link/Target As...'),
		a 15K text file containing a triangle with one-hundred rows.
	*/

	dat, err := ioutil.ReadFile(problem67TriangleFileName)
	if err != nil {
		log.Fatalf("problem 67 cannot read file: %v", err)
	}
	tri := convertTriangleToSlices(string(dat))
	triStr := []string{}
	for i, row := range tri {
		triStr = append(triStr, fmt.Sprintf("%d: %v", i, row))
	}
	res := collapseMaxesDownTriangle(tri)
	ans := mathUtils.MaxInSlice(res)
	return fmt.Sprintf("%v", ans)
}
