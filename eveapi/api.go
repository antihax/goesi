package eveapi

import (
	"fmt"

	"golang.org/x/oauth2"
)

// CCP basic XML Frame
type xmlAPIFrame struct {
	Version     int        `xml:"eveapi>version"`
	CurrentTime EVEXMLTime `xml:"currentTime"`
	CachedUntil EVEXMLTime `xml:"cachedUntil"`
}

type EVEAPICallListXML struct {
	xmlAPIFrame
	Entries []struct {
		Kind string `xml:"name,attr"`
		Rows []struct {
			Name        string `xml:"name,attr"`
			Description string `xml:"description,attr"`
			GroupID     int32  `xml:"groupID,attr"`
			AccessMask  int64  `xml:"accessMask,attr"`
			Type        string `xml:"type,attr"`
		} `xml:"row"`
	} `xml:"result>rowset"`
}

func (c *EVEAPIClient) CallListXML(auth oauth2.TokenSource) (*EVEAPICallListXML, error) {
	w := &EVEAPICallListXML{}
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("api/CallList.xml.aspx?accessToken=%s", tok.AccessToken)
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type EVEAPIKeyInfoXML struct {
	xmlAPIFrame
	Key struct {
		AccessMask int64      `xml:"accessMask,attr"`
		Type       string     `xml:"type,attr"`
		Expires    EVEXMLTime `xml:"expires,attr"`
		Info       struct {
			CharacterID     int64  `xml:"characterID,attr"`
			CharacterName   string `xml:"characterName,attr"`
			CorporationID   int64  `xml:"corporationID,attr"`
			CorporationName string `xml:"corporationName,attr"`
			AllianceID      int64  `xml:"allianceID,attr"`
			AllianceName    string `xml:"allianceName,attr"`
			FactionID       int64  `xml:"factionID,attr"`
			FactionName     string `xml:"factionName,attr"`
		} `xml:"rowset>row"`
	} `xml:"result>key"`
}

// APIKeyInfoXML queries the EVE API for information about the given API token.
//
// The accessType parameter should be one of: character, corporation.
func (c *EVEAPIClient) APIKeyInfoXML(auth oauth2.TokenSource, accessType string) (*EVEAPIKeyInfoXML, error) {
	w := &EVEAPIKeyInfoXML{}
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("account/APIKeyInfo.xml.aspx?accessToken=%s&accessType=%s", tok.AccessToken, accessType)
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
