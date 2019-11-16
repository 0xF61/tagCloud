package main

import (
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strings"
)

func getPage(url string, w http.ResponseWriter, r *http.Request, minFrequency int) []Tag {
	resp, err := http.Get(url)
	// handle the error if there is one
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		panic(err)
	}
	// do this now so it won't be forgotten
	defer resp.Body.Close()
	// reads html as a slice of bytes
	html, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		panic(err)
	}

	regx := regexp.MustCompile("(?s)<body>.*</body>")
	h := regx.FindString(string(html))
	h = strings.ToLower(h)

	wordMap := wordCount(h)
	var tags []Tag

	// add Tag to tags
	for i, j := range wordMap {
		if j > minFrequency {
			tags = append(tags, Tag{Title: i, Freq: j})
		}
	}
	// Burada tags atanmış bitmiş durumda

	// SORT
	keys := make([]string, 0, len(tags))
	for _, tag := range tags {
		keys = append(keys, tag.Title)
	}
	sort.Strings(keys)
	// Keyler sort edildi

	max := findMax(tags)

	var sortedTags []Tag
	for _, tag := range keys {
		sortedTags = append(sortedTags, Tag{Title: tag, Freq: findFreq(tag, tags)})
	}

	fontMin := 10
	fontMax := 20
	for i, tag := range sortedTags {
		// Calculate fontsize for each tag
		sortedTags[i].Size = int((float64(tag.Freq)/float64(max))*float64((fontMax-fontMin)) + float64(fontMin))
	}
	return sortedTags
}

func findFreq(tag string, tags []Tag) int {
	for _, j := range tags {
		if tag == j.Title {
			return j.Freq
		}
	}
	return 42
}

func wordCount(str string) map[string]int {
	wordList := strings.Fields(str)
	counts := make(map[string]int)
	for _, word := range wordList {
		if len(word) <= 3 || checkBanned(word) {
			continue
		}
		_, ok := counts[word]
		if ok {
			counts[word] += 1
		} else {
			counts[word] = 1
		}
	}

	return counts
}

func checkBanned(word string) bool {
	bannedList := []string{"\"", "<", ">", "'", "=", "/", "(", ";"}
	for _, banned := range bannedList {
		if strings.ContainsAny(word, banned) {
			return true
		}
	}
	return false
}

func findMax(tags []Tag) int {
	max := 0
	for _, tag := range tags {
		if tag.Freq > max {
			max = tag.Freq
		}
	}
	return max
}
