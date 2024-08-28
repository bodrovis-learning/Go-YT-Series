package huffman_compression

import (
	"fmt"
	"lesson13/heap"
	"strings"
)

func Decompress(encoded string, root *heap.Node) string {
	if len(encoded) == 0 {
		return ""
	}

	var decoded strings.Builder
	currentNode := root
	// aaaaaaaaaaaaaaaaaaaaa
	// 000000000000000000000
	if currentNode.Left == nil && currentNode.Right == nil {
		// In this case, the encoded string is just a repetition of the single character
		for range encoded {
			decoded.WriteByte(currentNode.Char)
		}
		return decoded.String()
	}

	for _, bit := range encoded {
		if bit == '0' {
			currentNode = currentNode.Left
		} else {
			currentNode = currentNode.Right
		}

		if currentNode.Left == nil && currentNode.Right == nil {
			decoded.WriteByte(currentNode.Char)
			currentNode = root
		}
	}

	return decoded.String()
}

func DecompressUsingCodes(encodedData string, codes map[byte]string) string {
	invertedCodes := make(map[string]byte)
	for char, code := range codes {
		invertedCodes[code] = char
	}

	var decoded strings.Builder
	currentCode := ""

	for _, bit := range encodedData {
		currentCode += string(bit)

		if char, found := invertedCodes[currentCode]; found {
			decoded.WriteByte(char) // Append the corresponding character to the decoded output
			currentCode = ""        // Reset currentCode for the next character
		}
	}

	return decoded.String()
}

func Compress(data string) (string, *heap.Node, map[byte]string) {
	if len(data) == 0 {
		return "", nil, nil
	}

	freqMap := make(map[byte]int)
	for i := 0; i < len(data); i++ {
		freqMap[data[i]]++
	}

	fmt.Println(freqMap)

	root := buildHuffmanTree(freqMap)
	codes := getHuffmanCodes(root)
	encodedData := encode(data, codes)

	return encodedData, root, codes
}

func buildHuffmanTree(freqMap map[byte]int) *heap.Node {
	h := &heap.HuffmanHeap{}

	for char, freq := range freqMap {
		h.Insert(&heap.Node{Char: char, Freq: freq})
	}

	for len(*h) > 1 {
		left := h.ExtractMin()
		right := h.ExtractMin()

		merged := &heap.Node{
			Freq:  left.Freq + right.Freq,
			Left:  left,
			Right: right,
		}

		h.Insert(merged)
	}

	return h.ExtractMin()
}

func getHuffmanCodes(root *heap.Node) map[byte]string {
	codes := make(map[byte]string)

	generateCodes(root, "", codes)

	return codes
}

func generateCodes(node *heap.Node, prefix string, codeMap map[byte]string) {
	if node == nil {
		return
	}

	if node.Left == nil && node.Right == nil {
		if prefix == "" {
			prefix = "0" // Assign "0" if this is the only node
		}
		codeMap[node.Char] = prefix
	}

	generateCodes(node.Left, prefix+"0", codeMap)
	generateCodes(node.Right, prefix+"1", codeMap)
}

func encode(data string, codes map[byte]string) string {
	var encoded strings.Builder

	for i := 0; i < len(data); i++ {
		encoded.WriteString(codes[data[i]])
	}

	return encoded.String()
}
