package routers

import "github.com/gin-gonic/gin"

type ChapterRouter struct{}

func (s *BookRouter) InitChapterRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	chapterRouter := Router.Group("chapter")
	{
		chapterRouter.GET("/content", chapterApi.GetContentByBookId)
	}
	return chapterRouter
}
