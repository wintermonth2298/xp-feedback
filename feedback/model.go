package feedback

type Feedback struct {
	ID      string `bson:"id,omitempty" json:"id"`
	Stars   int    `bson:"stars" json:"stars"`
	Comment string `bson:"comment" json:"comment"`
}
