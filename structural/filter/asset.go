package filter

type AssetDto struct {
	TypeID string `json:"typeId"`
	//TenantID    string           `json:"tenantId"`
	Name        string `json:"name"`
	ExternalID  string `json:"externalId"`
	ParentID    string `json:"parentId"`
	Description string `json:"description"`

	Variables []AssetVariableDto `json:"variables"`
}

type AssetVariableDto struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}
