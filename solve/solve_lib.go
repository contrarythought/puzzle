package solve

import (
	"fmt"
	"nymin/lib_dict"
	"strings"
	"sync"
)

const (
	// CHANGE TO 255
	MAX_BOUNDARY = 15
)

// TODO
func Solve(diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary) {
	var wg sync.WaitGroup
	for i := 0; i < MAX_BOUNDARY; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var word []rune
			AnalyzeFront(i, diagram, wp, d, word)
			word = nil
			AnalyzeBack(i, diagram, wp, d, word)
			word = nil
			AnalyzeVertLR(i, diagram, wp, d, word)
			word = nil
			AnalyzeVertRL(i, diagram, wp, d, word)
			word = nil
			AnalyzeDown(i, diagram, wp, d, word)
			word = nil
			AnalyzeUp(i, diagram, wp, d, word)
			word = nil
		}(i)
	}
	wg.Wait()
}

func AnalyzeFront(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {
	if index < 0 || index >= MAX_BOUNDARY {
		return
	}
	if (index+1)%15 == 0 {
		word = append(word, rune(diagram[index]))
		wp.Add(strings.TrimSpace(string(word)), d)
		index++
		return
	} else {
		word = append(word, rune(diagram[index]))
		fmt.Println(string(word))
		wp.Add(strings.TrimSpace(string(word)), d)
		index++
		AnalyzeFront(index, diagram, wp, d, word)
	}
}

func AnalyzeBack(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {
	if index < 0 || index >= MAX_BOUNDARY {
		return
	}
	if index%15 == 0 {
		word = append(word, rune(diagram[index]))
		wp.Add(strings.TrimSpace(string(word)), d)
		index--
		return
	} else {
		word = append(word, rune(diagram[index]))
		wp.Add(strings.TrimSpace(string(word)), d)
		index--
		AnalyzeBack(index, diagram, wp, d, word)
	}
}

// TODO
func AnalyzeVertLR(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {
	if index < 0 || index >= MAX_BOUNDARY {
		return
	}
	word = append(word, rune(diagram[index]))
	wp.Add(strings.TrimSpace(string(word)), d)
	index += 16
	AnalyzeVertLR(index, diagram, wp, d, word)
}

// TODO
func AnalyzeVertRL(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {
	if index < 0 || index >= MAX_BOUNDARY {
		return
	}
	word = append(word, rune(diagram[index]))
	wp.Add(strings.TrimSpace(string(word)), d)
	index -= 16
	AnalyzeVertRL(index, diagram, wp, d, word)
}

func AnalyzeDown(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {
	if index < 0 || index >= MAX_BOUNDARY {
		return
	}
	word = append(word, rune(diagram[index]))
	wp.Add(strings.TrimSpace(string(word)), d)
	index += 15
	AnalyzeDown(index, diagram, wp, d, word)
}

func AnalyzeUp(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {
	if index < 0 || index >= MAX_BOUNDARY {
		return
	}
	word = append(word, rune(diagram[index]))
	wp.Add(strings.TrimSpace(string(word)), d)
	index -= 15
	AnalyzeUp(index, diagram, wp, d, word)
}
