package entity

type Organization struct {
	ID          int64  `json:"id"`
	OrgID       string `json:"org_id"`
	OrgName     string `json:"org_name"`
	OrgStatus   string `json:"org_status"`
	OrgParentID string `json:"org_parent_id"`
}
