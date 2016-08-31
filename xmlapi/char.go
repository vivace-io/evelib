package xmlapi

import (
	"net/url"
	"strconv"
)

// CharAccountBalance returns CharAccountResponse which contains information for each
// account that character has access to.
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

// CharAccountResponse contains general API information in addition to
// a slice of CharAccounts.
type CharAccountResponse struct {
	Result
	Accounts []CharAccount `xml:"result>rowset>row"`
}

// CharAccount contains details on an account belonging to a character.
type CharAccount struct {
	AccountID  int     `xml:"accountID,attr"`
	AccountKey int     `xml:"accountKey,attr"`
	Balance    float64 `xml:"balance,attr"`
}

// CharAssetList returns CharAssetListResponse which contains a list of all assets
// belong to the player. If flat is set to true, assets are not nested.
func (this *Client) CharAssetList(key Key, charID int, flat bool) (*CharAssetListResponse, error) {
	// TODO
	return nil, nil
}

type CharAssetListResponse struct {
	// TODO
}

// CharBlueprints returns BlueprintsResponse, containing a list of blueprints
// belonging to the character.
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

// CharBookmarks returns BookmarksResponse, containing a list of BookmarksFolder
// which contains a list of Bookmarks
func (this *Client) CharBookmarks(key Key, charID int) (*BookmarksResponse, error) {
	var err error
	response := BookmarksResponse{}
	args := url.Values{}
	args.Add("characterID", strconv.Itoa(charID))
	err = this.fetch("/char/Bookmarks.xml.aspx", args, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}
