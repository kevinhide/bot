package models

//Category :""
type Category struct {
	UniqueID string `json:"uniqueId" bson:"uniqueId,omitempty"`
	Category string `json:"category" bson:"category,omitempty"`
	IsParent string `json:"isParent" bson:"isParent,omitempty"`
	Parent   string `json:"parent" bson:"parent,omitempty"`
	Type     string `json:"type" bson:"type,omitempty"`
}
