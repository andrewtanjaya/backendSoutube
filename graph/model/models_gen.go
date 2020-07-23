// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Comment struct {
	CommentID    string `json:"comment_id"`
	UserID       int    `json:"user_id"`
	VideoID      int    `json:"video_id"`
	Comment      string `json:"comment"`
	LikeCount    int    `json:"like_count"`
	DislikeCount int    `json:"dislike_count"`
	Day          int    `json:"day"`
	Month        int    `json:"month"`
	Year         int    `json:"year"`
	Hour         int    `json:"hour"`
	Minute       int    `json:"minute"`
	Second       int    `json:"second"`
}

type Dislikedcomment struct {
	DislikeID string `json:"dislike_id"`
	UserID    int    `json:"user_id"`
	CommentID int    `json:"comment_id"`
}

type Dislikedreply struct {
	DislikeID string `json:"dislike_id"`
	UserID    int    `json:"user_id"`
	ReplyID   int    `json:"reply_id"`
}

type Dislikedvid struct {
	DislikeID string `json:"dislike_id"`
	UserID    int    `json:"user_id"`
	VideoID   int    `json:"video_id"`
}

type Likedcomment struct {
	LikeID    string `json:"like_id"`
	UserID    int    `json:"user_id"`
	CommentID int    `json:"comment_id"`
}

type Likedreply struct {
	LikeID  string `json:"like_id"`
	UserID  int    `json:"user_id"`
	ReplyID int    `json:"reply_id"`
}

type Likedvid struct {
	LikeID  string `json:"like_id"`
	UserID  int    `json:"user_id"`
	VideoID int    `json:"video_id"`
}

type Playlist struct {
	PlaylistID   string `json:"playlist_id"`
	PlaylistName string `json:"playlist_name"`
	UserID       int    `json:"user_id"`
	Visibility   string `json:"visibility"`
	Day          int    `json:"day"`
	Month        int    `json:"month"`
	Year         int    `json:"year"`
	Hour         int    `json:"hour"`
	Minute       int    `json:"minute"`
	Second       int    `json:"second"`
}

