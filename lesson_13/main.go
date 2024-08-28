package main

import (
	"bufio"
	"fmt"
	"lesson13/huffman_compression"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter something: ")

	data, err := reader.ReadString('\n')
	if err != nil {
		os.Exit(1)
	}

	data = strings.TrimSpace(data)

	// Calculate and print the size of uncompressed data in bytes
	uncompressedSizeBytes := len(data) // Number of bytes in the UTF-8 encoded string
	uncompressedSizeBits := uncompressedSizeBytes * 8
	fmt.Printf("Uncompressed size: %d bytes (%d bits)\n", uncompressedSizeBytes, uncompressedSizeBits)

	// Compress the data
	encodedData, root, codes := huffman_compression.Compress(data)

	fmt.Println(codes)
	// Calculate and print the size of compressed data in bits
	compressedSizeBits := len(encodedData) // encodedData length is in bits since it's a binary string
	fmt.Printf("Compressed size: %d bits\n", compressedSizeBits)

	// Calculate the compression percentage
	compressionPercentage := 100 - (float64(compressedSizeBits) / float64(uncompressedSizeBits) * 100)
	fmt.Printf("Compression ratio: %.2f%%\n", compressionPercentage)

	fmt.Println("Encoded:", encodedData)

	decodedData := huffman_compression.Decompress(encodedData, root)
	fmt.Println("Decoded:", decodedData)

	decodedDataCodes := huffman_compression.DecompressUsingCodes(encodedData, codes)
	fmt.Println("Decoded using codes:", decodedDataCodes)
}
