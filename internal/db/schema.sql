CREATE TABLE IF NOT EXISTS likes (
  id integer not null primary key autoincrement, 
  count integer DEFAULT 0
);

INSERT OR IGNORE INTO likes (id, count) VALUES (0, 0);