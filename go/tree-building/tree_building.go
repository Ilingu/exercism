package tree

import (
	"errors"
)

type Record struct {
	ID     int
	Parent int
}

type Node struct {
	ID       int
	Children []*Node
}

func Build(records []Record) (*Node, error) {
	if len(records) <= 0 {
		return nil, nil
	}

	root := Node{ID: 0}
	sortRecords(records)

	if records[0].ID == 0 {
		if records[0].Parent != 0 {
			return nil, errors.New("invalid root node")
		}
	} else {
		return nil, errors.New("no root node")
	}

	records = records[1:] // shift first elem (root node)
	IdSeen, lastIdSeen := map[int]bool{}, 0
	for _, record := range records {
		if record.ID == 0 {
			return nil, errors.New("duplicate root node")
		}
		if record.Parent > record.ID {
			return nil, errors.New("higher id parent of lower id")
		}
		if IdSeen[record.ID] {
			return nil, errors.New("duplicate node")
		}

		IdSeen[record.ID] = true
		lastIdSeen = record.ID

		// 1. Find Record.Parent_ID == Node.ID in Root Node
		// 2. append to the found Node's Children the Record.ID
		children, found := FindInNodes(&root, record.Parent)
		if !found {
			return nil, errors.New("invalid tree")
		}

		*children = append((*children), &Node{ID: record.ID})
	}

	// Check integrity
	for i := 1; i <= lastIdSeen; i++ {
		if !IdSeen[i] {
			return nil, errors.New("non continuous")
		}
	}

	return &root, nil
}

func FindInNodes(root *Node, IdToFind int) (*[]*Node, bool) {
	if root.ID == IdToFind {
		return &root.Children, true
	}

	for _, child := range root.Children {
		child, found := FindInNodes(child, IdToFind)
		if !found {
			continue
		}

		return child, true
	}
	return nil, false
}

func sortRecords(arr []Record) {
	for i := range arr {
		for j := 0; j < len(arr)-i-1; j++ {
			// Sorting By Parent ID
			if arr[j].Parent > arr[j+1].Parent {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
			// Sorting By Parent's Children ID
			if arr[j].ID > arr[j+1].ID && arr[j].Parent == arr[j+1].Parent {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}

		// Assign ChildID
		// if i+1 <= len(arr)-1 && arr[i].Parent == arr[i+1].Parent {
		// 	arr[i].ChildID = counter
		// 	arr[i+1].ChildID = counter + 1
		// 	counter++
		// } else {
		// 	counter = 0
		// }
	}
}

/*

	{ID: 1, Parent: 0},
	{ID: 2, Parent: 0},
	{ID: 4, Parent: 1},
	{ID: 3, Parent: 2},
	{ID: 5, Parent: 2},
	{ID: 6, Parent: 2},

*/

/*

root (ID: 0, parent ID: 0)

	+-- 0child1 (ID: 1, parent ID: 0)
			+-- 1grandchild1 (ID: 5, parent ID: 1)
						+-- 2grandchild1 (ID: 6, parent ID: 5)

	+-- 0child2 (ID: 2, parent ID: 0)
			+-- 1grandchild1 (ID: 3, parent ID: 2)
						+-- 2grandchild1 (ID: 4, parent ID: 3)


--> [{1 0} {2 0} {5 1} {3 2} {4 3} {6 5}]
*/
