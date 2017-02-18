package eveapi

// CCP basic XML Frame
type xmlAPIFrame struct {
	Version     int        `xml:"eveapi>version"`
	CurrentTime EVEXMLTime `xml:"currentTime"`
	CachedUntil EVEXMLTime `xml:"cachedUntil"`
}
