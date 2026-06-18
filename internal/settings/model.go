package settings

type OrgSetting struct {
	ID    string `json:"id"`
	OrgID string `json:"org_id"`
	Key   string `json:"key"`
	Value string `json:"value"`
}
