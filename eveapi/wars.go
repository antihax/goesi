package eveapi

import "fmt"

const warsCollectionV1Type = "application/vnd.ccp.eve.WarsCollection-v1"

type WarsCollectionV1 struct {
	*EVEAPIClient
	crestPagedFrame

	Items []struct {
		HRef string
		ID   int
	}
}

const warKillmailsV1Type = "application/vnd.ccp.eve.WarKillmails-v1"

type WarKillmailsV1 struct {
	*EVEAPIClient
	crestPagedFrame

	Items []struct {
		HRef string
		ID   int
	}
}

const warV1Type = "application/vnd.ccp.eve.War-v1"

type WarV1 struct {
	*EVEAPIClient
	crestSimpleFrame

	TimeFinished  EVETime
	OpenForAllies bool
	TimeStarted   EVETime
	AllyCount     int
	TimeDeclared  EVETime

	Allies []struct {
		HRef string
		ID   int64
		Icon struct {
			HRef string
		}
		Name string
	}
	Aggressor struct {
		ShipsKilled int

		Name string
		HRef string

		Icon struct {
			HRef string
		}
		ID        int64
		IskKilled float64
	}
	Mutual bool

	Killmails string

	Defender struct {
		ShipsKilled int

		Name string
		HRef string

		Icon struct {
			HRef string
		}
		ID        int64
		IskKilled float64
	}
	ID int64
}

func (c *EVEAPIClient) WarsV1(page int) (*WarsCollectionV1, error) {
	w := &WarsCollectionV1{EVEAPIClient: c}
	url := c.base.CREST + fmt.Sprintf("wars/?page=%d", page)
	res, err := c.doJSON("GET", url, nil, w, warsCollectionV1Type, nil)
	if err != nil {
		return nil, err
	}
	w.getFrameInfo(url, res)
	return w, nil
}

func (c *WarsCollectionV1) NextPage() (*WarsCollectionV1, error) {
	w := &WarsCollectionV1{EVEAPIClient: c.EVEAPIClient}
	if c.Next.HRef == "" {
		return nil, nil
	}
	res, err := c.doJSON("GET", c.Next.HRef, nil, w, warsCollectionV1Type, nil)
	if err != nil {
		return nil, err
	}
	w.getFrameInfo(c.Next.HRef, res)
	return w, nil
}

func (c *EVEAPIClient) WarV1(href string) (*WarV1, error) {
	w := &WarV1{EVEAPIClient: c}
	res, err := c.doJSON("GET", href, nil, w, warV1Type, nil)
	if err != nil {
		return nil, err
	}
	w.getFrameInfo(res)
	return w, nil
}

func (c *EVEAPIClient) WarByID(id int) (*WarV1, error) {
	url := c.base.CREST + fmt.Sprintf("wars/%d/", id)
	return c.WarV1(url)
}

// GetKillmails provides a list of killmails associated to this war.
func (c *WarV1) KillmailsV1() (*WarKillmailsV1, error) {
	w := &WarKillmailsV1{EVEAPIClient: c.EVEAPIClient}
	res, err := c.doJSON("GET", c.Killmails, nil, w, warKillmailsV1Type, nil)
	if err != nil {
		return nil, err
	}
	w.getFrameInfo(c.Killmails, res)
	return w, nil
}
