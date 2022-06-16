package controller

import "simple-demo/define"

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            2,
		Author:        DemoUser,
		PlayUrl:       define.VideoRoot + "6889aed365f7470edfff41ed3bb7a978.mp4",
		CoverUrl:      define.ImgRoot + "bear.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            3,
		Author:        DemoUser,
		PlayUrl:       define.VideoRoot + "f0967b94811b88ccbffc6c034cb93b67.mp4",
		CoverUrl:      define.ImgRoot + "bear.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
	{
		Id:            4,
		Author:        DemoUser,
		PlayUrl:       define.VideoRoot + "eb2bdb5d435a8ea675c694c7ea7489fd.mp4",
		CoverUrl:      define.ImgRoot + "bear.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
	},
}

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
	},
}

var DemoUser = User{
	Id:            1,
	Name:          "TestUser",
	FollowCount:   0,
	FollowerCount: 0,
	IsFollow:      false,
}
