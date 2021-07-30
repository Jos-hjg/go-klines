package models

type CheckModel struct{
	Title string `form:"title"`
	Code string `form:"code"`
}

type DataModel struct{
	Name string
	Message []map[string]interface{}
}



