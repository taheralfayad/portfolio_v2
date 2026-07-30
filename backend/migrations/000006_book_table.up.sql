CREATE TYPE book_status AS ENUM ('not_yet_read', 'reading', 'complete', 'currently_reading');

CREATE TABLE book (
	id SERIAL PRIMARY KEY,
	title VARCHAR(255) NOT NULL,
	author VARCHAR(255) NOT NULL,
	last_read_time VARCHAR(255),
	percentage_finished NUMERIC(3, 2),
	status book_status NOT NULL DEFAULT 'not_yet_read',
	UNIQUE (title, author)
);

CREATE TYPE image_site AS ENUM ('home', 'books');

ALTER TABLE images ADD site image_site DEFAULT 'home';
