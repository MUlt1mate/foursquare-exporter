package model

import "time"

type CheckinsResponse struct {
	Response struct {
		Checkins struct {
			Items []Checkin `json:"items"`
		} `json:"checkins"`
	} `json:"response"`
}

type Checkin struct {
	CreatedAt int64  `json:"createdAt"`
	Venue     *Venue `json:"venue"`
}

type Venue struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Location struct {
		FormattedAddress []string `json:"formattedAddress"`
	} `json:"location"`
}

type CSVRecord struct {
	Subject     string
	StartDate   string
	StartTime   string
	EndDate     string
	EndTime     string
	AllDayEvent string
	Description string
	Location    string
	Private     string
}

func CheckinToCSVRecord(c Checkin) *CSVRecord {
	if c.Venue == nil {
		return nil
	}
	t := time.Unix(c.CreatedAt, 0)
	tEnd := t.Add(time.Hour)

	return &CSVRecord{
		Subject:     c.Venue.Name,
		StartDate:   t.Format("2006/01/02"),
		StartTime:   t.Format("03:04:05 PM"),
		EndDate:     tEnd.Format("2006/01/02"),
		EndTime:     tEnd.Format("03:04:05 PM"),
		AllDayEvent: "False",
		Description: "https://foursquare.com/v/foursquare-hq/" + c.Venue.ID,
		Location:    joinAddress(c.Venue.Location.FormattedAddress),
		Private:     "True",
	}
}

func (r CSVRecord) ToSlice() []string {
	return []string{
		r.Subject,
		r.StartDate,
		r.StartTime,
		r.EndDate,
		r.EndTime,
		r.AllDayEvent,
		r.Description,
		r.Location,
		r.Private,
	}
}

func joinAddress(parts []string) string {
	var result string
	for i, p := range parts {
		if i > 0 {
			result += " "
		}
		result += p
	}
	return result
}
