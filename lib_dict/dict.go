package lib_dict

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Dictionary struct {
	Set map[string]struct{}
}

func NewDictionary() *Dictionary {
	return &Dictionary{Set: make(map[string]struct{})}
}

func (d *Dictionary) AddToDict(dict_buf []byte) {
	var tmp_buf []byte
	for _, char := range string(dict_buf) {
		//fmt.Print("CHAR:", string(char), " ")
		//fmt.Printf(" HEX: %x ", string(char))
		if string(char) == "\n" {
			//fmt.Println()
			d.Add(string(tmp_buf))
			tmp_buf = nil
		} else if string(char) != "\r" {
			tmp_buf = append(tmp_buf, byte(char))
		}
	}
}

func (d *Dictionary) Has(word string) bool {
	_, ok := d.Set[word]
	return ok
}

func (d *Dictionary) Add(word string) {
	strings.TrimSpace(word)
	d.Set[word] = struct{}{}
}

type WordRepo struct {
	mu  sync.Mutex
	Set map[string]struct{}
}

func NewWordRepo() *WordRepo {
	return &WordRepo{Set: make(map[string]struct{})}
}

// if word is in the dictionary add to repo
func (wp *WordRepo) Add(word string, d *Dictionary) {
	wp.mu.Lock()
	defer wp.mu.Unlock()
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
	trimmed := strings.Trim(string(b), "\n\r\t")
	fmt.Fprintf(f, "%s", trimmed)
}
