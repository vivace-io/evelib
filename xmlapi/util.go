package xmlapi

import (
	"encoding/xml"
	"fmt"
	"time"
)

var (
	dateFormat = "2006-01-02 15:04:05"
)

type APIError struct {
	Code    int    `xml:"code,attr"`
	Message string `xml:",chardata"`
}

func (this APIError) Error() string {
	return fmt.Sprintf("%v (code:%v)", this.Message, this.Code)
}

type eTime struct {
	time.Time
}

func (this eTime) String() string {
	return this.Format(dateFormat)
}

func (this *eTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var v string
	d.DecodeElement(&v, &start)
	parse, err := time.Parse(dateFormat, v)
	if err != nil {
		return nil
	}
	*this = eTime{parse}
	return nil
}
