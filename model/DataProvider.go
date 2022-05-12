package model

type SortProBody struct {
	Property	string `json:"property"`
	Sort		string  `json:"sortorder"`
}

type Pagination struct {
	Size 	int `json:"size"`
	Page 	string `json:"page"`
	Total	int	`json:"total"`
}

type NodeDtail struct {
	UnavailCap	[]string `json:"unavailable-capabilities"`
	AvailCap	[]string `json:"available-capabilities"`
}

type NodeInfo struct {
	CoreMdlCap	string		`json:"core-model-capability"`
	Type		string		`json:"device-type"`
	Dtail		NodeDtail	`json:"node-details"`
	Id			string 		`json:"id"`
	NodeId		string 		`json:"node-id"`
	Host		string 		`json:"host"`
	Port		string 		`json:"port"`
	Username	string 		`json:"username"`
    Required	bool 		`json:"is-required"`
	Status		string		`json:"status"`
}

type DataProviderInput struct {
	Data	NodeInfo `json:"data-provider:input"`
}

type DataProviderOutput struct {
	Data	[]NodeInfo `json:"data"`
	Page	Pagination `json:"pagination"`
}

type OdlReadConnRes struct {
	Body	DataProviderOutput	`json:"data-provider:output"`
}