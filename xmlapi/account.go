package xmlapi

func (this *Client) AccountAPIKeyInfo(key Key) (*APIKeyInfoResponse, error) {
	var err error
	response := APIKeyInfoResponse{}
	err = this.fetch("/account/APIKeyInfo.xml.aspx", nil, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}

type APIKeyInfoResponse struct {
	Result
	KeyInfo APIKeyInfo `xml:"result>key"`
}

type APIKeyInfo struct {
	AccessMask int             `xml:"accessMask,attr"`
	Type       string          `xml:"type,attr"`
	Rows       []APIKeyInfoRow `xml:"rowset>row"`
}

type APIKeyInfoRow struct {
	ID              int    `xml:"characterID,attr"`
	Name            string `xml:"characterName,attr"`
	CorporationID   int    `xml:"corporationID,attr"`
	CorporationName string `xml:"corporationName,attr"`
	AllianceID      int    `xml:"allianceID,attr"`
	AllianceName    string `xml:"allianceName,attr"`
	FactionID       int    `xml:"factionID,attr"`
	FactionName     string `xml:"factionName,attr"`
}

func (this *Client) AccountStatus(key Key) (*AccountStatus, error) {
	var err error
	response := AccountStatus{}
	err = this.fetch("/account/AccountStatus.xml.aspx", nil, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}

type AccountStatus struct {
	Result
	PaidUntil    eTime `xml:"result>paidUntil"`
	CreateDate   eTime `xml:"result>createData"`
	LogonCount   int   `xml:"result>logonCount"`
	LogonMinutes int   `xml:"result>logonMinutes"`
}

func (this *Client) AccountCharacters(key Key) (*AccountCharactersResponse, error) {
	var err error
	response := AccountCharactersResponse{}
	err = this.fetch("/account/Characters.xml.aspx", nil, key, &response)
	if err != nil {
		return nil, err
	}
	return &response, response.Error
}

type AccountCharactersResponse struct {
	Result
	AccountCharacters AccountCharacters `xml:"result"`
}

type AccountCharacters struct {
	Rows []AccountCharactersRow `xml:"rowset>row"`
}

type AccountCharactersRow struct {
	Name            string `xml:"name,attr"`
	ID              int    `xml:"id,attr"`
	CorporationName string `xml:"corporationName,attr"`
	CorporationID   int    `xml:"corporationID,attr"`
	AllianceName    string `xml:"allianceName,attr"`
	AllianceID      int    `xml:"allianceID,attr"`
	FactionName     string `xml:"factionName,attr"`
	FactionID       int    `xml:"factionID,attr"`
}
