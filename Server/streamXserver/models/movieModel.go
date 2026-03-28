package model

import (
	"go.mongodb.org/mongo-driver/v2/bson"
)

type Genre struct{
	GernreId int
	GenreName string
}
type Ranking struct{
	RankingValue int
	RankingName string
}


type Movie struct {
	ID bson.ObjectID `bson:"_id" json:"_id"`//giving excat perticular id name as database so when we r gonna call it matches according to atabase
	ImdbID string
	Title string
	PosterPath string
	YouTubeID string
	Genre []Genre
	AdminReview string
	Ranking Ranking
}