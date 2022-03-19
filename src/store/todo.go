package store

type Todo struct {
	Id     int    `json:"id"`
	Title  string `json:"title" binding:"required"`
	Status bool   `json:"status" binding:"required"`
}

var Id = 0

var TodoManagement = []Todo{
	{
		Id:     100,
		Title:  "kakkaka",
		Status: true,
	},
	{
		Id:     200,
		Title:  "kakkaka",
		Status: true,
	},
}
