package main

import (
	"fmt"
	"github.com/jaco60/zipcodes"
	sort2 "sort"
	"strings"
)

func sort(tab []string) []string {
	res := make([]string, len(tab))
	copy(res, tab)
	sort2.Strings(res)
	return res
}

// Use case...
func main() {
	cities := zipcodes.MapZipToCities
	fmt.Println(strings.Join(sort(cities["95000"]), ", "))
	fmt.Println(strings.Join(sort(cities["95800"]), ", "))
	fmt.Println(strings.Join(sort(cities["31870"]), ", "))
	fmt.Println(strings.Join(sort(cities["72380"]), ", "))

	zips := zipcodes.MapCityToZips
	fmt.Println(strings.Join(sort(zips["PARIS"]), ", "))
	fmt.Println(strings.Join(sort(zips["LAGARDELLE SUR LEZE"]), ", "))
	fmt.Println(strings.Join(sort(zips["ANTOIGNE"]), ", "))

}
