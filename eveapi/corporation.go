package eveapi

import (
	"fmt"

	"golang.org/x/oauth2"
)

// CorporationSheetXML contains parsed data from XML API
type CorporationSheetXML struct {
	xmlAPIFrame
	CorporationID   int64  `xml:"result>corporationID"`
	CorporationName string `xml:"result>corporationName"`
	Ticker          string `xml:"result>ticker"`
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

// CorporationPublicSheetXML queries the XML API for a given corporationID.
func (c *EVEAPIClient) CorporationPublicSheetXML(corporationID int64) (*CorporationSheetXML, error) {
	w := &CorporationSheetXML{}

	url := c.base.XML + fmt.Sprintf("corp/CorporationSheet.xml.aspx?corporationID=%d", corporationID)
	_, err := c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type CorporationSheetDetailXML struct {
	CorporationSheetXML
	Divisions []struct {
		Name     string `xml:"name,attr"`
		Accounts []struct {
			Key         string `xml:"accountKey,attr"`
			Description string `xml:"description,attr"`
		} `xml:"row"`
	} `xml:"result>rowset"`
}

// CorporationSheetXML queries the XML API for details of the token's corporation.
func (c *EVEAPIClient) CorporationSheetXML(auth oauth2.TokenSource) (*CorporationSheetDetailXML, error) {
	w := &CorporationSheetDetailXML{}
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("corp/CorporationSheet.xml.aspx?accessToken=%s&accessType=corporation", tok.AccessToken)
	_, err = c.doXML("GET", url, nil, w)
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

type CorporationAssetsXML struct {
	xmlAPIFrame
	Entries []struct {
		ItemID      int64 `xml:"itemID,attr"`
		LocationID  int64 `xml:"locationID,attr"`
		TypeID      int64 `xml:"typeID,attr"`
		Quantity    int64 `xml:"quantity,attr"`
		FlagID      int64 `xml:"flagID,attr"`
		Singleton   bool  `xml:"singleton,attr"`
		RawQuantity int64 `xml:"rawQuantity,attr"`
	} `xml:"result>rowset>row"`
}

// CorporationAssetsXML queries the XML API for assets owned by corporationID.
func (c *EVEAPIClient) CorporationAssetsXML(auth oauth2.TokenSource, corporationID int64) (*CorporationAssetsXML, error) {
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("corp/AssetList.xml.aspx?corporationID=%d&accessToken=%s&flat=1", corporationID, tok.AccessToken)
	w := &CorporationAssetsXML{}
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

type CorporationMarketOrdersXML struct {
	xmlAPIFrame
	Entries []struct {
		OrderID      int64      `xml:"orderID,attr"`
		CharID       int64      `xml:"charID,attr"`
		StationID    int64      `xml:"stationID,attr"`
		VolEntered   int32      `xml:"volEntered,attr"`
		VolRemaining int32      `xml:"volRemaining,attr"`
		MinVolume    int32      `xml:"minVolume,attr"`
		OrderState   int32      `xml:"orderState,attr"`
		TypeID       int64      `xml:"typeID,attr"`
		Range        int32      `xml:"range,attr"`
		AccountKey   int32      `xml:"accountKey,attr"`
		Duration     int32      `xml:"duration,attr"`
		Escrow       float64    `xml:"escrow,attr"`
		Price        float64    `xml:"price,attr"`
		Bid          bool       `xml:"bid,attr"`
		Issued       EVEXMLTime `xml:"issued,attr"`
	} `xml:"result>rowset>row"`
}

// CorporationMarketOrdersXML queries the XML API for orders placed by characters in corporationID.
func (c *EVEAPIClient) CorporationMarketOrdersXML(auth oauth2.TokenSource, corporationID int64) (*CorporationMarketOrdersXML, error) {
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("corp/MarketOrders.xml.aspx?corporationID=%d&accessToken=%s", corporationID, tok.AccessToken)
	w := &CorporationMarketOrdersXML{}
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}

// CorporationMarketOrdersXML queries the XML API for a specific order.
func (c *EVEAPIClient) CorporationMarketOrderXML(auth oauth2.TokenSource, corporationID, orderID int64) (*CorporationMarketOrdersXML, error) {
	tok, err := auth.Token()
	if err != nil {
		return nil, err
	}
	url := c.base.XML + fmt.Sprintf("corp/MarketOrders.xml.aspx?corporationID=%d&orderID=%d&accessToken=%s", corporationID, orderID, tok.AccessToken)
	w := &CorporationMarketOrdersXML{}
	_, err = c.doXML("GET", url, nil, w)
	if err != nil {
		return nil, err
	}
	return w, nil
}
