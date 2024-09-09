package robotname

import (
	"math/rand/v2"
	"strconv"
)

var NAMES = map[string]bool{}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if len(r.name) == 0 {
		r.name = generateNewName()
	}
	return r.name, nil

}

func generateNewName() string {
	for {
		result := generateAName()
		if isANewName(result) {
			NAMES[result] = true
			return result
		}
	}
}

func generateAName() string {
	result := ""
	for i := 0; i < 2; i++ {
		r1 := rune('A' - 1 + rand.IntN(26))
		result += string(r1)
	}

	number := rand.IntN(999-100) + 100
	result += strconv.Itoa(number)
	return result
}

func (r *Robot) Reset() {
	r.name = ""
	NAMES = make(map[string]bool)
}

func isANewName(name string) bool {
	_, ok := NAMES[name]
	return !ok
}
