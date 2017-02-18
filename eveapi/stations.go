package eveapi

import "fmt"

// CharacterInfo returned data from XML API
type ConquerableStationsXML struct {
	xmlAPIFrame
	Stations []struct {
		StationID     int64  `xml:"stationID,attr"`
		StationName   string `xml:"stationName,attr"`
		StationTypeID int64  `xml:"stationTypeID,attr"`
		SolarSystemID int64  `xml:"solarSystemID,attr"`
		CorporationID int64  `xml:"corporationID,attr"`
	} `xml:"result>rowset>row"`
}

// GetCharacterInfo queries the XML API for a given characterID.
func (c *EVEAPIClient) ConquerableStationsListXML() (*ConquerableStationsXML, error) {
	w := &ConquerableStationsXML{}

	url := c.base.XML + fmt.Sprintf("/eve/ConquerableStationList.xml.aspx")
	_, err := c.doXML("GET", url, nil, w, nil, nil)
	if err != nil {
		return nil, err
	}
	return w, nil
}
