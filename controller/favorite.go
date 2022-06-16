package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-demo/service"
	"strconv"
)

// FavoriteAction 点赞操作 没有实际效果，只是检查token是否有效
func FavoriteAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

func FavoriteList(c *gin.Context) {
	userid := c.Query("user_id")
	UID, _ := strconv.ParseUint(userid, 10, 32)
	token := c.Query("token")

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		videoList_, err := service.GetLikeVideo(int64(UID))
		videoList, err := GenerateVideo(videoList_)
		log.Println(videoList)
		log.Println(err)
		log.Println(UID)
		c.JSON(http.StatusOK, FeedResponse{
			Response:  Response{StatusCode: 0, StatusMsg: "Get success"},
			VideoList: videoList,
		})
	}

}
