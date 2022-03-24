package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

type Jsn struct {
	Name   string  `json:"name"`
	Price  float64 `json:"price"`
	Rating float64 `json:"rating"`
}

func main() {
	file, err := os.Open("data.json")
	if err != nil {
		panic(err.Error())
	}
	dataBytes, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	item := Jsn{
		Name:   "",
		Price:  0,
		Rating: 0,
	}
	r := strings.NewReader(string(dataBytes))
	dec := json.NewDecoder(r)
	for {
		tok, err := dec.Token()
		if err != nil {
			break
		}
		if tok == nil {
			break
		}
		switch tp := tok.(type) {
		case json.Delim:
			str := tp.String()
			if str == "[" || str == "{" {
				for dec.More() {
					var u Jsn
					err := dec.Decode(&u)
					if err != nil {
						break
					} else {
						if u.Price > item.Price && u.Rating > item.Rating {
							item = u
						}
					}
				}
			}
		}
	}
	fmt.Printf("%v", item)

	// CSV file

	f, err := os.Open("data.csv")
	if err != nil {
		return
	}
	defer f.Close()

	lines, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return
	}

	biggest := lines[1]

	for i := 2; i < len(lines); i++ {
		line := lines[i]
		lineId, _ := strconv.Atoi(line[0])
		biggestId, _ := strconv.Atoi(biggest[0])

		if lineId > biggestId {
			biggest = line
		}
	}

	fmt.Println(biggest)
}
