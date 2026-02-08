package export

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/MUlt1mate/foursquare-exporter/model"
)

var header = []string{
	"Subject", "Start Date", "Start Time", "End Date", "End Time",
	"All Day Event", "Description", "Location", "Private",
}

func WriteCSV(filename string, checkins []model.Checkin) error {
	f, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("creating file: %w", err)
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	if err := w.Write(header); err != nil {
		return fmt.Errorf("writing header: %w", err)
	}

	for _, c := range checkins {
		record := model.CheckinToCSVRecord(c)
		if record == nil {
			continue
		}
		if err := w.Write(record.ToSlice()); err != nil {
			return fmt.Errorf("writing record: %w", err)
		}
	}

	return nil
}
