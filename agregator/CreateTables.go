package agregator

import "fmt"

func (agr *Agregator) CreateTableAgr() {

	sql := `

	-- Уничтожаем существующие таблицы таблицы		
	
	DROP TABLE IF EXISTS users			CASCADE;
	DROP TABLE IF EXISTS forum			CASCADE;
	DROP TABLE IF EXISTS thread			CASCADE;
	DROP TABLE IF EXISTS post				CASCADE;
	DROP TABLE IF EXISTS vote				CASCADE;	
	DROP TABLE IF EXISTS usersforum		CASCADE;



	-- Уничтожаем существующие тригеры.

	DROP TRIGGER IF EXISTS user_in_forum_thread_trigger		ON thread;
	DROP TRIGGER IF EXISTS add_thread_trigger				ON thread;
	DROP TRIGGER IF EXISTS add_post_trigger					ON post;



	-- Подключаем CITEXT.
	
	CREATE EXTENSION IF NOT EXISTS CITEXT;



	-- Создаем таблицы.

	-- Таблица пользователей.

	CREATE TABLE IF NOT EXISTS users (
		nickname		CITEXT							NOT NULL	PRIMARY KEY,
		fullname		VARCHAR,
		about			TEXT,
		email			CITEXT			UNIQUE );



	-- Таблица форумов.
	
	CREATE TABLE IF NOT EXISTS forum (
		slug			CITEXT			 				NOT NULL	PRIMARY KEY,
		title			VARCHAR							NOT NULL,
		nickname		CITEXT							NOT NULL 	REFERENCES users(nickname),
		posts			INTEGER							DEFAULT 0,
		threads			INTEGER							DEFAULT 0 );


--CREATE UNIQUE INDEX IF NOT EXISTS forum_unique on forum(slug);



	-- Таблица тредов.

	CREATE TABLE IF NOT EXISTS thread (
		id				BIGSERIAL									PRIMARY KEY,
		slug			CITEXT							NOT NULL,
		author			CITEXT							NOT NULL	REFERENCES users(nickname),
		forum			CITEXT							NOT NULL	REFERENCES forum(slug),
		created			TIMESTAMP WITH TIME ZONE		DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
		title			VARCHAR							NOT NULL,
		message			VARCHAR							NOT NULL,
		roots			INTEGER							NOT NULL	DEFAULT 0,
		votes			INTEGER							DEFAULT 0 );

	-- Индекс, который делает все не пустые slugs уникальными.

	CREATE UNIQUE INDEX IF NOT EXISTS thread_slug_unique ON thread(slug) WHERE slug <> '';

	--create unique index if not exists thread_unique on thread(slug, id);



	-- Таблица постов.

	CREATE TABLE IF NOT EXISTS post (
		id					BIGSERIAL								PRIMARY KEY,
		author				CITEXT						NOT NULL	REFERENCES users(nickname),
		forum				CITEXT						NOT NULL	REFERENCES forum(slug),
		thread_id			BIGSERIAL								REFERENCES thread(id),
		created				TIMESTAMP WITH TIME ZONE	DEFAULT (current_timestamp AT TIME ZONE 'UTC'),
		is_edited			BOOLEAN						DEFAULT FALSE,
		message				VARCHAR						NOT NULL,
		path				TEXT						NOT NULL	DEFAULT 0,
		parent				BIGINT						NOT NULL	DEFAULT 0,
		children			INTEGER						NOT NULL 	DEFAULT 0 );


create index if not exists "post_materialized_path_flat_sort" on post(thread_id, id);

create index if not exists "post_materialized_path_tree_sort" on post (thread_id, path, id);

create unique index if not exists "post_materialized_path_patent_tree_sort"
on post (thread_id, substring(path from 1 for 6) desc, substring(path from 8) asc, id);

create index if not exists "thread_created" on thread(forum, created);



	-- Таблица голосов.

	CREATE TABLE IF NOT EXISTS vote (
		id              BIGSERIAL									PRIMARY KEY,
		nickname        CITEXT							NOT NULL	REFERENCES users(nickname),
		thread_id		BIGINT							NOT NULL	REFERENCES thread(id),
		voice			INTEGER							DEFAULT 0,
		CONSTRAINT unique_vote UNIQUE(nickname, thread_id) );



	-- Таблица, хранящая информацию о полбзователях в форуме.

	CREATE TABLE IF NOT EXISTS usersforum (
		forum				CITEXT							NOT NULL	REFERENCES	forum(slug),
  		nickname			CITEXT		COLLATE ucs_basic 	NOT NULL 	REFERENCES users(nickname),
  		CONSTRAINT unique_users UNIQUE(forum, nickname) );



	-- Создаем функции.

	-- Функция, добавляющая пользователей, создавших тред.

	CREATE OR REPLACE FUNCTION "user_in_forum_on_create_thread"() RETURNS TRIGGER AS $$
		BEGIN
  			INSERT INTO usersforum(forum, nickname) VALUES(new.forum, new.author) ON CONFLICT DO NOTHING;
  			RETURN NULL;
		END
	$$ language plpgsql;



	-- Функция, увеличивающая кол-во тредов в таблице forum.

	CREATE OR REPLACE FUNCTION "counter_threads"() RETURNS TRIGGER AS $$
		BEGIN
  			UPDATE forum SET threads = threads + 1 WHERE slug = new.forum;
  			RETURN NULL;
		END
	$$ language plpgsql;



	-- Функция, увеличивающая кол-во постов в таблице forum.

	CREATE OR REPLACE FUNCTION "counter_posts"() RETURNS TRIGGER AS $$
		BEGIN
  			UPDATE forum SET posts = posts + 1 WHERE slug = NEW.forum;
  			RETURN NULL;
		END
	$$ language plpgsql;



	-- Создаем триггеры

	-- Создаем триггер, который срабатывает при каждом добавлении новых тредов.
	-- Натравливаем функцию "user_in_forum_on_create_thread"().

	CREATE TRIGGER  user_in_forum_thread_trigger
	AFTER INSERT ON thread
	FOR EACH ROW 
	EXECUTE PROCEDURE "user_in_forum_on_create_thread"();



	-- Создаем триггер, который срабатывает при каждом добавлении новых тредов.
	-- Натравливаем функцию "counter_threads"().

	CREATE TRIGGER "add_thread_trigger" 
	AFTER INSERT ON thread
	FOR EACH ROW
	EXECUTE PROCEDURE "counter_threads"();



	-- Создаем триггер, который срабатывает при каждом добавлении новых постов.
	-- Натравливаем функцию "counter_posts"().

	CREATE TRIGGER "add_post_trigger" 
	AFTER INSERT ON post
	FOR EACH ROW 
	EXECUTE PROCEDURE "counter_posts"();
`

	_, err := agr.Connection.Exec(sql)
	fmt.Println(err)
}
