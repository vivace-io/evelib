package xmlapi

import (
	"net/url"
	"strconv"
)

func (this *Client) CharAccountBalance(key Key, charID int) (*CharAccountResponse, error) {
	var err error
	response := CharAccountResponse{}
	args := url.Values{}
	args.Add("characterID", strconv.Itoa(charID))
	err = this.fetch("/char/AccountBalance.xml.aspx", args, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}

type CharAccountResponse struct {
	Result
	Accounts []CharAccount `xml:"result>rowset>row"`
}

type CharAccount struct {
	AccountID  int     `xml:"accountID,attr"`
	AccountKey int     `xml:"accountKey,attr"`
	Balance    float64 `xml:"balance,attr"`
}

func (this *Client) CharAssetList(key Key, charID int, flat bool) (*CharAssetListResponse, error) {
	// TODO
	return nil, nil
}

type CharAssetListResponse struct {
	// TODO
}

func (this *Client) CharBlueprints(key Key, charID int) (*BlueprintsResponse, error) {
	var err error
	response := BlueprintsResponse{}
	args := url.Values{}
	args.Add("characterID", strconv.Itoa(charID))
	err = this.fetch("/char/Blueprints.xml.aspx", args, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}
