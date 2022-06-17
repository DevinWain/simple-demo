package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"simple-demo/model"
	"simple-demo/service"
	"strconv"
)

type UserListResponse struct {
	Response
	UserList []User `json:"user_list"`
}

// RelationAction no practical effect, just check if token is valid
func RelationAction(c *gin.Context) {
	token := c.Query("token")

	if _, exist := usersLoginInfo[token]; exist {
		c.JSON(http.StatusOK, Response{StatusCode: 0})
	} else {
		c.JSON(http.StatusOK, Response{StatusCode: 1, StatusMsg: "User doesn't exist"})
	}
}

// FollowList all users have same follow list
func FollowList(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("user_id")
	UID, _ := strconv.ParseUint(userid, 10, 32)

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		followList_, err := service.GetFollowList(int64(UID))
		followList, err := GenerateFollow(followList_)
		log.Println(followList)
		log.Println(err)
		log.Println(UID)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Get success"},
			UserList: followList,
		})
	}
}

// FollowerList all users have same follower list
func FollowerList(c *gin.Context) {
	token := c.Query("token")
	userid := c.Query("user_id")
	UID, _ := strconv.ParseUint(userid, 10, 32)

	if token == "" {
		c.JSON(http.StatusOK, Response{
			StatusCode: 1,
			StatusMsg:  "You haven't logged in yet",
		})
	} else {
		fanList_, err := service.GetFanList(int64(UID))
		fanList, err := GenerateFollower(fanList_)
		log.Println(fanList)
		log.Println(err)
		log.Println(UID)
		c.JSON(http.StatusOK, UserListResponse{
			Response: Response{StatusCode: 0, StatusMsg: "Get success"},
			UserList: fanList,
		})
	}
}

func GenerateFollower(followerList []model.User) ([]User, error) {
	userID := 4
	fancount, _ := service.GetFanCount(uint(userID))
	followcount, _ := service.GetFollowCount(uint(userID))

	res := make([]User, len(followerList))

	for i, v := range followerList {
		res[i].Id = int64(v.UserID)
		res[i].Name = v.UserName
		res[i].FollowCount = followcount
		res[i].FollowerCount = fancount
		res[i].IsFollow, _ = service.IsFollow(uint(userID), v.UserID)
	}
	return res, nil
}

func GenerateFollow(followList []model.User) ([]User, error) {
	userID := 1
	fancount, _ := service.GetFanCount(uint(userID))
	followcount, _ := service.GetFollowCount(uint(userID))

	res := make([]User, len(followList))

	for i, v := range followList {
		res[i].Id = int64(v.UserID)
		res[i].Name = v.UserName
		res[i].FollowCount = followcount
		res[i].FollowerCount = fancount
		res[i].IsFollow, _ = service.IsFollow(uint(userID), v.UserID)
	}
	return res, nil
}
