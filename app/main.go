package main

import (
	"fmt"
	"github.com/jaco60/zipcodes"
	"sort"
	"strings"
)

func sorted(tab []string) []string {
	res := make([]string, len(tab))
	copy(res, tab)
	sort.Strings(res)
	return res
}

// Use case...
func main() {

	var zip, city string

	for {
		fmt.Print("Entrez un code postal (xxx pour terminer) : ")
		fmt.Scan(&zip)
		cities, ok := zipcodes.MapZipToCities[zip]
		if ok {
			fmt.Println(strings.Join(sorted(cities), ", "))
		} else if zip[0] == 'x' {
			break
		}
	}
	for {
		fmt.Print("Entrez un nom de commune (xxx pour terminer) : ")
		fmt.Scan(&city)
		zips, ok := zipcodes.MapCityToZips[city]
		if ok {
			fmt.Println(strings.Join(sorted(zips), ", "))
		} else if city[0] == 'x' {
			break
		}
	}

}
