package csv_process

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"
)

type EthereumPrice struct {
	Price float64
	Date  time.Time
}

func LoadDataFrom(path string) ([]EthereumPrice, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening prices file: %w", err)
	}

	defer file.Close()

	r := csv.NewReader(file)

	var pricePairs []EthereumPrice

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, fmt.Errorf("error reading from CSV file: %w", err)
		}

		parsedDate, err := time.Parse("2006-01-02", record[0])
		if err != nil {
			return nil, fmt.Errorf("cannot parse date: %w", err)
		}

		parsedPrice, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			return nil, fmt.Errorf("cannot parse price: %w", err)
		}

		pricePairs = append(pricePairs, EthereumPrice{
			Price: parsedPrice,
			Date:  parsedDate,
		})
	}

	return pricePairs, nil
}
