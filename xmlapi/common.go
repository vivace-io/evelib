package xmlapi

type BlueprintsResponse struct {
	Result
	Blueprints []Blueprint `xml:"result>rowset>row"`
}

type Blueprint struct {
	ItemID             int    `xml:"itemID,attr"`
	LocationID         int    `xml:"locationID,attr"`
	TypeID             int    `xml:"typeID,attr"`
	TypeName           string `xml:"typeName,attr"`
	Quantity           int    `xml:"quantity,attr"`
	FlagID             int    `xml:"flagID,attr"`
	TimeEfficiency     int8   `xml:"timeEfficiency,attr"`
	MaterialEfficiency int8   `xml:"materialEfficiency,attr"`
	Runs               int    `xml:"runs,attr"`
}
