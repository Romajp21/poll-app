package poll

import (
	"poll-app/db"
)

func GetPolls() ([]Poll, error) {
	rows, err := db.DB.Query(`
        SELECT p.Id, p.Question, o.Id, o.Text, o.Votes 
        FROM Poll p
        LEFT JOIN Optio o ON p.Id = o.PollId`)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	pollsMap := make(map[int]*Poll)

	for rows.Next() {
		var pollId, optionId, votes int
		var question, text string

		err := rows.Scan(&pollId, &question, &optionId, &text, &votes)
		if err != nil {
			return nil, err
		}

		if poll, exists := pollsMap[pollId]; exists {
			poll.Option = append(poll.Option, Optio{Id: optionId, Text: text, Votes: votes})
		} else {
			pollsMap[pollId] = &Poll{
				Id:       pollId,
				Question: question,
				Option:   []Optio{{Id: optionId, Text: text, Votes: votes}},
			}
		}
	}

	var polls []Poll
	for _, poll := range pollsMap {
		polls = append(polls, *poll)
	}

	return polls, nil
}

func Vote(optionId int) error {
	_, err := db.DB.Exec(`UPDATE Option SET Votes = Votes + 1 WHERE Id = ?`, optionId)
	return err
}
