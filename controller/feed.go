package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
	"time"
)

type FeedResponse struct {
	Response
	VideoList      []Video        `json:"video_list,omitempty"`
	ModelVideoList []model.Videos `json:"model_video_list"`
	NextTime       int64          `json:"next_time,omitempty"`
}

// Feed same demo video list for every request
func Feed(c *gin.Context) {

	token := c.Query("token")
	if token == "" {
		videoList, _ := service.GetVideoByNoLoginToken()

		c.JSON(http.StatusOK, FeedResponse{
			Response:       Response{StatusCode: 0},
			ModelVideoList: videoList,
			NextTime:       time.Now().Unix(),
		})
	} else {
		videoList, _ := service.GetVideoByLoginToken(token)
		c.JSON(http.StatusOK, FeedResponse{
			Response:       Response{StatusCode: 0},
			ModelVideoList: videoList,
			NextTime:       time.Now().Unix(),
		})
	}

}
