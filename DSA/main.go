package main

import (
	"encoding/csv"
	"log"
	"os"
	"strconv"
	"time"
)

func readCsvFile(filePath string) []int {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	values := make([]int, 0, len(records[0]))
	var v int
	for _, i := range records[0] {
		v, _ = strconv.Atoi(i)
		values = append(values, v)
	}
	return values
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}

func main() {
	defer timeTrack(time.Now(), "main")

	records := readCsvFile("./values.csv")

}
