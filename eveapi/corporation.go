package eveapi

import "fmt"

// CharacterInfo returned data from XML API
type CorporationSheetXML struct {
	xmlAPIFrame
	CorporationID   int64  `xml:"result>corporationID"`
	CorporationName string `xml:"result>corporationName"`
	Ticker          string `xml:"result>ricker"`
	CEOID           int64  `xml:"result>ceoID"`
	CEOName         string `xml:"result>ceoName"`
	StationID       int64  `xml:"result>stationID"`
	StationName     string `xml:"result>stationName"`
	Description     string `xml:"result>description"`
	AllianceID      int64  `xml:"result>allianceID"`
	AllianceName    string `xml:"result>allianceName"`
	FactionID       int64  `xml:"result>factionID"`
	URL             string `xml:"result>url"`
	MemberCount     int64  `xml:"result>memberCount"`
	Shares          int64  `xml:"result>shares"`
	Logo            struct {
		GraphicID int64 `xml:"grapicID,attr"`
		Shape1    int64 `xml:"shape1,attr"`
		Shape2    int64 `xml:"shape2,attr"`
		Shape3    int64 `xml:"shape3,attr"`
		Color1    int64 `xml:"color1,attr"`
		Color2    int64 `xml:"color2,attr"`
		Color3    int64 `xml:"color3,attr"`
	} `xml:"result>logo"`
}

// GetCharacterInfo queries the XML API for a given characterID.
func (c *EVEAPIClient) CorporationPublicSheetXML(corporationID int64) (*CorporationSheetXML, error) {
	w := &CorporationSheetXML{}

	url := c.base.XML + fmt.Sprintf("corp/CorporationSheet.xml.aspx?corporationID=%d", corporationID)
	_, err := c.doXML("GET", url, nil, w, nil, nil)
	if err != nil {
		return nil, err
	}
	return w, nil
}
