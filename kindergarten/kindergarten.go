package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

// Define the Garden type here.
type Garden struct {
	plants map[string]Cups
}

type Cups struct {
	cups []string
}

func NewGarden(diagram string, children []string) (*Garden, error) {
	sortedChildren := sortChildren(children)

	result := map[string]Cups{}
	split := strings.Split(diagram, "\n")
	err := verifyInput(split, sortedChildren)
	if err != nil {
		return nil, err
	}
	for index, child := range sortedChildren {
		var cups []string

		_, ok := result[child]
		if ok {
			return nil, errors.New("Invalid")
		}

		plant1, b := getPlant(split[1][index*2])
		if b == false {
			return nil, errors.New("Invalid")
		}
		cups = append(cups, plant1)

		plant2, b := getPlant(split[1][index*2+1])
		if b == false {
			return nil, errors.New("Invalid")
		}
		cups = append(cups, plant2)

		plant3, b := getPlant(split[2][index*2])
		if b == false {
			return nil, errors.New("Invalid")
		}
		cups = append(cups, plant3)

		plant4, b := getPlant(split[2][index*2+1])
		if b == false {
			return nil, errors.New("Invalid")
		}
		cups = append(cups, plant4)

		result[child] = Cups{
			cups: cups,
		}
	}

	return &Garden{
		plants: result,
	}, nil
}

func sortChildren(children []string) []string {
	var sortedChildren []string
	sortedChildren = append(sortedChildren, children...)
	sort.Strings(sortedChildren)
	return sortedChildren
}

func verifyInput(split []string, children []string) error {
	if len(split) != 3 {
		return errors.New("Invalid")
	}
	if len(split[1]) != len(split[2]) {
		return errors.New("Invalid")
	}
	if len(split[1]) != len(children)*2 {
		return errors.New("Invalid")
	}

	return nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	plants := g.plants
	cups, ok := plants[child]
	if ok {
		return cups.cups, true
	}
	return nil, false
}

func getPlant(encoding uint8) (string, bool) {
	encodingMap := map[uint8]string{
		'G': "grass",
		'V': "violets",
		'R': "radishes",
		'C': "clover",
	}
	value, ok := encodingMap[encoding]
	return value, ok
}
