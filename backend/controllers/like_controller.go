package controllers

import (
	"fxboard/global"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
)

func LikeArticle(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := "article:" + articleID + ":likes"

	likes, err := global.RedisDB.Incr(likeKey).Result()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully liked the article", "likes": likes})
}

func GetArticleLikes(ctx *gin.Context) {
	articleID := ctx.Param("id")

	likeKey := "article:" + articleID + ":likes"

	likesText, err := global.RedisDB.Get(likeKey).Result()

	if err == redis.Nil {
		ctx.JSON(http.StatusOK, gin.H{"likes": 0})
		return
	} else if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	likes, err := strconv.Atoi(likesText)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "invalid likes value"})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"likes": likes})
}
