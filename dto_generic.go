package cvr

import (
	"encoding/json"
	"fmt"
	"time"
)

type SidstOpdateret struct {
	SidstOpdateret *CVRDato `json:"sidstOpdateret"`
}

type SidstIndlaest struct {
	SidstIndlaest CVRDato `json:"sidstIndlaest"`
}

type CVRDato time.Time

const (
	CVRDateFormat1 = time.RFC3339
	CVRDateFormat2 = "2006-01-02T15:04:05+00:04:05"
)

func (d *CVRDato) UnmarshalJSON(data []byte) error {
	dateFormats := []string{CVRDateFormat1, CVRDateFormat2}
	date := string(data[1 : len(data)-1])
	for _, dateFormat := range dateFormats {
		t, err := time.Parse(dateFormat, date)
		if err != nil {
			continue
		}

		*d = CVRDato(t)
		return nil
	}

	fmt.Printf("ERROR: Failed to parse date %s using formats %v\n", date, dateFormats)
	return nil
}

type Periode struct {
	GyldigFra *GyldigDato `json:"gyldigFra"`
	GyldigTil *GyldigDato `json:"gyldigTil"`
}

func (me Periode) Less(you Periode) bool {
	meOngoing := me.GyldigTil == nil
	youOngoing := you.GyldigTil == nil

	// Both are ongoing
	if meOngoing && youOngoing {
		return true
	}

	// only youOngoing means me < you
	if youOngoing {
		return true
	}

	// only meOngoing means me > you
	if meOngoing {
		return false
	}

	// Both have end dates, check if me <= you
	return time.Time(*me.GyldigTil).Before(time.Time(*you.GyldigTil))
}

const gyldigDatoFormat = "2006-01-02"

type GyldigDato time.Time

func (d *GyldigDato) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(gyldigDatoFormat, string(data[1:len(data)-1]))
	if err != nil {
		return err
	}

	*d = GyldigDato(t)
	return nil
}

func (d *GyldigDato) MarshalJSON() ([]byte, error) {
	if d == nil {
		return nil, nil
	}

	t := time.Time(*d).Format(gyldigDatoFormat)
	res, err := json.Marshal(t)

	return res, err
}
