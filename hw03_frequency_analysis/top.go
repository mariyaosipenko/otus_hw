package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

func Top10(str string) []string {
	w := strings.Fields(str)
	if len(w) == 0 {
		return nil
	}

	wordFrequency := make(map[string]int, len(w))
	for _, v := range w {
		pv := strings.ToLower(strings.Trim(v, "!¡,.'-"))
		if len(pv) > 0 {
			wordFrequency[pv]++
		}
	}
	keys := make([]string, 0, len(wordFrequency))

	for key := range wordFrequency {
		keys = append(keys, key)
	}

	sort.Slice(keys, func(i, j int) bool {
		if wordFrequency[keys[i]] == wordFrequency[keys[j]] {
			return keys[i] < keys[j]
		}
		return wordFrequency[keys[i]] > wordFrequency[keys[j]]
	})

	var cut int
	if len(keys) < 10 {
		cut = len(keys)
	} else {
		cut = 10
	}

	return keys[:cut]
}
