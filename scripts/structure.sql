CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  title VARCHAR(100),
  description VARCHAR(500),
  created_at TIMESTAMP
);

CREATE INDEX IF NOT EXISTS index_todos ON todos(id);
