package eveapi

import (
	"fmt"

	"golang.org/x/oauth2"
)

const contactCreateV1Type = "application/vnd.ccp.eve.ContactCreate-v1"

type ContactCreateV1 struct {
	Standing    float64 `json:"standing,omitempty"`
	ContactType string  `json:"contactType,omitempty"`
	Contact     struct {
		Href string `json:"href,omitempty"`
		Name string `json:"name,omitempty"`
		ID   int64  `json:"id,omitempty"`
	} `json:"contact,omitempty"`
	Watched bool `json:"watched,omitempty"`
}

func (c *EVEAPIClient) ContactSetV1(auth oauth2.TokenSource, characterID int64, id int64, ref string, standing float64) error {

	contact := ContactCreateV1{Standing: standing}
	contact.Contact.ID = id
	contact.Contact.Href = ref

	url := fmt.Sprintf(c.base.CREST+"characters/%d/contacts/", characterID)
	_, err := c.doJSON("POST", url, contact, nil, contactCreateV1Type, auth)
	if err != nil {
		return err
	}

	return nil
}

func (c *EVEAPIClient) ContactDeleteV1(auth oauth2.TokenSource, characterID int64, id int64, ref string) error {
	contact := ContactCreateV1{}
	contact.Contact.ID = id
	contact.Contact.Href = ref

	url := fmt.Sprintf(c.base.CREST+"characters/%d/contacts/%d/", characterID, id)
	ret := make(map[string]interface{})
	_, err := c.doJSON("DELETE", url, contact, &ret, contactCreateV1Type, auth)
	if err != nil {
		return err
	}

	return nil
}

const contactCollectionV1Type = "application/vnd.ccp.eve.ContactCollection-v1"

type ContactCollectionV1 struct {
	*EVEAPIClient
	auth oauth2.TokenSource
	crestPagedFrame
	Items []struct {
		Standing  float64
		Character characterReference
		Contact   struct {
			Href string
			Name string
			ID   int64
		}
		Href        string
		ContactType string
		Watched     bool
		Blocked     bool
	}
}

func (c *EVEAPIClient) ContactsV1(auth oauth2.TokenSource, characterID int64) (*ContactCollectionV1, error) {

	ret := &ContactCollectionV1{EVEAPIClient: c, auth: auth}

	url := fmt.Sprintf(c.base.CREST+"characters/%d/contacts/", characterID)

	res, err := c.doJSON("GET", url, nil, ret, contactCollectionV1Type, auth)
	if err != nil {
		return nil, err
	}

	ret.getFrameInfo(url, res)
	return ret, nil
}

func (c *ContactCollectionV1) NextPage() (*ContactCollectionV1, error) {
	ret := &ContactCollectionV1{EVEAPIClient: c.EVEAPIClient, auth: c.auth}
	if c.Next.HRef == "" {
		return nil, nil
	}
	res, err := c.doJSON("GET", c.Next.HRef, nil, ret, contactCollectionV1Type, c.auth)
	if err != nil {
		return nil, err
	}
	ret.getFrameInfo(c.Next.HRef, res)
	return ret, nil
}

func (c *ContactCollectionV1) PreviousPage() (*ContactCollectionV1, error) {
	ret := &ContactCollectionV1{EVEAPIClient: c.EVEAPIClient, auth: c.auth}
	if c.Previous.HRef == "" {
		return nil, nil
	}
	res, err := c.doJSON("GET", c.Previous.HRef, nil, ret, contactCollectionV1Type, c.auth)
	if err != nil {
		return nil, err
	}
	ret.getFrameInfo(c.Previous.HRef, res)
	return ret, nil
}
