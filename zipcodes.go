// Package zipcodes uses https://www.data.gouv.fr/fr/datasets/base-officielle-des-codes-postaux/
// Second init() (commented out) uses https://sql.sh/736-base-donnees-villes-francaises
// See https://github.com/samueltardieu/AusweisBot/blob/master/src/main/scala/ZipCodes.scala

package zipcodes

import (
	"bufio"
	"encoding/csv"
	"io"
	"log"
	"os"
)

var (
	MapZipToCities = make(map[string][]string)
	MapCityToZips  = make(map[string][]string)
)

func init() {
	fdesc, err := os.Open("laposte_hexasmal.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fdesc.Close()

	reader := csv.NewReader(bufio.NewReader(fdesc))
	reader.Comma = ';'
	_, _ = reader.Read() // Discard header line

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		zip, city := record[2], record[3]
		if city == "" {
			city = record[4]
		}
		MapZipToCities[zip] = append(MapZipToCities[zip], city)
		MapCityToZips[city] = append(MapCityToZips[city], zip)
	}
}

/*func init() {
	fdesc, err := os.Open("villes_france.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer fdesc.Close()

	reader := csv.NewReader(bufio.NewReader(fdesc))

	for {
		record, err := reader.Read()
		if err == io.EOF  {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		for _, zip := range strings.Split(record[8], "-") {
			MapZipToCities[zip] = append(MapZipToCities[zip], record[5])
			MapCityToZips[record[5]] = append(MapCityToZips[record[5]], zip)
		}
	}
}*/
