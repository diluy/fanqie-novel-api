package api

import "fanqie-novel-api/service"

type ApiGroup struct {
	BookAPI
	ChapterAPI
}

var (
	bookService    = service.ServiceGroupApp.BookService
	chapterService = service.ServiceGroupApp.ChapterService
)

var ApiGroupApp = new(ApiGroup)
