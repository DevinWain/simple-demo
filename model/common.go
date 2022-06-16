package model

//type Response struct {
//	StatusCode int32  `json:"status_code"`
//	StatusMsg  string `json:"status_msg,omitempty"`
//}

type Videos struct {
	Id     int64  `gorm:"primaryKey;column:video_id" json:"video_id,omitempty"`
	Author Users  `gorm:"foreignKey;column:user_id,column:user_name" json:"author"`
	UserId int64  `gorm:"primaryKey;column:user_id" json:"user_id,omitempty"`
	Title  string `gorm:"column:title"`
	//CreatedAt int
	PlayUrl  string `gorm:"column:play_url" json:"play_url" json:"play_url,omitempty"`
	CoverUrl string `gorm:"column:cover_url" json:"cover_url,omitempty"`

	FavoriteCount int64 `gorm:"primaryKey;column:userid" json:"favorite_count,omitempty"`
	CommentCount  int64 `gorm:"primaryKey;column:userid" json:"comment_count,omitempty"`
	IsFavorite    bool  `gorm:"primaryKey;column:userid" json:"is_favorite,omitempty"`
}

//type MComment struct {
//	Id         int64  `json:"id,omitempty"`
//	User       User   `json:"user"`
//	Content    string `json:"content,omitempty"`
//	CreateDate string `json:"create_date,omitempty"`
//}

type Users struct {
	Id       int64  `gorm:"primaryKey;column:user_id" `
	Name     string `gorm:"column:user_name;not null" `
	Password string `gorm:"column:password;not null"`
	//Token         string `gorm:"column:token;not null"`
	FollowCount   int64 `gorm:"column:follow_count" json:"follow_count,omitempty;default:0"`
	FollowerCount int64 `gorm:"column:follower_count" json:"follower_count,omitempty;default:0"`
	IsFollow      bool  `gorm:"column:is_follow" json:"is_follow,omitempty;default:false"`
}

type Favorites struct {
	Id      int64  `gorm:"primaryKey"`
	User    Users  `gorm:"foreignKey:user_id"`
	UserId  int64  `gorm:"not null"`
	Video   Videos `gorm:"foreignKey:video_id"`
	VideoId int64  `gorm:"not null"`
}
