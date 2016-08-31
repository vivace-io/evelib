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

// BookmarksResponse contains any number of bookmark folders.
type BookmarksResponse struct {
	Result
	Folders []BookmarksFolder `xml:"result>rowset>row"`
}

// BookmarksFolder contains any number of bookmarks.
type BookmarksFolder struct {
	FolderID   int        `xml:"folderID"`
	FolderName string     `xml:"folderName"`
	CreatorID  int        `xml:"creatorID"`
	Bookmarks  []Bookmark `xml:"rowset>row"`
}

// Bookmark is a bookmark in Eve
type Bookmark struct {
	BookmarkID int `xml:"bookmarkID,attr"`
	CreatorID  int `xml:"creatorID,attr"`
	// Created    eTime   `xml:"created,attr"`
	ItemID     int     `xml:"itemID,attr"`
	TypeID     int     `xml:"typeID,attr"`
	LocationID int     `xml:"locationID,attr"`
	X          float64 `xml:"x,attr"`
	Y          float64 `xml:"y,attr"`
	Z          float64 `xml:"z,attr"`
	Memo       string  `xml:"memo,attr"`
	Not        string  `xml:"note,attr"`
}
