package main

import (
	"fmt"
	"os"

	"github.com/MUlt1mate/foursquare-exporter/api"
	"github.com/MUlt1mate/foursquare-exporter/export"
)

func main() {
	token := os.Getenv("FOURSQUARE_TOKEN")

	if token == "" {
		fmt.Fprintln(os.Stderr, "FOURSQUARE_TOKEN environment variable is required")
		os.Exit(1)
	}

	client := api.NewClient(token)

	checkins, err := client.GetAllCheckins()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error fetching checkins: %v\n", err)
		os.Exit(1)
	}

	if err := export.WriteCSV("checkins.csv", checkins); err != nil {
		fmt.Fprintf(os.Stderr, "Error writing CSV: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Exported %d checkins to checkins.csv\n", len(checkins))
}
