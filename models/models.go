package models

type Product struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
	Category int `json:"category,omitempty"`
	Quantity int `json:"quantity,omitempty"`
	Price float64 `json:"price,omitempty"`
	Image string `json:"image,omitempty"`
}

type Category struct {
	Id int `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}

// https://api.flickr.com/services
// Key:
// 5b6cb2eeb1a833e847cfbcc972c6c27c

// Secret:
// 5aa11c3c7c76bb8b