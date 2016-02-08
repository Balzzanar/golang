package main

import (
	"fmt"
	//"io/ioutil"
	"strings"
	"strconv"
	"sort"
	)

func main(){
	fmt.Println("----- Day 2 -----")
	fmt.Println(calc_pkg("2x3x4"))

}

/**
 * L x W x H
 */
func calc_pkg(dim string) int64{
	var total int64 = 0
	dims := strings.Split(dim, "x")
	lenght,_ := int(strconv.ParseInt(dims[0], 10, 32)
	width,_ := strconv.ParseInt(dims[1], 10, 32)
	height,_ := strconv.ParseInt(dims[2], 10, 32)

	var list []int64
	append(list, lenght * width)
	append(list, width * height)
	append(list, height * lenght)
	sort.Ints(list)

	fmt.Println(list)
//	total += 2 * lw
//	total += 2 * wh
//	total += 2 * hl
	return total
}