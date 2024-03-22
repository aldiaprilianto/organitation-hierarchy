package response

type OrgHierarchyResponse struct {
	OrgID     string                 `json:"org_id"`
	OrgName   string                 `json:"org_name"`
	OrgChilds []OrgHierarchyResponse `json:"org_childs"`
}
