package eveapi

import "fmt"

const npcCorporationsCollectionV1Type = "application/vnd.ccp.eve.NPCCorporationsCollection-v1"

type NPCCorporationsCollectionV1 struct {
	*EVEAPIClient
	crestPagedFrame

	Items []struct {
		itemReference
		Description  string
		Headquarters itemReference
		LoyaltyStore struct {
			Href string
		}
		Ticker string
	}
}

func (c *EVEAPIClient) NPCCorporationsV1(page int64) (*NPCCorporationsCollectionV1, error) {
	ret := &NPCCorporationsCollectionV1{EVEAPIClient: c}
	url := c.base.CREST + fmt.Sprintf("corporations/npccorps/?page=%d", page)

	res, err := c.doJSON("GET", url, nil, ret, npcCorporationsCollectionV1Type, nil)
	if err != nil {
		return nil, err
	}

	ret.getFrameInfo(url, res)
	return ret, nil
}

func (c *NPCCorporationsCollectionV1) NextPage() (*NPCCorporationsCollectionV1, error) {
	w := &NPCCorporationsCollectionV1{EVEAPIClient: c.EVEAPIClient}
	if c.Next.HRef == "" {
		return nil, nil
	}
	res, err := c.doJSON("GET", c.Next.HRef, nil, w, npcCorporationsCollectionV1Type, nil)
	if err != nil {
		return nil, err
	}
	w.getFrameInfo(c.Next.HRef, res)
	return w, nil
}

func (c *NPCCorporationsCollectionV1) PreviousPage() (*NPCCorporationsCollectionV1, error) {
	w := &NPCCorporationsCollectionV1{EVEAPIClient: c.EVEAPIClient}
	if c.Previous.HRef == "" {
		return nil, nil
	}
	res, err := c.doJSON("GET", c.Previous.HRef, nil, w, npcCorporationsCollectionV1Type, nil)
	if err != nil {
		return nil, err
	}
	w.getFrameInfo(c.Next.HRef, res)
	return w, nil
}
