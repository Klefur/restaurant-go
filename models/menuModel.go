package models

import (
	"time"
	"encoding/json"
	"gorm.io/gorm"
)

// Menu struct
type Menu struct {
	gorm.Model
	Name        *string		`json:"name" gorm:"not null"`
	Category 	*string		`json:"category" gorm:"not null"`
	Foods 		*[]Food		`json:"foods"`
	Start_date 	*time.Time	`json:"start_date"`
	End_date 	*time.Time	`json:"end_date"`
}

// Custom unmarshal function for Menu
func (m *Menu) UnmarshalJSON(data []byte) error {
    type Alias Menu
    aux := &struct {
        Start_date *string `json:"start_date"`
        End_date   *string `json:"end_date"`
        *Alias
    }{
        Alias: (*Alias)(m),
    }

    if err := json.Unmarshal(data, &aux); err != nil {
        return err
    }

    var err error
    if aux.Start_date != nil {
        m.Start_date, err = parseDate(*aux.Start_date)
        if err != nil {
            return err
        }
    }

    if aux.End_date != nil {
        m.End_date, err = parseDate(*aux.End_date)
        if err != nil {
            return err
        }
    }

    return nil
}

func parseDate(dateStr string) (*time.Time, error) {
    layout := "2006-01-02"
    parsedDate, err := time.Parse(layout, dateStr)
    if err != nil {
        return nil, err
    }
    return &parsedDate, nil
}