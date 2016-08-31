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

func (this *Client) CharCharacterSheet(key Key, charID int) (*CharacterSheetResponse, error) {
	var err error
	response := CharacterSheetResponse{}
	args := url.Values{}
	args.Add("characterID", strconv.Itoa(charID))
	err = this.fetch("/char/CharacterSheet.xml.aspx", args, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}

type CharacterSheetResponse struct {
	Result
	CharacterSheet CharacterSheet `xml:"result"`
}

type CharacterSheet struct {
	CharacterID       int    `xml:"characterID"`
	Name              string `xml:"name"`
	HomeStationID     int    `xml:"homeStationID"`
	DoB               eTime  `xml:"DoB"`
	Race              string `xml:"race"`
	BloodLineID       int    `xml:"bloodLineID"`
	BloodLine         string `xml:"bloodLine"`
	AncestryID        int    `xml:"ancestryID"`
	Ancestry          string `xml:"ancestry"`
	Gender            string `xml:"gender"`
	CorporationName   string `xml:"corporationName"`
	CorporationID     int    `xml:"corporationID"`
	AllianceName      string `xml:"allianceName"`
	AllianceID        int    `xml:"allianceID"`
	FactionName       string `xml:"factionName"`
	FactionID         int    `xml:"factionID"`
	CloneTypeID       int    `xml:"cloneTypeID"`
	CloneName         string `xml:"cloneName"`
	CloneSkillPoints  int    `xml:"cloneSkillPoints"`
	FreeSkillPoints   int    `xml:"freeSkillPoints"`
	FreeRespecs       int8   `xml:"freeRespecs"`
	CloneJumpDate     eTime  `xml:"cloneJumpDate"`
	LastRespecDate    eTime  `xml:"lastRespecDate"`
	LastTimedRespec   eTime  `xml:"lastTimedRespec"`
	RemoteStationDate eTime  `xml:"remoteStationDate"`
}

func (this *Client) CharChatChannels(key Key, charID int) (*ChatChannelsResponse, error) {
	var err error
	response := ChatChannelsResponse{}
	args := url.Values{}
	args.Add("characterID", strconv.Itoa(charID))
	err = this.fetch("/char/ChatChannels.xml.aspx", args, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}

type ChatChannelsResponse struct {
	Result
	Channels []ChatChannel `xml:"result>rowset>row"`
}

type ChatChannel struct {
	ChannelID     int    `xml:"channelID,attr"`
	OwnerID       int    `xml:"ownerID,attr"`
	OwnerName     string `xml:"ownerName,attr"`
	DisplayName   string `xml:"displayName,attr"`
	ComparisonKey string `xml:"comparisonKey",attr`
	HasPassword   bool   `xml:"hasPassword,attr"`
	MOTD          string `xml:"motd,attr"`
	//Groups        []ChatAccessorGroup `xml:"rowset>row"`
}
