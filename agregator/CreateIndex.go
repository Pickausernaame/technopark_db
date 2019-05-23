package agregator

func (agr *Agregator) CreateIndexes() {
	sql := `

DROP INDEX IF EXIST flat_sort;
DROP INDEX IF EXIST tree_sort;
DROP INDEX IF EXIST parent_tree_sort;
DROP INDEX IF EXIST thread_created;
DROP INDEX IF EXIST thread_forum_index;



CREATE INDEX IF NOT EXIST flat_sort ON post(thread_id);

CREATE INDEX IF NOT EXIST tree_sort ON post(thread_id, path);

CREATE INDEX IF NOT EXIST parent_tree_sort ON post (thread_id, substring(path from 1 for 6) desc);

CREATE INDEX IF NOT EXIST thread_created ON thread(forum, created);

CREATE INDEX IF NOT EXIST thread_forum_index ON thread (forum);`

	agr.Connection.Exec(sql)
}
