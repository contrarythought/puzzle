package solve

import (
	"nymin/lib_dict"
	"sync"
)

const (
	MAX_CHAR = 255
)

// TODO
func Solve(diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary) {
	var wg sync.WaitGroup
	for i := 0; i < MAX_CHAR; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			var word []rune
			AnalyzeFront(i, diagram, wp, d, word)
			word = nil
			AnalyzeBack(i, diagram, wp, d, word)
			word = nil
			AnalyzeVert(i, diagram, wp, d, word)
			word = nil
		}(i)
	}
	wg.Wait()
}

func AnalyzeFront(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {
	if index < 0 || index >= MAX_CHAR {
		return
	}
	// TODO - figure this condition out...maybe use array of key indices?
	if index > 0 && index%15 == 0 {
		return
	} else {
		word = append(word, rune(diagram[index]))
		wp.Add(string(word), d)
		index++
		AnalyzeFront(index, diagram, wp, d, word)
	}
}

func AnalyzeBack(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {

}

func AnalyzeVert(index int, diagram string, wp *lib_dict.WordRepo, d *lib_dict.Dictionary, word []rune) {

}
