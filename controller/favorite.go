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

// FavoriteList 点赞列表 所有用户都有相同的收藏夹视频列表
//func FavoriteList(c *gin.Context) {
//	c.JSON(http.StatusOK, VideoListResponse{
//		Response: Response{
//			StatusCode: 0,
//		},
//		VideoList: DemoVideos,
//	})
//}

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
		videoList, err := service.GetLikeVideo(int64(UID))
		log.Println(videoList)
		log.Println(err)
		log.Println(UID)
		c.JSON(http.StatusOK, FeedResponse{
			Response:       Response{StatusCode: 0, StatusMsg: "Get success"},
			ModelVideoList: videoList,
		})
	}

}
