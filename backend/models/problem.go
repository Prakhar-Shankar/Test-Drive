package models

type Problem struct{
	ID string `json:"id"`
	Title string `json:"title"`
	File string `json:"file"`
	Code string `json:"code,omitempty"`
}