package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	word := []string{"kita", "atik", "tika", "aku", "kia", "makan", "kua"}
	wordDetail := wordToLetter(word)
	mapWordDetail := wordGrouping(wordDetail)
	mapWord := getWordOnly(mapWordDetail)
	fmt.Println(mapWord)
}

type WordDetail struct {
	Word   string
	Length int
	Letter map[string]int
	Index  int
}

func wordToLetter(words []string) []WordDetail {
	var wordDetail []WordDetail
	for i, w := range words {
		letters := make(map[string]int)
		for _, l := range strings.Split(w, "") {

			letters[l]++
		}
		wordDetail = append(wordDetail, WordDetail{
			Word:   w,
			Length: len(w),
			Letter: letters,
			Index:  i,
		})
	}
	return wordDetail
}

func wordGrouping(wordDetail []WordDetail) map[string][]WordDetail {
	mapWordDetail := make(map[string][]WordDetail)

	for _, wd := range wordDetail {
		anyAnagram := false
		for k, mwd := range mapWordDetail {
			mwd0 := mwd[0]
			if mwd0.Length == wd.Length && reflect.DeepEqual(mwd0.Letter, wd.Letter) {
				mapWordDetail[k] = append(mapWordDetail[k], wd)
				anyAnagram = true
				break
			}
		}
		if anyAnagram == false {
			mapWordDetail[wd.Word] = []WordDetail{wd}
		}
	}
	return mapWordDetail
}

func getWordOnly(mapWordDetail map[string][]WordDetail) map[int][]string {
	mapWord := make(map[int][]string)
	i := 0
	for _, mwd := range mapWordDetail {
		for _, wd := range mwd {
			mapWord[i] = append(mapWord[i], wd.Word)
		}
		i++
	}
	return mapWord
}
