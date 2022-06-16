package service

import (
	"context"
	"math"
	"simple-demo/model"
)

// CreateVideo 创建视频信息
func CreateVideo(ctx context.Context, video *model.Video) error {
	return model.DB.WithContext(ctx).Create(&video).Error
}

// GetUserIDByToken 根据用户id查token
func GetUserIDByToken(token string) (*model.Users, error) {
	var user model.Users
	if err := model.DB.Where("token = ?", token).First(&user).Error; err != nil {
		return &user, err
	}
	return &user, nil
}

// GetVideoByUserID 根据用户id查视频
func GetVideoByUserID(userID uint) ([]model.Videos, error) {
	userid := model.Users{Id: int64(userID)}.Id
	var videoList []model.Videos
	if err := model.DB.Where("user_id=?", userid).Find(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// GetVideoByLoginToken 根据用户token提供视频
func GetVideoByLoginToken(token string) ([]model.Videos, error) {
	userinfo, _ := GetUserIDByToken(token)
	userid := userinfo.Id
	var videoList []model.Videos
	if err := model.DB.Where("user_id!=?", userid).Find(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// GetVideoByNoLoginToken 给非登录用户提供视频
func GetVideoByNoLoginToken() ([]model.Videos, error) {
	var videoList []model.Videos
	if err := model.DB.Select("*").Find(&videoList).Error; err != nil {
		return nil, err
	}
	return videoList, nil
}

// GetLikeCount 返回视频点赞数
func GetLikeCount(ctx context.Context, videoID int64) (int64, error) {
	video := model.Video{VideoID: uint(videoID)}
	return model.DB.WithContext(ctx).Model(&video).Association("Likes").Count(), nil
}

// GetCommentCount 返回视频评论数
func GetCommentCount(ctx context.Context, videoID int64) (int64, error) {
	video := model.Video{VideoID: uint(videoID)}
	return model.DB.WithContext(ctx).Model(&video).Association("Comments").Count(), nil
}

// IsFavorite 返回是否点赞
func IsFavorite(ctx context.Context, videoID int64, userID int64) (bool, error) {
	user := model.User{UserID: uint(userID)}
	return model.DB.WithContext(ctx).Model(&user).Where("`like`.video_id = ?", videoID).Association("Likes").Count() > 0, nil
}

// GetVideoByTime 根据时间戳返回最近count个视频,还需要返回next time
func GetVideoByTime(ctx context.Context, latestTime int64, count int64) ([]*model.Video, int64, error) {
	var videos []*model.Video
	if err := model.DB.WithContext(ctx).Where("created_at < ?", latestTime).Limit(int(count)).Order("created_at DESC").Find(&videos).Error; err != nil {
		return nil, 0, err
	}
	var nextTime int64 = math.MaxInt32
	if len(videos) != 0 { // 查到了新视频
		nextTime = int64(videos[0].CreatedAt)
	}
	return videos, nextTime, nil
}

// LikeVideo 点赞视频
func LikeVideo(ctx context.Context, userID int64, videoID int64) error {
	user := model.User{UserID: uint(userID)}
	video := model.Video{VideoID: uint(videoID)}
	return model.DB.WithContext(ctx).Model(&user).Association("Likes").Append(&video)
}

// UnLikeVideo 取消点赞视频
func UnLikeVideo(ctx context.Context, userID int64, videoID int64) error {
	user := model.User{UserID: uint(userID)}
	video := model.Video{VideoID: uint(videoID)}
	return model.DB.WithContext(ctx).Model(&user).Association("Likes").Delete(&video)
}

func GetLikeVideo(userID int64) ([]model.Videos, error) {
	//var favoriteList []model.Favorites
	//var videoList []model.Videos
	//model.DB.Where("user_id=?", userID).Find(&favoriteList)
	//
	//model.DB.Table("favorites").Select("favorites.video_id,videos.user_id,videos.title,videos.play_url,videos.cover_url").
	//	Where("favorites.user_id=?", userID).
	//	Joins("LEFT JOIN videos ON favorites.video_id = videos.video_id").
	//	Find(&videoList)

	UserId := userID
	var favoriteList []model.Favorites
	var videoList []model.Videos
	model.DB.Where("user_id=?", UserId).Find(&favoriteList)

	model.DB.Table("favorites").Select("favorites.video_id,videos.user_id,videos.title,videos.play_url,videos.cover_url").
		Where("favorites.user_id=?", UserId).
		Joins("LEFT JOIN videos ON favorites.video_id = videos.video_id").
		Find(&videoList)
	//fmt.Println(videoList)
	return videoList, nil
}

// CreateComment 新增评论,需要dal层返回评论详情
func CreateComment(ctx context.Context, userID int64, videoID int64, content string) (*model.Comment, error) {
	comment := model.Comment{
		UserID:  uint(userID),
		VideoID: uint(videoID),
		Content: content,
	}
	if err := model.DB.WithContext(ctx).Create(&comment).Error; err != nil {
		return nil, err
	}
	return &comment, nil
}

// DeleteComment 删除评论
func DeleteComment(ctx context.Context, commentID int64) error {
	return model.DB.WithContext(ctx).Delete(&model.Comment{}, commentID).Error
}

// GetComment 查询评论,需要dal层返回评论详情,有可能有多条评论
func GetComment(ctx context.Context, videoID int64) ([]*model.Comment, error) {
	video := model.Video{VideoID: uint(videoID)}
	if err := model.DB.WithContext(ctx).Preload("Comments").Find(&video).Error; err != nil {
		return nil, err
	}
	return video.Comments, nil
}
