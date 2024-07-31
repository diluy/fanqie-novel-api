package routers

import "fanqie-novel-api/api"

type RouterGroup struct {
	BookRouter
	ChapterRouter
	SystemRouter
}

var (
	bookApi    = api.ApiGroupApp.BookAPI
	chapterApi = api.ApiGroupApp.ChapterAPI
)

var RouterGroupApp = new(RouterGroup)
