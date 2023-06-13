package main

type book struct {
	ISIN   string  `json:"isin"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
	Price  float32 `json:"price"`
}
