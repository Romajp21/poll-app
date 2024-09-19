package poll

type Poll struct {
	Id       int
	Question string
	Option   []Optio
}

type Optio struct {
	Id     int
	PollId int
	Text   string
	Votes  int
}
