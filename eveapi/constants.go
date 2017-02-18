package eveapi

import (
	"encoding/xml"
	"strings"
	"time"
)

const (
	BASE_API_VERSION = "application/vnd.ccp.eve.Api-v5"
	BASE_MEDIA_TYPE  = "application/json"
	USER_AGENT       = "https://github.com/antihax/eveapi"
)

type EveURI struct {
	AppManagement string
	CREST         string
	Images        string
	Login         string
	XML           string
}

var eveTQ = EveURI{
	AppManagement: "https://developers.eveonline.com/",
	CREST:         "https://crest-tq.eveonline.com/",
	Images:        "https://image.eveonline.com/",
	Login:         "https://login.eveonline.com/",
	XML:           "https://api.eveonline.com/",
}

var eveSisi = EveURI{
	AppManagement: "https://developers.testeveonline.com/",
	CREST:         "https://api-sisi.testeveonline.com/",
	Images:        "https://image.testeveonline.com/",
	Login:         "https://sisilogin.testeveonline.com/",
	XML:           "https://api.testeveonline.com/",
}

type EVETime struct {
	time.Time
}

// Cannot properly Unmarshal CCP's time stamps?
const eveTimeLayout = "2006-01-02T15:04:05"

func (c *EVETime) UnmarshalJSON(b []byte) (err error) {
	t := string(b)
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveTimeLayout, t)
	return
}
func (c *EVETime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t string
	d.DecodeElement(&t, &start)
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveTimeLayout, t)
	return
}
func (c *EVETime) UnmarshalXMLAttr(a xml.Attr) (err error) {
	t := a.Value
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveTimeLayout, t)
	return
}

type EVEXMLTime struct {
	time.Time
}

// Cannot properly Unmarshal CCP's time stamps?
const eveXMLTimeLayout = "2006-01-02 15:04:05"

func (c *EVEXMLTime) UnmarshalJSON(b []byte) (err error) {
	t := string(b)
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveXMLTimeLayout, t)
	return
}

func (c *EVEXMLTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t string
	d.DecodeElement(&t, &start)
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveXMLTimeLayout, t)
	return
}

func (c *EVEXMLTime) UnmarshalXMLAttr(a xml.Attr) (err error) {
	t := a.Value
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveXMLTimeLayout, t)
	return
}

type EVEKillmailTime struct {
	time.Time
}

// Cannot properly Unmarshal CCP's KillmailTime stamps?
const eveKillmailTimeLayout = "2006.01.02 15:04:05"

func (c *EVEKillmailTime) UnmarshalJSON(b []byte) (err error) {
	t := string(b)
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveKillmailTimeLayout, t)
	return
}

func (c *EVEKillmailTime) UnmarshalXML(d *xml.Decoder, start xml.StartElement) (err error) {
	var t string
	d.DecodeElement(&t, &start)
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveKillmailTimeLayout, t)
	return
}

func (c *EVEKillmailTime) UnmarshalXMLAttr(a xml.Attr) (err error) {
	t := a.Value
	t = strings.Replace(t, `"`, "", -1)
	c.Time, err = time.Parse(eveKillmailTimeLayout, t)
	return
}
