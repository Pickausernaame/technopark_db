package agregator

import "fmt"

func (agr *Agregator) CreateIndexes() {
	sql := `

DROP INDEX IF EXISTS flat_sort;
DROP INDEX IF EXISTS tree_sort;
DROP INDEX IF EXISTS parent_tree_sort;
DROP INDEX IF EXISTS thread_created;
DROP INDEX IF EXISTS thread_forum_index;



CREATE INDEX IF NOT EXISTS flat_sort ON post(thread_id, id);

CREATE INDEX IF NOT EXISTS tree_sort ON post(thread_id, path, id);

CREATE INDEX IF NOT EXISTS parent_tree_sort ON post (thread_id, substring(path from 1 for 6) desc, id);

CREATE INDEX IF NOT EXISTS thread_created ON thread(forum, created);

CREATE INDEX IF NOT EXISTS thread_forum_index ON thread (forum);

--CREATE INDEX IF NOT EXISTS thread_vote ON thread (id, vote);


`

	_, err := agr.Connection.Exec(sql)
	fmt.Println(err)
}
