package main

import (
	"fmt"
	"github.com/jlevitt/katas/numfilter"
)

func main() {
	filter := numfilter.ShortenNumber([]string{"", "k", "m"}, 1000)
	fmt.Printf("%v\n", filter("314"))
	fmt.Printf("%v\n", filter("234324"))
	fmt.Printf("%v\n", filter("98234324"))
	fmt.Printf("%v\n", filter("98234324023"))

	filter = numfilter.ShortenNumber([]string{"B", "KB", "MB", "GB"}, 1024)
	fmt.Printf("%v\n", filter("pippi"))
	fmt.Printf("%v\n", filter("32"))
	fmt.Printf("%v\n", filter("2100"))
	fmt.Printf("%v\n", filter("3145728"))
	fmt.Printf("%v\n", filter("pippi"))
	fmt.Printf("%v\n", filter("13922402"))
}
