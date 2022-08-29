package main

import (
	"log"
	"nymin/lib_dict"
	"nymin/solve"
	"os"
)

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
	dictionary.AddToDict(dict_buf)

	wordRepo := lib_dict.NewWordRepo()
	solve.Solve(string(buffer), wordRepo, dictionary)

	// prints out all words in the picture
	wordRepo.PrintWordsJSON()
}
