package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"fmt"
	"log"
	"soutube/graph/generated"
	"soutube/graph/model"
)

func (r *mutationResolver) CreateVideo(ctx context.Context, input *model.NewVideo) (*model.Video, error) {
	log.Println("bambang")
	video := model.Video{
		VideoURL:       input.VideoURL,
		VideoTitle:     input.VideoTitle,
		VideoDesc:      input.VideoDesc,
		VideoCat:       input.VideoCat,
		VideoThumb:     input.VideoThumb,
		PlaylistID:     input.PlaylistID,
		LikeCount:      input.LikeCount,
		DislikeCount:   input.DislikeCount,
		AgeRestriction: input.AgeRestriction,
		Visibility:     input.Visibility,
		LocationVideo:  input.LocationVideo,
		Status:         input.Status,
		Premium:        input.Premium,
		ViewCount:      input.ViewCount,
		UserID:         input.UserID,
		VideoDuration:  input.VideoDuration,
		Day:            input.Day,
		Month:          input.Month,
		Year:           input.Year,
		Hour:           input.Hour,
		Minute:         input.Minute,
		Second:         input.Second,
	}
	_, err := r.DB.Model(&video).Insert()
	//	Query(&video, `INSERT INTO videos(
	//video_url, video_title, video_desc, video_cat, video_thumb, playlist_id, like_count, dislike_count, age_restriction, visibility, location_video, status, premium, view_count, user_id)
	//VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`, video.VideoURL, video.VideoTitle, video.VideoDesc, video.VideoCat, video.VideoThumb, video.PlaylistID, video.LikeCount, video.DislikeCount, video.AgeRestriction, video.Visibility, video.LocationVideo, video.Status, video.Premium, video.ViewCount, video.UserID)
	if err != nil {
		log.Println(err)
		return nil, errors.New(video.VideoTitle)
	}
	return &video, nil
}

func (r *mutationResolver) UpdateVideo(ctx context.Context, id int, input *model.NewVideo) (*model.Video, error) {
	var video model.Video
	_, err := r.DB.Query(&video, `SELECT * FROM videos WHERE video_id = ?`, id)
	if err != nil {
		return nil, errors.New("Not Found")
	}
	video.VideoDuration = input.VideoDuration
	video.Second = input.Second
	video.Minute = input.Minute
	video.Hour = input.Hour
	video.Year = input.Year
	video.Month = input.Month
	video.Day = input.Day
	video.ViewCount = input.ViewCount
	video.UserID = input.UserID
	video.Visibility = input.Visibility
	video.LikeCount = input.LikeCount
	video.DislikeCount = input.DislikeCount
	video.VideoTitle = input.VideoTitle
	video.Premium = input.Premium
	video.Status = input.Status
	video.LocationVideo = input.LocationVideo
	video.AgeRestriction = input.AgeRestriction
	video.PlaylistID = input.PlaylistID
	video.VideoThumb = input.VideoThumb
	video.VideoCat = input.VideoCat
	video.VideoDesc = input.VideoDesc
	video.VideoURL = input.VideoURL

	_, upErr := r.DB.Model(&video).Where("video_id = ?", id).Update()

	if upErr != nil {
		log.Println(upErr)
		return nil, errors.New("Update Failed")
	}
	return &video, nil
}

