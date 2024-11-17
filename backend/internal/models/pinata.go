package models

type Pinata struct {
	Cid *string `json:"cid"`
}

type PinataResponse struct {
	Url string `json:"url"`
}