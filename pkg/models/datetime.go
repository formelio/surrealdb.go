package models

import (
	"fmt"
	"time"

	"github.com/fxamacker/cbor/v2"
	"github.com/surrealdb/surrealdb.go/pkg/constants"
)

// CustomDateTime embeds time.Time
type CustomDateTime struct {
	time.Time
}

func (d *CustomDateTime) MarshalCBOR() ([]byte, error) {
	enc := getCborEncoder()

	totalNS := d.Nanosecond()

	s := totalNS / constants.OneSecondToNanoSecond
	ns := totalNS % constants.OneSecondToNanoSecond

	return enc.Marshal(cbor.Tag{
		Number:  TagCustomDatetime,
		Content: [2]int64{int64(s), int64(ns)},
	})
}

func (d *CustomDateTime) UnmarshalCBOR(data []byte) error {
	dec := getCborDecoder()

	var temp [2]int64
	err := dec.Unmarshal(data, &temp)
	if err != nil {
		return err
	}

	s := temp[0]
	ns := temp[1]

	*d = CustomDateTime{time.Unix(s, ns)}

	return nil
}

func (d *CustomDateTime) String() string {
	layout := "2006-01-02T15:04:05Z"
	return d.Format(layout)
}

func (d *CustomDateTime) SurrealString() string {
	return fmt.Sprintf("<datetime> '%s'", d.String())
}
