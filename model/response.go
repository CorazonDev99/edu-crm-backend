package model

type DataList struct {
	List       interface{} `json:"data"`
	Pagination Pagination  `json:"pagination"`
}
