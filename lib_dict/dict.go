package lib_dict

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

type Dictionary struct {
	Set map[string]struct{}
}

func NewDictionary() *Dictionary {
	return &Dictionary{Set: make(map[string]struct{})}
}

func (d *Dictionary) Has(word string) bool {
	_, ok := d.Set[word]
	return ok
}

func (d *Dictionary) Add(word string) {
	d.Set[word] = struct{}{}
}

type WordRepo struct {
	Set map[string]struct{}
}

func NewWordRepo() *WordRepo {
	return &WordRepo{Set: make(map[string]struct{})}
}

// if word is in the dictionary add to repo
func (wp *WordRepo) Add(word string, d *Dictionary) {
	if d.Has(word) {
		wp.Set[word] = struct{}{}
	}
}

func (wp *WordRepo) PrintWordsJSON() {
	f, err := os.Create("clues.json")
	if err != nil {
		log.Fatal(err)
	}
	b, err := json.MarshalIndent(wp.Set, "", "	")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Fprintf(f, "%s", string(b))
}
