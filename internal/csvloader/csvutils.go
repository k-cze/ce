package csvloader

import (
	"encoding/csv"
	"os"
)

func ProcessCSV[R any, D interface{ Append(*R) }](p *Processor[R, D]) error {
	f, err := os.Open(p.Filename())
	if err != nil {
		return err
	}
	defer f.Close()

	reader := csv.NewReader(f)
	records, err := reader.ReadAll()
	if err != nil {
		return err
	}

	for _, line := range records[1:] {
		if err := p.processCSVLine(line); err != nil {
			return err
		}
	}

	return nil
}
