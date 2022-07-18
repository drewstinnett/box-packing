package packit

import (
	"errors"
	"sort"
)

type Packit struct {
	Boxes map[int]*Box
	Items []BoxItem
}

// QuickPackWithMaxSize packs all items into boxes with a maximum size. Quick
// just sorts from largest to smallest and does a single pass through
func (p *Packit) QuickPackWithMaxSize(maxSize int64) error {
	// Sort largest to smallest
	sort.Sort(sort.Reverse(BoxItemsBySize(p.Items)))
	// Add at least 1 box
	p.AddBoxWithCapacity(maxSize)

	for _, item := range p.Items {
		// Implementation requires that maxSize is greater than or equal to the size of the largest file
		// If maxSize == 0, everything goes in the same suitcase
		if maxSize == 0 {
			p.Boxes[0].AddItem(item)
		} else {
			if item.GetSize() > maxSize {
				return errors.New("index containes at least one file that is too large")
			}
			if !p.sortItemIntoExistingBox(item) {
				p.sortItemIntoNewBox(item, maxSize)
			}
		}
	}
	return nil
}

func (p *Packit) SlowPackWithMaxSize(maxSize int64) error {
	// Sort largest to smallest
	sort.Sort(sort.Reverse(BoxItemsBySize(p.Items)))
	// Add at least 1 box
	p.AddBoxWithCapacity(maxSize)

	for _, item := range p.Items {
		// Implementation requires that maxSize is greater than or equal to the size of the largest file
		// If maxSize == 0, everything goes in the same suitcase
		if maxSize == 0 {
			p.Boxes[0].AddItem(item)
		} else {
			if item.GetSize() > maxSize {
				return errors.New("index containes at least one file that is too large")
			}
			if !p.sortItemIntoExistingBox(item) {
				p.sortItemIntoNewBox(item, maxSize)
			}
		}
	}
	return nil
}

// Return indicates if there was a box to insert into
func (p *Packit) sortItemIntoExistingBox(item BoxItem) bool {
	// for index, box := range p.Boxes {
	for i := 0; i < len(p.Boxes); i++ {
		if item.GetSize() <= p.Boxes[i].Remaining() {
			p.Boxes[i].AddItem(item)
			return true
		}
	}
	return false
}

func (p *Packit) sortItemIntoNewBox(item BoxItem, c int64) int {
	newIndex := p.AddBoxWithCapacity(c)
	p.Boxes[newIndex].AddItem(item)
	return newIndex
}

func (p *Packit) AddItem(item BoxItem) error {
	p.Items = append(p.Items, item)
	return nil
}

// add a new box to the map of boxes, incrementing the index
func (p *Packit) AddBoxWithCapacity(c int64) int {
	curBoxLen := len(p.Boxes)
	if curBoxLen == 0 {
		p.Boxes = map[int]*Box{
			0: {
				Capacity: c,
			},
		}
	} else {
		p.Boxes[curBoxLen] = &Box{
			Capacity: c,
		}
	}
	return curBoxLen
}
