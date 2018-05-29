-- CREATE USER twodo_user;
-- CREATE DATABASE twodo_db;
-- GRANT ALL PRIVILEGES ON DATABASE twodo_db TO twodo_user;

CREATE TABLE IF NOT EXISTS todos (
  id SERIAL PRIMARY KEY,
  description VARCHAR(500),
  done BOOLEAN
);

CREATE INDEX IF NOT EXISTS index_todos ON todos(id);
