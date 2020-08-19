package goutils

import (
	"encoding/csv"
	"io"
	"log"
)

func ReadCsv(f io.Reader) <-chan map[string]string {
	r := csv.NewReader(f)
	r.ReuseRecord = true
	r.TrimLeadingSpace = true

	ch := make(chan map[string]string)

	go func() {
		defer close(ch)

		withHeader := true
		var headers []string

		for {
			record, err := r.Read()
			if err != nil {
				if err == io.EOF {
					log.Println("FINISH CSV data")
					return
				} else {
					log.Fatal(err)
					return
				}
			}
			if withHeader {
				headers = make([]string, len(record))
				copy(headers, record)
				withHeader = false
				continue
			}

			data := map[string]string{}

			for idx, header := range headers {
				data[header] = record[idx]
			}

			ch <- data
		}
	}()

	return ch
}
