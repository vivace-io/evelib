package xmlapi

func (this *Client) AccountAPIKeyInfo(key Key) (*APIKeyInfoResponse, error) {
	var err error
	response := APIKeyInfoResponse{}
	err = this.fetch("/account/APIKeyInfo.xml.aspx", key, &response)
	if err != nil {
		return nil, err
	}
	return &response, err
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
