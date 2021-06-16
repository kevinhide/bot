package models

type discussby struct {
	BillRelatedIssues     string `json:"billRelatedIssues" bson:"billRelatedIssues,omitempty"`
	ServiceRequestRealted string `json:"serviceRequestRealted" bson:"serviceRequestRealted,omitempty"`
	AccountRelatedIssues  string `json:"accountRelatedIssues" bson:"accountRelatedIssues,omitempty"`
}
