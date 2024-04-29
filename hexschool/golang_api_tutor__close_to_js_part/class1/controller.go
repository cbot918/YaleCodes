package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	R Repository
}

func NewController(r Repository) Controller {
	return Controller{
		R: r,
	}
}

func (c *Controller) CreatePostController(ctx *gin.Context) {

	var err error

	post := Post{
		Title: "test",
		Body:  "test body",
	}

	post, err = c.R.CreatePost(post)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, post)
}

func (c *Controller) ReadPostsController(ctx *gin.Context) {
	posts, err := c.R.ReadPosts()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, posts)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
