CREATE TABLE IF NOT EXISTS users (
	id BIGSERIAL NOT NULL PRIMARY KEY,
	username VARCHAR(50) NOT NULL,
	password VARCHAR(100) NOT NULL,
	token VARCHAR(100) NOT NULL
);