package main

import (
	"log"
	"nymin/lib_dict"
	"nymin/solve"
	"os"
)

func AddToDict(dict_buf []byte, dictionary *lib_dict.Dictionary) {
	var tmp_buf []byte
	for _, char := range string(dict_buf) {
		if string(char) == "\n" {
			dictionary.Add(string(tmp_buf))
			tmp_buf = nil
		} else {
			tmp_buf = append(tmp_buf, byte(char))
		}
	}
}

func main() {
	buffer, err := os.ReadFile("diagram.txt")
	if err != nil {
		log.Fatal(err)
	}

	dict_buf, err := os.ReadFile("dictionary.txt")
	if err != nil {
		log.Fatal(err)
	}

	dictionary := lib_dict.NewDictionary()
	AddToDict(dict_buf, dictionary)

	wordRepo := lib_dict.NewWordRepo()
	solve.Solve(string(buffer), wordRepo, dictionary)
}
