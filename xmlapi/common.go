package xmlapi

// BlueprintsResponse contains a slice of Blueprints
type BlueprintsResponse struct {
	Result
	Blueprints []Blueprint `xml:"result>rowset>row"`
}

// Blueprint contains information on a single blueprint, copy or original.
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
