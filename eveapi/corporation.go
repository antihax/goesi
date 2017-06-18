package eveapi

import (
	"fmt"

	"golang.org/x/oauth2"
)

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
	_, err := c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type CorporationIndustryJobsXML struct {
	xmlAPIFrame
	Entries []struct {
		JobID                int64      `xml:"jobID,attr"`
		InstallerID          int64      `xml:"installerID,attr"`
		InstallerName        string     `xml:"installerName,attr"`
		FacilityID           int64      `xml:"facilityID,attr"`
		SolarSystemName      string     `xml:"solarSystemName,attr"`
		SolarSystemID        int64      `xml:"solarSystemID,attr"`
		StationID            int64      `xml:"stationID,attr"`
		ActivityID           int64      `xml:"activityID,attr"`
		BlueprintID          int64      `xml:"blueprintID,attr"`
		BlueprintTypeID      int64      `xml:"blueprintTypeID,attr"`
		BlueprintTypeName    string     `xml:"blueprintTypeName,attr"`
		BlueprintLocationID  int64      `xml:"blueprintLocationID,attr"`
		OutputLocationID     int64      `xml:"outputLocationID,attr"`
		ProductTypeID        int64      `xml:"productTypeID,attr"`
		Runs                 int64      `xml:"runs,attr"`
		Cost                 float64    `xml:"cost,attr"`
		LicensedRuns         int64      `xml:"licensedRuns,attr"`
		Probability          float64    `xml:"probability,attr"`
		ProductTypeName      string     `xml:"productTypeName,attr"`
		Status               int64      `xml:"status,attr"`
		TimeInSeconds        int64      `xml:"timeInSeconds,attr"`
		StartDate            EVEXMLTime `xml:"startDate,attr"`
		EndDate              EVEXMLTime `xml:"endDate,attr"`
		PauseDate            EVEXMLTime `xml:"pauseDate,attr"`
		CompletedDate        EVEXMLTime `xml:"completedDate,attr"`
		CompletedCharacterID int64      `xml:"completedCharacterID,attr"`
		SuccessfulRuns       int64      `xml:"successfulRuns,attr"`
	} `xml:"result>rowset>row"`
}

// CorporationIndustryJobsXML queries the XML API for active industry jobs for corporationID.
func (c *EVEAPIClient) CorporationIndustryJobsXML(auth oauth2.TokenSource, corporationID int64) (*CorporationIndustryJobsXML, error) {
	w := &CorporationIndustryJobsXML{}
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("corp/IndustryJobs.xml.aspx?corporationID=%d&accessToken=%s", corporationID, tok.AccessToken)
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// CorporationIndustryJobsHistoryXML queries the XML API for finished industry jobs for corporationID.
func (c *EVEAPIClient) CorporationIndustryJobsHistoryXML(auth oauth2.TokenSource, corporationID int64) (*CorporationIndustryJobsXML, error) {
	w := &CorporationIndustryJobsXML{}
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("corp/IndustryJobsHistory.xml.aspx?corporationID=%d&accessToken=%s", corporationID, tok.AccessToken)
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type CorporationBlueprintsXML struct {
	xmlAPIFrame
	Entries []struct {
		ItemID             int64  `xml:"itemID,attr"`
		LocationID         int64  `xml:"locationID,attr"`
		TypeID             int64  `xml:"typeID,attr"`
		TypeName           string `xml:"typeName,attr"`
		Quantity           int64  `xml:"quantity,attr"`
		FlagID             int64  `xml:"flagID,attr"`
		TimeEfficiency     int64  `xml:"timeEfficiency,attr"`
		MaterialEfficiency int64  `xml:"materialEfficiency,attr"`
		Runs               int64  `xml:"runs,attr"`
	} `xml:"result>rowset>row"`
}

// CorporationBlueprintsXML queries the XML API for blueprints owned by corporationID.
func (c *EVEAPIClient) CorporationBlueprintsXML(auth oauth2.TokenSource, corporationID int64) (*CorporationBlueprintsXML, error) {
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("corp/Blueprints.xml.aspx?corporationID=%d&accessToken=%s", corporationID, tok.AccessToken)
	w := &CorporationBlueprintsXML{}
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
