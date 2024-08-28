package heap

type Node struct {
	Char  byte  // Character or symbol
	Freq  int   // Frequency of the character
	Left  *Node // Left child node
	Right *Node // Right child node
}

type HuffmanHeap []*Node

func (h *HuffmanHeap) Insert(node *Node) {
	*h = append(*h, node)

	h.upHeap(len(*h) - 1)
}

func (h *HuffmanHeap) ExtractMin() *Node {
	if len(*h) == 0 {
		return nil
	}

	min := (*h)[0]
	last := len(*h) - 1

	(*h)[0] = (*h)[last]
	*h = (*h)[:last]

	h.downHeap(0)

	return min
}

func (h *HuffmanHeap) upHeap(idx int) {
	parent := (idx - 1) / 2

	for idx > 0 && (*h)[idx].Freq < (*h)[parent].Freq {
		(*h)[idx], (*h)[parent] = (*h)[parent], (*h)[idx]
		idx = parent
		parent = (idx - 1) / 2
	}
}

func (h *HuffmanHeap) downHeap(idx int) {
	lastIdx := len(*h) - 1

	for {
		left := 2*idx + 1
		right := 2*idx + 2
		smallest := idx

		if left <= lastIdx && (*h)[left].Freq < (*h)[smallest].Freq {
			smallest = left
		}
		if right <= lastIdx && (*h)[right].Freq < (*h)[smallest].Freq {
			smallest = right
		}
		if smallest == idx {
			break
		}
		(*h)[idx], (*h)[smallest] = (*h)[smallest], (*h)[idx]
		idx = smallest
	}
}
