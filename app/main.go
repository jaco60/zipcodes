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

	for {
		zip := ""
		fmt.Print("Entrez un code postal (Retour pour terminer) : ")
		fmt.Scanln(&zip)
		cities, ok := zipcodes.MapZipToCities[zip]
		if ok {
			fmt.Println(strings.Join(sorted(cities), ", "))
		} else if zip == "" {
			break
		}
	}
	for {
		city := ""
		fmt.Print("Entrez un nom de commune (Retour pour terminer) : ")
		fmt.Scanln(&city)
		zips, ok := zipcodes.MapCityToZips[city]
		if ok {
			fmt.Println(strings.Join(sorted(zips), ", "))
		} else if city == "" {
			break
		}
	}

}
