CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  title VARCHAR,
  description VARCHAR,
  created_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS index_todos ON todos(id);
