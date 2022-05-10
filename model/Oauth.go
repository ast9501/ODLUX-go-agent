package model

type Provider struct {
	Id		string	`json: "id"`
	Title	string	`json: "title"`
	Url		string	`json: "loginUrl"`
}