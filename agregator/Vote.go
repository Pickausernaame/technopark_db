package agregator

import (
	"github.com/Pickausernaame/technopark_db/models"
)

func (agr *Agregator) GetThreadVotesBySlug(slug string) (id int, err error) {
	sql := `
	SELECT id FROM thread
		WHERE slug = $1;`
	err = agr.Connection.QueryRow(sql, slug).Scan(&id)
	return
}

func (agr *Agregator) GetVote(Nickname string, id int) (vote models.Vote, err error) {
	var voice int
	sql := `
	SELECT voice FROM vote
		WHERE nickname = $1 AND thread_id = $2;`
	err = agr.Connection.QueryRow(sql, Nickname, id).Scan(&voice)
	vote.Nickname = Nickname
	vote.Voice = voice
	vote.Id = id
	return
}

func (agr *Agregator) UpdateVote(vote models.Vote) (err error) {
	sql := `
	UPDATE vote SET voice = $3
		WHERE nickname = $1 AND thread_id = $2;`
	_, err = agr.Connection.Exec(sql, vote.Nickname, vote.Id, vote.Voice)
	return
}

func (agr *Agregator) InsertVote(vote models.Vote) (err error) {
	sql := `
	INSERT INTO vote (nickname, thread_id, voice)
		VALUES($1, $2, $3);`
	_, err = agr.Connection.Exec(sql, vote.Nickname, vote.Id, vote.Voice)
	return
}

func (agr *Agregator) UpdateThreadVote(vote int, id int) (thread models.Thread, err error) {
	sql := `
	UPDATE thread SET votes = votes + $1
		WHERE id = $2
	RETURNING author, created, forum, id, message, slug, title, votes;`
	err = agr.Connection.QueryRow(sql, vote, id).Scan(&thread.Author, &thread.Created, &thread.Forum, &thread.Id, &thread.Message, &thread.Slug, &thread.Title, &thread.Votes)
	return
}
