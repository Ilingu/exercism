package kindergarten

import (
	"errors"
	"sort"
	"strings"
)

var letterToName = map[rune]string{'V': "violets", 'R': "radishes", 'C': "clover", 'G': "grass"}

type Garden struct {
	diagram       string
	ClassChildren []string
}

func NewGarden(diagram string, RawChildren []string) (*Garden, error) {
	if strings.ToUpper(diagram) != diagram || diagram[0] != 10 { //Check if first byte is "\n"
		return nil, errors.New("wrong format")
	}

	seen := map[string]bool{}
	for _, child := range RawChildren {
		if seen[child] {
			return nil, errors.New("duplicate name")
		}
		seen[child] = true
	}

	diagram = strings.TrimPrefix(diagram, "\n")
	rows := strings.Split(diagram, "\n")
	if len(rows) != 2 {
		return nil, errors.New("invalid number of rows")
	} else if len(rows[0]) != len(rows[1]) {
		return nil, errors.New("mismatched rows")
	} else if len(rows[0])%2 != 0 || len(rows[1])%2 != 0 {
		return nil, errors.New("odd number of cups")
	}

	SortedChildren := append([]string{}, RawChildren...)
	sort.Strings(SortedChildren) // Sorting Children in alphabetical order

	return &Garden{diagram: diagram, ClassChildren: SortedChildren}, nil
}

func (g *Garden) Plants(child string) ([]string, bool) {
	if !strings.Contains(strings.Join(g.ClassChildren, ", "), child) {
		return nil, false
	}

	OrderedDiagram := g.OrderDiagram()
	i, _ := IndexOfChild(child, g.ClassChildren)

	childPlants, namedChildPlants := OrderedDiagram[i], []string{}
	for _, plant := range childPlants {
		plantName := letterToName[plant]
		namedChildPlants = append(namedChildPlants, plantName)
	}

	return namedChildPlants, true
}

/* Helper Functions */

/* Convert the diagram grid into a array of string, representing children plants (array ordered by alphabetical order).
e.g: Index 0 is always the 4 Alice's plants...*/
func (g *Garden) OrderDiagram() []string {
	rows, rowChildPlant := strings.Split(g.diagram, "\n"), [2][]string{{}, {}}
	for i, row := range rows {
		for j := range row {
			if (j+1)%2 == 0 {
				rowChildPlant[i] = append(rowChildPlant[i], row[j-1:j+1])
			}
		}
	}

	childrenPlants := []string{}
	for i := range rowChildPlant[0] {
		childrenPlants = append(childrenPlants, rowChildPlant[0][i]+rowChildPlant[1][i])
	}
	return childrenPlants
}

func IndexOfChild(child string, ClassChildren []string) (int, error) {
	for i, InClassChild := range ClassChildren {
		if child == InClassChild {
			return i, nil
		}
	}
	return 0, errors.New("child not in this class")
}
