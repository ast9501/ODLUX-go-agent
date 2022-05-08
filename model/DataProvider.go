package model

type SortProBody struct {
	Property	string `json:"property"`
	Sort		string  `json:"sortorder"`
}

type Pagination struct {
	Size int `json: "size"`
	Page int `json: "page"`
}

type NodeInfo struct {
	Id		string `json:"id"`
	NodeId	string `json:"node-id"`
	Host	string `json:"host"`
	Port	string `json: "port"`
	Username	string `json: "username"`
	Passwd		string `json: "password"`
    Required	string `json: "is-required"`
}

type DataProvider struct {
	Data	NodeInfo `json: "data-provider:input"`
}