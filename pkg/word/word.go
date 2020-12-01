package word

import (
	"fmt"
	"math/rand"
	"regexp"
)

var (
	vowels = []rune{
		'a', 'e', 'i', 'o', 'u', 'y',
	}
	consonants = []rune{
		'b', 'c', 'd', 'f', 'g', 'h',
		'j', 'k', 'l', 'm', 'n', 'p',
		'q', 'r', 's', 't', 'v', 'w',
		'x', 'z',
	}
	mutantMakers = []func(string) (string, bool){
		changeVowel,
		changeConsonant,
		removeLetter,
		addVowel,
		addConsonant,
	}
)

func MutateTimes(w string, n int) ([]string, error) {
	mutants := map[string]bool{}
	i := 0
	for {
		i++
		if i == 1000 {
			return nil, fmt.Errorf(
				"couldn't create %v mutations of %v after %v attempts",
				n, w, i)
		}
		m, ok := mutantMakers[rand.Intn(len(mutantMakers))](w)
		if !ok {
			continue
		}
		if _, have := mutants[m]; have {
			continue
		}
		if i == 1000 {
			return nil, fmt.Errorf(
				"couldn't create %v mutations of %v after %v attempts",
				n, w, i)
		}
		mutants[m] = true
		if len(mutants) == n {
			ms := []string{}
			for m := range mutants {
				ms = append(ms, m)
			}
			return ms, nil
		}
	}
}

func changeVowel(w string) (string, bool) {
	re := regexp.MustCompile(`[aeiouy]`)
	matches := re.FindAllIndex([]byte(w), -1)
	if len(matches) == 0 {
		return "", false
	}
	i := rand.Intn(len(matches))
	match := matches[i]
	for {
		mutant := []rune(w)
		current := mutant[match[0]]
		candidate := vowels[rand.Intn(len(vowels))]
		if candidate != current {
			mutant[match[0]] = candidate
			return string(mutant), true
		}
	}
}

func changeConsonant(w string) (string, bool) {
	if len(w) < 3 {
		return "", false // too short
	}
	// Just change the middle consonants so the word still looks close.
	middle := []byte(w)[1 : len(w)-2]
	re := regexp.MustCompile(`[bcdfghjklmnpqrstvwxz]`)
	matches := re.FindAllIndex(middle, -1)
	if len(matches) == 0 {
		return "", false
	}
	i := rand.Intn(len(matches)) + 1
	for {
		mutant := []rune(w)
		current := mutant[i]
		candidate := consonants[rand.Intn(len(consonants))]
		if candidate != current {
			mutant[i] = candidate
			return string(mutant), true
		}
	}
}

func removeLetter(w string) (string, bool) {
	if len(w) < 2 {
		return "", false // too short
	}
	i := rand.Intn(len(w))
	mutant := append([]rune(w)[0:i], []rune(w)[i+1:len(w)]...)
	return string(mutant), true
}

func addVowel(w string) (string, bool) {
	i := rand.Intn(len(w))
	newcomer := vowels[rand.Intn(len(vowels))]
	mutant := []rune(w)[0:i]
	mutant = append(mutant, newcomer)
	mutant = append(mutant, []rune(w)[i:len(w)]...)
	return string(mutant), true
}

func addConsonant(w string) (string, bool) {
	i := rand.Intn(len(w))
	newcomer := consonants[rand.Intn(len(consonants))]
	mutant := []rune(w)[0:i]
	mutant = append(mutant, newcomer)
	mutant = append(mutant, []rune(w)[i:len(w)]...)
	return string(mutant), true
}
