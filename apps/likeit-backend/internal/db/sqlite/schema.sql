CREATE TABLE IF NOT EXISTS likes (
  pk integer not null primary key autoincrement,
  count integer DEFAULT 0
);

CREATE TABLE IF NOT EXISTS users (
  pk integer not null primary key autoincrement,
  id text not null unique
);

INSERT OR IGNORE INTO likes (pk, count) VALUES (0, 0);