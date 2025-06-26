package main

import (
	"time"

	"github.com/SiddhantKandi/HTTPServer/internal/database"
	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Name      string     `json:"name"`
	APIKey    string     `json:"api_key"`
}

type Feed struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	Name      string     `json:"name"`
	URL       string     `json:"url"`
	UserID    uuid.UUID  `json:"user_id"`
}

type FeedFollow struct {
	ID        uuid.UUID  `json:"id"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	UserID    uuid.UUID  `json:"user_id"`
	FeedID    uuid.UUID  `json:"feed_id"`
}

func databaseUsertoUser(dbUser database.User) User{
	return User{
		ID:dbUser.ID,
		CreatedAt:dbUser.CreatedAt,
		UpdatedAt:dbUser.UpdatedAt,
		Name:dbUser.Name,
		APIKey:dbUser.ApiKey,
	}
}


func databaseFeedtoFeed(dbFeed database.Feed) Feed{
	return Feed {
		ID:dbFeed.ID,
		CreatedAt:dbFeed.CreatedAt,
		UpdatedAt:dbFeed.UpdatedAt,
		Name:dbFeed.Name,
		URL:dbFeed.Url,
		UserID: dbFeed.UserID,
	}
}

func databaseFeedstoFeeds(dbFeeds []database.Feed) []Feed{
	feeds := []Feed{}

	for _, dbFeed := range dbFeeds{
		feeds = append(feeds, databaseFeedtoFeed(dbFeed))
	}

	return feeds
}

func databaseFeedFollowtoFeedFollow(dbFeed database.Feedfollow) FeedFollow{
	return FeedFollow {
		ID:dbFeed.ID,
		CreatedAt:dbFeed.CreatedAt,
		UpdatedAt:dbFeed.UpdatedAt,
		UserID: dbFeed.UserID,
		FeedID: dbFeed.FeedID,
	}
}