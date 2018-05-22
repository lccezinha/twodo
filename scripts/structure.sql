CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  description VARCHAR(500),
  done BOOLEAN
);

CREATE INDEX IF NOT EXISTS index_todos ON todos(id);
