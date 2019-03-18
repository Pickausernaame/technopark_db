package agregator

import "fmt"

func (agr *Agregator) CreateTableAgr() {
	//TODO таблица с постами

	sql := `
		
	DROP TABLE IF EXISTS users CASCADE;
	DROP TABLE IF EXISTS forum CASCADE;
	DROP TABLE IF EXISTS thread CASCADE;
	DROP TABLE IF EXISTS post CASCADE;
	DROP TABLE IF EXISTS vote CASCADE;	


	CREATE EXTENSION IF NOT EXISTS CITEXT;

	CREATE TABLE IF NOT EXISTS users (
	nickname		CITEXT							NOT NULL	PRIMARY KEY,
	fullname		VARCHAR,
	about			TEXT,
	email			CITEXT			UNIQUE
);

CREATE TABLE IF NOT EXISTS forum (
	slug			CITEXT			 				NOT NULL	PRIMARY KEY,
	title			VARCHAR							NOT NULL,
	owner_nickname	CITEXT							NOT NULL 	REFERENCES users(nickname),
	posts			INTEGER							DEFAULT 0,
	threads			INTEGER							DEFAULT 0
);

CREATE TABLE IF NOT EXISTS thread (
	id				BIGSERIAL									PRIMARY KEY,
	slug			CITEXT			UNIQUE,
	author_nickname	CITEXT							NOT NULL	REFERENCES users(nickname),
	forum_slug		CITEXT							NOT NULL	REFERENCES forum(slug),
	created			TIMESTAMP WITH TIME ZONE		DEFAULT NOW(),
	title			VARCHAR							NOT NULL,
	message			VARCHAR							NOT NULL,
	votes			INTEGER							DEFAULT 0
);

CREATE TABLE IF NOT EXISTS post (
	id					BIGSERIAL								PRIMARY KEY,
	author_nickname		CITEXT						NOT NULL	REFERENCES users(nickname),
	forum_slug			CITEXT						NOT NULL	REFERENCES forum(slug),
	thread_slug			CITEXT									REFERENCES thread(slug),
	thread_id			BIGSERIAL								REFERENCES thread(id),
	created				TIMESTAMP WITH TIME ZONE	DEFAULT NOW(),
	isEdited			BOOLEAN						DEFAULT FALSE,
	message				VARCHAR						NOT NULL,
	parent				BIGINT						NULL		REFERENCES post(id)
);

CREATE TABLE IF NOT EXISTS vote (
	id              BIGSERIAL									PRIMARY KEY,
	nickname        CITEXT							NOT NULL	REFERENCES users(nickname),
	thread          BIGINT							NOT NULL	REFERENCES thread(id),
	voice			INTEGER							DEFAULT 0,
	CONSTRAINT unique_vote UNIQUE(nickname, thread)
);

`

	_, err := agr.Connection.Exec(sql)
	fmt.Println(err)

}
