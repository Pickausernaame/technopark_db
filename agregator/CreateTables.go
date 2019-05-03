package agregator

import "fmt"

func (agr *Agregator) CreateTableAgr() {

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
	nickname		CITEXT							NOT NULL 	REFERENCES users(nickname),
	posts			INTEGER							DEFAULT 0,
	threads			INTEGER							DEFAULT 0
);

CREATE TABLE IF NOT EXISTS thread (
	id				BIGSERIAL									PRIMARY KEY,
	slug			CITEXT							NOT NULL,
	author			CITEXT							NOT NULL	REFERENCES users(nickname),
	forum			CITEXT							NOT NULL	REFERENCES forum(slug),
	created			TIMESTAMP WITH TIME ZONE		DEFAULT NOW(),
	title			VARCHAR							NOT NULL,
	message			VARCHAR							NOT NULL,
	roots			INTEGER							NOT NULL	DEFAULT 0,
	votes			INTEGER							DEFAULT 0
);

create unique index if not exists thread_slug_unique on thread(slug) where slug <> '';

CREATE TABLE IF NOT EXISTS post (
	id					BIGSERIAL								PRIMARY KEY,
	author				CITEXT						NOT NULL	REFERENCES users(nickname),
	forum				CITEXT						NOT NULL	REFERENCES forum(slug),
	thread_id			BIGSERIAL								REFERENCES thread(id),
	created				TIMESTAMP WITH TIME ZONE	DEFAULT NOW(),
	isEdited			BOOLEAN						DEFAULT FALSE,
	message				VARCHAR						NOT NULL,
	path				TEXT						NOT NULL DEFAULT 0,
	children			INTEGER						NOT NULL  DEFAULT 0	
);

CREATE TABLE IF NOT EXISTS vote (
	id              BIGSERIAL									PRIMARY KEY,
	nickname        CITEXT							NOT NULL	REFERENCES users(nickname),
	thread          BIGINT							NOT NULL	REFERENCES thread(id),
	voice			INTEGER							DEFAULT 0,
	CONSTRAINT unique_vote UNIQUE(nickname, thread)
);

`
	//parent				BIGINT						NULL		REFERENCES post(id)
	_, err := agr.Connection.Exec(sql)
	fmt.Println(err)
	// 	thread				CITEXT						NOT NULL	REFERENCES thread(slug),
}
