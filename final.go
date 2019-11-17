package main

type User struct {
	Id       int    `json:"id" form:"id"`
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}
type Bookmark struct {
	Id       int    `json:"id" form:"id"`
	IdRecipe int    `json:"idrecipe" form:"idrecipe"`
	Type     string `json:"type" form:"type"`
}
type dataRecipe struct {
	Id   int    `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
	Url  string `json:"url" form:"url"`
}

type listFilter struct {
	Filter string `json:"filter" form:"filter"`
	Url    string `json:"url" form:"url"`
}

func main()  {
	
}