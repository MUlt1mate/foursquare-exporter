# Foursquare Exporter

Exports your Foursquare check-in history to a CSV file in Google Calendar import format.

## Project Structure

```
├── main.go          # Entry point, configuration
├── api/client.go    # Foursquare API client with pagination
├── model/checkin.go # Data models and CSV record conversion
├── export/csv.go    # CSV file writer
```

## Prerequisites

- Go 1.25+
- Foursquare OAuth token — obtain one at https://developer.foursquare.com/docs/explore
    - or get to https://swarmapp.com/history and take it from request in dev tools

## Usage

```sh
FOURSQUARE_TOKEN=your_token_here go run main.go
```

This will create a `checkins.csv` file in the current directory.

## CSV Format

The output CSV is compatible with Google Calendar import:

| Column        | Example                                           |
|---------------|---------------------------------------------------|
| Subject       | Venue name                                        |
| Start Date    | 2024/01/15                                        |
| Start Time    | 02:30:00 PM                                       |
| End Date      | 2024/01/15                                        |
| End Time      | 03:30:00 PM                                       |
| All Day Event | False                                             |
| Description   | https://foursquare.com/v/foursquare-hq/{venue_id} |
| Location      | Formatted address                                 |
| Private       | True                                              |

Each check-in is exported as a 1-hour event.

## Building

```sh
go build -o foursquare-exporter .
./foursquare-exporter
```