type Reply struct {
	ReplyID   string `json:"reply_id"`
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	Reply     string `json:"reply"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
	Second    int    `json:"second"`
}

type Subscriber struct {
	SubsID       string `json:"subs_id"`
	UserID       int    `json:"user_id"`
	SubscriberID int    `json:"subscriber_id"`
}

type User struct {
	UserID          string `json:"user_id"`
	Membership      bool   `json:"membership"`
	ImgURL          string `json:"img_url"`
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
	SubscriberCount int    `json:"subscriber_count"`
	Day             int    `json:"day"`
	Month           int    `json:"month"`
	Year            int    `json:"year"`
	Hour            int    `json:"hour"`
	Minute          int    `json:"minute"`
	Second          int    `json:"second"`
}

type Video struct {
	VideoID        string `json:"video_id"`
	VideoURL       string `json:"video_url"`
	VideoTitle     string `json:"video_title"`
	VideoDesc      string `json:"video_desc"`
	VideoCat       string `json:"video_cat"`
	VideoThumb     string `json:"video_thumb"`
	PlaylistID     int    `json:"playlist_id"`
	LikeCount      int    `json:"like_count"`
	DislikeCount   int    `json:"dislike_count"`
	AgeRestriction string `json:"age_restriction"`
	Visibility     string `json:"visibility"`
	LocationVideo  string `json:"location_video"`
	Status         string `json:"status"`
	Premium        string `json:"premium"`
	ViewCount      int    `json:"view_count"`
	UserID         int    `json:"user_id"`
	VideoDuration  int    `json:"video_duration"`
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Year           int    `json:"year"`
	Hour           int    `json:"hour"`
	Minute         int    `json:"minute"`
	Second         int    `json:"second"`
}

type View struct {
	ViewID    string `json:"view_id"`
	UserID    int    `json:"user_id"`
	VideoID   int    `json:"video_id"`
	ViewCount int    `json:"view_count"`
}

type ListPlaylist struct {
	ListID     string `json:"list_id"`
	PlaylistID int    `json:"playlist_id"`
	VideoID    int    `json:"video_id"`
}

type NewComment struct {
	UserID       int    `json:"user_id"`
	VideoID      int    `json:"video_id"`
	Comment      string `json:"comment"`
	LikeCount    int    `json:"like_count"`
	DislikeCount int    `json:"dislike_count"`
	Day          int    `json:"day"`
	Month        int    `json:"month"`
	Year         int    `json:"year"`
	Hour         int    `json:"hour"`
	Minute       int    `json:"minute"`
	Second       int    `json:"second"`
}

type NewDislikedcomment struct {
	UserID    int `json:"user_id"`
	CommentID int `json:"comment_id"`
}

type NewDislikedvid struct {
	UserID  int `json:"user_id"`
	VideoID int `json:"video_id"`
}

type NewDislikereply struct {
	UserID  int `json:"user_id"`
	ReplyID int `json:"reply_id"`
}

type NewLikedcomment struct {
	UserID    int `json:"user_id"`
	CommentID int `json:"comment_id"`
}

type NewLikedvid struct {
	UserID  int `json:"user_id"`
	VideoID int `json:"video_id"`
}

type NewLikereply struct {
	UserID  int `json:"user_id"`
	ReplyID int `json:"reply_id"`
}

type NewPlaylist struct {
	PlaylistName string `json:"playlist_name"`
	UserID       int    `json:"user_id"`
	Visibility   string `json:"visibility"`
	Day          int    `json:"day"`
	Month        int    `json:"month"`
	Year         int    `json:"year"`
	Hour         int    `json:"hour"`
	Minute       int    `json:"minute"`
	Second       int    `json:"second"`
}

type NewReply struct {
	CommentID int    `json:"comment_id"`
	UserID    int    `json:"user_id"`
	Reply     string `json:"reply"`
	Day       int    `json:"day"`
	Month     int    `json:"month"`
	Year      int    `json:"year"`
	Hour      int    `json:"hour"`
	Minute    int    `json:"minute"`
	Second    int    `json:"second"`
}

type NewSubscriber struct {
	UserID       int `json:"user_id"`
	SubscriberID int `json:"subscriber_id"`
}

type NewUser struct {
	Membership      bool   `json:"membership"`
	ImgURL          string `json:"img_url"`
	Email           string `json:"email"`
	UserName        string `json:"user_name"`
	SubscriberCount int    `json:"subscriber_count"`
	Day             int    `json:"day"`
	Month           int    `json:"month"`
	Year            int    `json:"year"`
	Hour            int    `json:"hour"`
	Minute          int    `json:"minute"`
	Second          int    `json:"second"`
}

type NewVideo struct {
	VideoURL       string `json:"video_url"`
	VideoTitle     string `json:"video_title"`
	VideoDesc      string `json:"video_desc"`
	VideoCat       string `json:"video_cat"`
	VideoThumb     string `json:"video_thumb"`
	PlaylistID     int    `json:"playlist_id"`
	LikeCount      int    `json:"like_count"`
	DislikeCount   int    `json:"dislike_count"`
	AgeRestriction string `json:"age_restriction"`
	Visibility     string `json:"visibility"`
	LocationVideo  string `json:"location_video"`
	Status         string `json:"status"`
	Premium        string `json:"premium"`
	ViewCount      int    `json:"view_count"`
	UserID         int    `json:"user_id"`
	VideoDuration  int    `json:"video_duration"`
	Day            int    `json:"day"`
	Month          int    `json:"month"`
	Year           int    `json:"year"`
	Hour           int    `json:"hour"`
	Minute         int    `json:"minute"`
	Second         int    `json:"second"`
}

type NewView struct {
	UserID    int `json:"user_id"`
	VideoID   int `json:"video_id"`
	ViewCount int `json:"view_count"`
}

type NewlistPlaylist struct {
	PlaylistID int `json:"playlist_id"`
	VideoID    int `json:"video_id"`
}
