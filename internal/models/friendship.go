package models

type Friendship struct {
    UserID   int    `json:"user_id"`
    FriendID int    `json:"friend_id"`
    Status   string `json:"status"`
}