func (r *mutationResolver) DeleteVideo(ctx context.Context, id int) (bool, error) {
	var commenter model.Video
	_, err := r.DB.Query(&commenter, `SELECT * FROM videos WHERE video_id = ?`, id)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, delErr := r.DB.Model(&commenter).Where("video_id = ?", id).Delete()
	if delErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateUser(ctx context.Context, input *model.NewUser) (*model.User, error) {
	if input.ImgURL != "" && input.Email != "" && input.UserName != "" {
		user := model.User{
			Membership:      input.Membership,
			ImgURL:          input.ImgURL,
			Email:           input.Email,
			UserName:        input.UserName,
			SubscriberCount: input.SubscriberCount,
			Day:             input.Day,
			Month:           input.Month,
			Year:            input.Year,
			Hour:            input.Hour,
			Minute:          input.Minute,
			Second:          input.Second,
		}
		_, err := r.DB.Model(&user).Insert()

		if err != nil {
			return nil, errors.New("Insert new user failed")
		}
		return &user, nil
	}
	return nil, errors.New("Please input all")
}

func (r *mutationResolver) UpdateUser(ctx context.Context, userID int, input *model.NewUser) (*model.User, error) {
	var commenter model.User
	_, err := r.DB.Query(&commenter, `SELECT * FROM users WHERE user_id = ?`, userID)
	if err != nil {
		return nil, errors.New("Not Found")
	}
	commenter.Membership = input.Membership
	commenter.SubscriberCount = input.SubscriberCount

	_, upErr := r.DB.Model(&commenter).Where("user_id = ?", userID).Update()

	if upErr != nil {
		log.Println(upErr)
		return nil, errors.New("Update Failed")
	}
	return &commenter, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreatePlaylist(ctx context.Context, input *model.NewPlaylist) (*model.Playlist, error) {
	playlist := model.Playlist{
		UserID:       input.UserID,
		PlaylistName: input.PlaylistName,
		Visibility:   input.Visibility,
		Day:          input.Day,
		Month:        input.Month,
		Year:         input.Year,
		Hour:         input.Hour,
		Minute:       input.Minute,
		Second:       input.Second,
	}

	_, err := r.DB.Model(&playlist).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert playlist failed")
	}
	return &playlist, nil
}

func (r *mutationResolver) UpdatePlaylist(ctx context.Context, id int, input *model.NewPlaylist) (*model.Playlist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeletePlaylist(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateList(ctx context.Context, input *model.NewlistPlaylist) (*model.ListPlaylist, error) {
	playlist := model.ListPlaylist{
		PlaylistID: input.PlaylistID,
		VideoID:    input.VideoID,
	}

	_, err := r.DB.Query(&playlist, `INSERT INTO listplaylist (playlist_id,video_id) VALUES (?,?)`, playlist.PlaylistID, playlist.VideoID)

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert playlist failed")
	}
	return &playlist, nil
}

func (r *mutationResolver) DeleteList(ctx context.Context, listID int) (bool, error) {
	var commenter model.ListPlaylist
	_, err := r.DB.Query(&commenter, `SELECT * FROM listplaylist WHERE list_id = ?`, listID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, delErr := r.DB.Model(&commenter).Where("list_id = ?", listID).Delete()
	if delErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateComment(ctx context.Context, input *model.NewComment) (*model.Comment, error) {
	comment := model.Comment{
		UserID:       input.UserID,
		VideoID:      input.VideoID,
		Comment:      input.Comment,
		LikeCount:    input.LikeCount,
		DislikeCount: input.DislikeCount,
		Day:          input.Day,
		Month:        input.Month,
		Year:         input.Year,
		Hour:         input.Hour,
		Minute:       input.Minute,
		Second:       input.Second,
	}

	_, err := r.DB.Model(&comment).Insert()

	if err != nil {
		log.Println(err)
		return nil, errors.New("Insert comment failed")
	}
	return &comment, nil
}

func (r *mutationResolver) UpdateComment(ctx context.Context, commentID int, comment string, day int, month int, year int, hour int, minute int, second int) (*model.Comment, error) {
	var commenter model.Comment
	_, err := r.DB.Query(&commenter, `SELECT * FROM comments WHERE comment_id = ?`, commentID)
	if err != nil {
		return nil, errors.New("Not Found")
	}
	commenter.Comment = comment
	commenter.Day = day
	commenter.Month = month
	commenter.Year = year
	commenter.Hour = hour
	commenter.Minute = minute
	commenter.Second = second

	_, upErr := r.DB.Model(&commenter).Where("comment_id = ?", commentID).Update()

	if upErr != nil {
		log.Println(upErr)
		return nil, errors.New("Update Failed")
	}
	return &commenter, nil
}

func (r *mutationResolver) UpdateLike(ctx context.Context, commentID int, likeCount int, dislikeCount int) (*model.Comment, error) {
	var commenter model.Comment
	_, err := r.DB.Query(&commenter, `SELECT * FROM comments WHERE comment_id = ?`, commentID)
	if err != nil {
		return nil, errors.New("Not Found")
	}
	commenter.LikeCount = likeCount
	commenter.DislikeCount = dislikeCount

	_, upErr := r.DB.Model(&commenter).Where("comment_id = ?", commentID).Update()

	if upErr != nil {
		log.Println(upErr)
		return nil, errors.New("Update Failed")
	}
	return &commenter, nil
}

func (r *mutationResolver) DeleteComment(ctx context.Context, id int) (bool, error) {
	var commenter model.Comment
	_, err := r.DB.Query(&commenter, `SELECT * FROM comments WHERE comment_id = ?`, id)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, delErr := r.DB.Model(&commenter).Where("comment_id = ?", id).Delete()
	if delErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateReply(ctx context.Context, input *model.NewReply) (*model.Reply, error) {
	reply := model.Reply{
		Reply:     input.Reply,
		UserID:    input.UserID,
		CommentID: input.CommentID,
		Day:       input.Day,
		Month:     input.Month,
		Year:      input.Year,
		Hour:      input.Hour,
		Minute:    input.Minute,
		Second:    input.Second,
	}
	_, err := r.DB.Model(&reply).Insert()

	if err != nil {
		return nil, errors.New("Insert new user failed")
	}
	return &reply, nil
}

func (r *mutationResolver) UpdateReply(ctx context.Context, id int, input *model.NewReply) (*model.Reply, error) {
	var reply model.Reply
	_, err := r.DB.Query(&reply, `SELECT * FROM replies WHERE reply_id = ?`, id)
	if err != nil {
		return nil, errors.New("Not Found")
	}
	reply.Reply = input.Reply
	reply.Day = input.Day
	reply.Month = input.Month
	reply.Year = input.Year
	reply.Hour = input.Hour
	reply.Minute = input.Minute
	reply.Second = input.Second

	_, upErr := r.DB.Model(&reply).Where("reply_id = ?", id).Update()

	if upErr != nil {
		log.Println(upErr)
		return nil, errors.New("Update Failed")
	}
	return &reply, nil
}

func (r *mutationResolver) DeleteReply(ctx context.Context, id int) (bool, error) {
	var reply model.Reply
	_, err := r.DB.Query(&reply, `SELECT * FROM replies WHERE reply_id = ?`, id)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Model(&reply).Where("reply_id = ?", id).Delete()
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateLikedvid(ctx context.Context, input *model.NewLikedvid) (*model.Likedvid, error) {
	liked := model.Likedvid{
		UserID:  input.UserID,
		VideoID: input.VideoID,
	}
	_, err := r.DB.Query(&liked, `INSERT INTO likedvideos (user_id,video_id) VALUES (?,?)`, liked.UserID, liked.VideoID)

	if err != nil {
		return nil, errors.New("Insert new liked failed")
	}
	return &liked, nil
}

func (r *mutationResolver) UpdateLikedvid(ctx context.Context, likeID int, input *model.NewLikedvid) (*model.Likedvid, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteLikedvid(ctx context.Context, videoID int, userID int) (bool, error) {
	var liked model.Likedvid
	_, err := r.DB.Query(&liked, `SELECT * FROM likedvideos WHERE video_id = ? AND user_id = ?`, videoID, userID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Query(&liked, `DELETE FROM likedvideos WHERE video_id = ? AND user_id = ?`, videoID, userID)
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateDislikedvid(ctx context.Context, input *model.NewDislikedvid) (*model.Dislikedvid, error) {
	disliked := model.Dislikedvid{
		UserID:  input.UserID,
		VideoID: input.VideoID,
	}
	_, err := r.DB.Query(&disliked, `INSERT INTO dislikedvideos (user_id,video_id) VALUES (?,?)`, disliked.UserID, disliked.VideoID)

	if err != nil {
		return nil, errors.New("Insert new disliked failed")
	}
	return &disliked, nil
}

func (r *mutationResolver) UpdateDislikedvid(ctx context.Context, dislikeID int, input *model.NewDislikedvid) (*model.Dislikedvid, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDislikedvid(ctx context.Context, videoID int, userID int) (bool, error) {
	var disliked model.Dislikedvid
	_, err := r.DB.Query(&disliked, `SELECT * FROM dislikedvideos WHERE video_id = ? AND user_id = ?`, videoID, userID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Query(&disliked, `DELETE FROM dislikedvideos WHERE video_id = ? AND user_id = ?`, videoID, userID)
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateLikedcomment(ctx context.Context, input *model.NewLikedcomment) (*model.Likedcomment, error) {
	liked := model.Likedcomment{
		UserID:    input.UserID,
		CommentID: input.CommentID,
	}
	_, err := r.DB.Query(&liked, `INSERT INTO likedcomment (user_id,comment_id) VALUES (?,?)`, liked.UserID, liked.CommentID)

	if err != nil {
		return nil, errors.New("Insert new disliked failed")
	}
	return &liked, nil
}

func (r *mutationResolver) UpdateLikedcomment(ctx context.Context, likeID int, input *model.NewLikedcomment) (*model.Likedcomment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteLikedcomment(ctx context.Context, commentID int, userID int) (bool, error) {
	var disliked model.Likedcomment
	_, err := r.DB.Query(&disliked, `SELECT * FROM likedcomment WHERE comment_id = ? AND user_id = ?`, commentID, userID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Query(&disliked, `DELETE FROM likedcomment WHERE comment_id = ? AND user_id = ?`, commentID, userID)
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateDislikedcomment(ctx context.Context, input *model.NewDislikedcomment) (*model.Dislikedcomment, error) {
	liked := model.Dislikedcomment{
		UserID:    input.UserID,
		CommentID: input.CommentID,
	}
	_, err := r.DB.Query(&liked, `INSERT INTO dislikedcomment (user_id,comment_id) VALUES (?,?)`, liked.UserID, liked.CommentID)

	if err != nil {
		return nil, errors.New("Insert new disliked failed")
	}
	return &liked, nil
}

func (r *mutationResolver) UpdateDislikedcomment(ctx context.Context, dislikeID int, input *model.NewDislikedcomment) (*model.Dislikedcomment, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDislikedcomment(ctx context.Context, commentID int, userID int) (bool, error) {
	var disliked model.Dislikedcomment
	_, err := r.DB.Query(&disliked, `SELECT * FROM dislikedcomment WHERE comment_id = ? AND user_id = ?`, commentID, userID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Query(&disliked, `DELETE FROM dislikedcomment WHERE comment_id = ? AND user_id = ?`, commentID, userID)
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateLikedreply(ctx context.Context, input *model.NewLikereply) (*model.Likedreply, error) {
	liked := model.Likedreply{
		UserID:  input.UserID,
		ReplyID: input.ReplyID,
	}
	_, err := r.DB.Query(&liked, `INSERT INTO likedreply (user_id,reply_id) VALUES (?,?)`, liked.UserID, liked.ReplyID)

	if err != nil {
		return nil, errors.New("Insert new liked failed")
	}
	return &liked, nil
}

func (r *mutationResolver) UpdateLikedreply(ctx context.Context, likeID int, input *model.NewLikereply) (*model.Likedreply, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteLikedreply(ctx context.Context, replyID int, userID int) (bool, error) {
	var disliked model.Likedreply
	_, err := r.DB.Query(&disliked, `SELECT * FROM likedreply WHERE reply_id = ? AND user_id = ?`, replyID, userID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Query(&disliked, `DELETE FROM likedreply WHERE reply_id = ? AND user_id = ?`, replyID, userID)
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateDislikedreply(ctx context.Context, input *model.NewDislikereply) (*model.Dislikedreply, error) {
	liked := model.Dislikedreply{
		UserID:  input.UserID,
		ReplyID: input.ReplyID,
	}
	_, err := r.DB.Query(&liked, `INSERT INTO dislikedreply (user_id,reply_id) VALUES (?,?)`, liked.UserID, liked.ReplyID)

	if err != nil {
		return nil, errors.New("Insert new liked failed")
	}
	return &liked, nil
}

func (r *mutationResolver) UpdateDislikedreply(ctx context.Context, dislikeID int, input *model.NewDislikereply) (*model.Dislikedreply, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) DeleteDislikedreply(ctx context.Context, replyID int, userID int) (bool, error) {
	var disliked model.Dislikedreply
	_, err := r.DB.Query(&disliked, `SELECT * FROM dislikedreply WHERE reply_id = ? AND user_id = ?`, replyID, userID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Query(&disliked, `DELETE FROM dislikedreply WHERE reply_id = ? AND user_id = ?`, replyID, userID)
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *mutationResolver) CreateView(ctx context.Context, input *model.NewView) (*model.View, error) {
	liked := model.View{
		UserID:    input.UserID,
		VideoID:   input.VideoID,
		ViewCount: input.ViewCount,
	}
	_, err := r.DB.Query(&liked, `INSERT INTO viewvideos (user_id,video_id,view_count) VALUES (?,?,?)`, liked.UserID, liked.VideoID, liked.ViewCount)

	if err != nil {
		return nil, errors.New("Insert new liked failed")
	}
	return &liked, nil
}

func (r *mutationResolver) UpdateView(ctx context.Context, viewID int, input *model.NewView) (*model.View, error) {
	var reply model.View
	_, err := r.DB.Query(&reply, `SELECT * FROM viewvideos WHERE view_id = ?`, viewID)
	if err != nil {
		return nil, errors.New("Not Found")
	}
	reply.ViewCount = input.ViewCount

	_, upErr := r.DB.Query(&reply, `UPDATE viewvideos
	SET view_count=?
	WHERE view_id = ?;`, input.ViewCount, viewID)

	if upErr != nil {
		log.Println(upErr)
		return nil, errors.New("Update Failed")
	}
	return &reply, nil
}

func (r *mutationResolver) CreateSubs(ctx context.Context, input *model.NewSubscriber) (*model.Subscriber, error) {
	liked := model.Subscriber{
		UserID:       input.UserID,
		SubscriberID: input.SubscriberID,
	}
	_, err := r.DB.Query(&liked, `INSERT INTO subscribers (user_id,subscriber_id) VALUES (?,?)`, liked.UserID, liked.SubscriberID)

	if err != nil {
		return nil, errors.New("Insert new liked failed")
	}
	return &liked, nil
}

func (r *mutationResolver) DeleteSubs(ctx context.Context, userID int, subscriberID int) (bool, error) {
	var disliked model.Subscriber
	_, err := r.DB.Query(&disliked, `SELECT * FROM subscribers WHERE user_id = ? AND subscriber_id = ?`, userID, subscriberID)
	if err != nil {
		return false, errors.New("Not Found")
	}
	_, upErr := r.DB.Query(&disliked, `DELETE FROM subscribers WHERE user_id = ? AND subscriber_id = ?`, userID, subscriberID)
	if upErr != nil {
		return false, errors.New("Delete Failed")
	}
	return true, nil
}

func (r *queryResolver) Videos(ctx context.Context) ([]*model.Video, error) {
	var videos []*model.Video

	err := r.DB.Model(&videos).Select()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {
	var users []*model.User

	err := r.DB.Model(&users).Select()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query users")
	}

	return users, nil
}

func (r *queryResolver) Playlists(ctx context.Context) ([]*model.Playlist, error) {
	var playlists []*model.Playlist

	err := r.DB.Model(&playlists).Select()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return playlists, nil
}

func (r *queryResolver) Replies(ctx context.Context) ([]*model.Reply, error) {
	var replies []*model.Reply

	err := r.DB.Model(&replies).Select()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return replies, nil
}

func (r *queryResolver) Comments(ctx context.Context) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Select()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}
	return comments, nil
}

func (r *queryResolver) Likedvids(ctx context.Context) ([]*model.Likedvid, error) {
	var liked []*model.Likedvid

	_, err := r.DB.Query(&liked, `SELECT * FROM likedvideos`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query liked")
	}
	return liked, nil
}

func (r *queryResolver) Dislikedvids(ctx context.Context) ([]*model.Dislikedvid, error) {
	var disliked []*model.Dislikedvid

	_, err := r.DB.Query(&disliked, `SELECT * FROM dislikedvideos`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query disliked")
	}
	return disliked, nil
}

func (r *queryResolver) Likedcomments(ctx context.Context) ([]*model.Likedcomment, error) {
	var liked []*model.Likedcomment

	_, err := r.DB.Query(&liked, `SELECT * FROM likedComment`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query liked")
	}
	return liked, nil
}

func (r *queryResolver) Dislikedcomments(ctx context.Context) ([]*model.Dislikedcomment, error) {
	var liked []*model.Dislikedcomment

	_, err := r.DB.Query(&liked, `SELECT * FROM dislikedComment`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query liked")
	}
	return liked, nil
}

func (r *queryResolver) Likedreply(ctx context.Context) ([]*model.Likedreply, error) {
	var liked []*model.Likedreply

	_, err := r.DB.Query(&liked, `SELECT * FROM likedreply`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query liked")
	}
	return liked, nil
}

func (r *queryResolver) Dislikedreply(ctx context.Context) ([]*model.Dislikedreply, error) {
	var liked []*model.Dislikedreply

	_, err := r.DB.Query(&liked, `SELECT * FROM dislikedreply`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query liked")
	}
	return liked, nil
}

func (r *queryResolver) Views(ctx context.Context) ([]*model.View, error) {
	var liked []*model.View

	_, err := r.DB.Query(&liked, `SELECT * FROM viewvideos`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query liked")
	}
	return liked, nil
}

func (r *queryResolver) Subscribers(ctx context.Context) ([]*model.Subscriber, error) {
	var liked []*model.Subscriber

	_, err := r.DB.Query(&liked, `SELECT * FROM subscribers`)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query liked")
	}
	return liked, nil
}

func (r *queryResolver) ListPlaylists(ctx context.Context) ([]*model.ListPlaylist, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) GetUserID(ctx context.Context, email string) (*model.User, error) {
	var user model.User
	_, err := r.DB.Query(&user, `SELECT * FROM users WHERE email = ?`, email)
	if err != nil {
		fmt.Println(err)
		fmt.Println(email)
		return nil, errors.New("Not Found")
	}
	return &user, nil
}

func (r *queryResolver) GetUser(ctx context.Context, userID int) (*model.User, error) {
	var user model.User
	_, err := r.DB.Query(&user, `SELECT * FROM users WHERE user_id = ?`, userID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(userID)
		return nil, errors.New("Not Found")
	}
	return &user, nil
}

func (r *queryResolver) GetVideo(ctx context.Context, videoID int) (*model.Video, error) {
	var video model.Video
	_, err := r.DB.Query(&video, `SELECT * FROM videos WHERE video_id = ?`, videoID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(videoID)
		return nil, errors.New("Not Found")
	}
	return &video, nil
}

func (r *queryResolver) GetVideoUser(ctx context.Context, userID int) ([]*model.Video, error) {
	var video []*model.Video
	_, err := r.DB.Query(&video, `SELECT * FROM videos WHERE user_id = ?`, userID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(userID)
		return nil, errors.New("Not Found")
	}
	return video, nil
}

func (r *queryResolver) GetAllComment(ctx context.Context) ([]*model.Comment, error) {
	var comments []*model.Comment

	err := r.DB.Model(&comments).Select()
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return comments, nil
}

func (r *queryResolver) GetComment(ctx context.Context, videoID int) ([]*model.Comment, error) {
	var comment []*model.Comment
	_, err := r.DB.Query(&comment, `SELECT * FROM comments WHERE video_id = ?`, videoID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(videoID)
		return nil, errors.New("Not Found")
	}
	return comment, nil
}

func (r *queryResolver) GetCommentByID(ctx context.Context, commentID int) (*model.Comment, error) {
	var comment model.Comment
	_, err := r.DB.Query(&comment, `SELECT * FROM comments WHERE comment_id = ?`, commentID)
	if err != nil {
		return nil, errors.New("Not Found")
	}
	return &comment, nil
}

func (r *queryResolver) GetPlayListUser(ctx context.Context, userID int) ([]*model.Playlist, error) {
	var playlist []*model.Playlist
	_, err := r.DB.Query(&playlist, `SELECT * FROM playlists WHERE user_id = ?`, userID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(userID)
		return nil, errors.New("Not Found")
	}
	return playlist, nil
}

func (r *queryResolver) GetReplyComment(ctx context.Context, commentID int) ([]*model.Reply, error) {
	var reply []*model.Reply
	_, err := r.DB.Query(&reply, `SELECT * FROM replies WHERE comment_id = ?`, commentID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(commentID)
		return nil, errors.New("Not Found")
	}
	return reply, nil
}

func (r *queryResolver) GetLikedVideo(ctx context.Context, videoID int) ([]*model.Likedvid, error) {
	var liked []*model.Likedvid
	_, err := r.DB.Query(&liked, `SELECT * FROM likedvideos WHERE video_id = ?`, videoID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(videoID)
		return nil, errors.New("Not Found")
	}
	return liked, nil
}

func (r *queryResolver) CheckLikedVideo(ctx context.Context, videoID int, userID int) (*model.Likedvid, error) {
	var liked model.Likedvid
	_, err := r.DB.Query(&liked, `SELECT * FROM likedvideos WHERE video_id = ? AND user_id = ?`, videoID, userID)
	if &liked == nil {
		fmt.Println(err)
		fmt.Println(videoID)
		return nil, errors.New("Not Found")
	}
	return &liked, nil
}

func (r *queryResolver) GetDislikedVideo(ctx context.Context, videoID int) ([]*model.Dislikedvid, error) {
	var disliked []*model.Dislikedvid
	_, err := r.DB.Query(&disliked, `SELECT * FROM dislikedvideos WHERE video_id = ?`, videoID)
	if err != nil {
		fmt.Println(err)
		fmt.Println(videoID)
		return nil, errors.New("Not Found")
	}
	return disliked, nil
}

func (r *queryResolver) CheckDislikedVideo(ctx context.Context, videoID int, userID int) (*model.Dislikedvid, error) {
	var disliked model.Dislikedvid
	_, err := r.DB.Query(&disliked, `SELECT * FROM dislikedvideos WHERE video_id = ? AND user_id = ?`, videoID, userID)
	if &disliked == nil {
		fmt.Println(err)
		fmt.Println(videoID)
		return nil, errors.New("Not Found")
	}
	return &disliked, nil
}

func (r *queryResolver) GetLikedComment(ctx context.Context, commentID int) ([]*model.Likedcomment, error) {
	var liked []*model.Likedcomment
	_, err := r.DB.Query(&liked, `SELECT * FROM likedcomment WHERE comment_id = ?`, commentID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return liked, nil
}

func (r *queryResolver) CheckLikedComment(ctx context.Context, commentID int, userID int) (*model.Likedcomment, error) {
	var disliked model.Likedcomment
	_, err := r.DB.Query(&disliked, `SELECT * FROM likedcomment WHERE comment_id = ? AND user_id = ?`, commentID, userID)
	if &disliked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return &disliked, nil
}

func (r *queryResolver) GetDislikedComment(ctx context.Context, commentID int) ([]*model.Dislikedcomment, error) {
	var liked []*model.Dislikedcomment
	_, err := r.DB.Query(&liked, `SELECT * FROM dislikedcomment WHERE comment_id = ?`, commentID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return liked, nil
}

func (r *queryResolver) CheckDislikedComment(ctx context.Context, commentID int, userID int) (*model.Dislikedcomment, error) {
	var disliked model.Dislikedcomment
	_, err := r.DB.Query(&disliked, `SELECT * FROM dislikedcomment WHERE comment_id = ? AND user_id = ?`, commentID, userID)
	if &disliked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return &disliked, nil
}

func (r *queryResolver) GetLikedReply(ctx context.Context, replyID int) ([]*model.Likedreply, error) {
	var liked []*model.Likedreply
	_, err := r.DB.Query(&liked, `SELECT * FROM likedreply WHERE reply_id = ?`, replyID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return liked, nil
}

func (r *queryResolver) CheckLikedReply(ctx context.Context, replyID int, userID int) (*model.Likedreply, error) {
	var liked model.Likedreply
	_, err := r.DB.Query(&liked, `SELECT * FROM likedreply WHERE reply_id = ? AND user_id = ?`, replyID, userID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return &liked, nil
}

func (r *queryResolver) GetDislikedReply(ctx context.Context, replyID int) ([]*model.Dislikedreply, error) {
	var liked []*model.Dislikedreply
	_, err := r.DB.Query(&liked, `SELECT * FROM dislikedreply WHERE reply_id = ?`, replyID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return liked, nil
}

func (r *queryResolver) CheckDislikedReply(ctx context.Context, replyID int, userID int) (*model.Dislikedreply, error) {
	var liked model.Dislikedreply
	_, err := r.DB.Query(&liked, `SELECT * FROM dislikedreply WHERE reply_id = ? AND user_id = ?`, replyID, userID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return &liked, nil
}

func (r *queryResolver) GetView(ctx context.Context, videoID int) ([]*model.View, error) {
	var liked []*model.View
	_, err := r.DB.Query(&liked, `SELECT * FROM viewvideos WHERE video_id = ?`, videoID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return liked, nil
}

func (r *queryResolver) CheckView(ctx context.Context, videoID int, userID int) (*model.View, error) {
	var liked model.View
	_, err := r.DB.Query(&liked, `SELECT * FROM viewvideos WHERE video_id = ? AND user_id = ?`, videoID, userID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return &liked, nil
}

func (r *queryResolver) GetList(ctx context.Context, playlistID int) ([]*model.ListPlaylist, error) {
	var liked []*model.ListPlaylist
	_, err := r.DB.Query(&liked, `SELECT * FROM listplaylist WHERE playlist_id = ?`, playlistID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return liked, nil
}

func (r *queryResolver) CheckList(ctx context.Context, videoID int) (*model.ListPlaylist, error) {
	var liked model.ListPlaylist
	_, err := r.DB.Query(&liked, `SELECT * FROM listplaylist WHERE video_id = ? `, videoID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return &liked, nil
}

func (r *queryResolver) GetPlaylist(ctx context.Context, playlistID int) (*model.Playlist, error) {
	var liked model.Playlist
	_, err := r.DB.Query(&liked, `SELECT * FROM playlists WHERE playlist_id = ? `, playlistID)
	if &liked == nil {
		fmt.Println(err)
		return nil, errors.New("Not Found")
	}
	return &liked, nil
}

func (r *queryResolver) CheckPlaylistLike(ctx context.Context, playlistName string) ([]*model.Playlist, error) {
	var videos []*model.Playlist

	_, err := r.DB.Query(&videos, `SELECT * FROM playlists WHERE LOWER(playlist_name) LIKE LOWER(?)`, playlistName)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

func (r *queryResolver) CheckVideoLike(ctx context.Context, videoTitle string) ([]*model.Video, error) {
	var videos []*model.Video

	_, err := r.DB.Query(&videos, `SELECT * FROM videos WHERE LOWER(video_title) LIKE LOWER(?)`, videoTitle)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

func (r *queryResolver) CheckChannelLike(ctx context.Context, userName string) ([]*model.User, error) {
	var videos []*model.User

	_, err := r.DB.Query(&videos, `SELECT * FROM users WHERE LOWER(user_name) LIKE LOWER(?)`, userName)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

func (r *queryResolver) GetSubscriber(ctx context.Context, userID int) ([]*model.Subscriber, error) {
	var videos []*model.Subscriber

	_, err := r.DB.Query(&videos, `SELECT * FROM subscribers WHERE user_id = ?`, userID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

func (r *queryResolver) CheckSubs(ctx context.Context, userID int, subscriberID int) (*model.Subscriber, error) {
	var videos model.Subscriber

	_, err := r.DB.Query(&videos, `SELECT * FROM subscribers WHERE user_id = ? AND subscriber_id = ?`, userID, subscriberID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query subscriber")
	}

	return &videos, nil
}

func (r *queryResolver) GetMySubs(ctx context.Context, subscriberID int) ([]*model.Subscriber, error) {
	var videos []*model.Subscriber

	_, err := r.DB.Query(&videos, `SELECT * FROM subscribers WHERE subscriber_id = ?`, subscriberID)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

func (r *queryResolver) GetRelated(ctx context.Context, videoCat string, locationVideo string) ([]*model.Video, error) {
	var videos []*model.Video

	_, err := r.DB.Query(&videos, `SELECT * FROM videos WHERE video_cat = ? AND location_video = ?`, videoCat, locationVideo)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

func (r *queryResolver) GetRelatedCat(ctx context.Context, videoCat string) ([]*model.Video, error) {
	var videos []*model.Video

	_, err := r.DB.Query(&videos, `SELECT * FROM videos WHERE video_cat = ? `, videoCat)
	if err != nil {
		fmt.Println(err)
		return nil, errors.New("Failed to Query videos")
	}

	return videos, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
