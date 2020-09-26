--

DROP TABLE IF EXISTS reviews;

DROP TABLE IF EXISTS reviewers;

DROP TABLE IF EXISTS book_authors;

DROP TABLE IF EXISTS books;

DROP TYPE IF EXISTS category;

DROP TABLE IF EXISTS authors;

DROP TABLE IF EXISTS publishers;

--

CREATE TABLE publishers (
	id SERIAL PRIMARY KEY,
	name VARCHAR(80) NOT NULL
);

CREATE TABLE authors (
  id SERIAL PRIMARY KEY,

  name VARCHAR(80) NOT NULL,
  surname VARCHAR(80) NOT NULL
);

ALTER TABLE authors ADD CONSTRAINT unique_author_name UNIQUE (name, surname);

CREATE TYPE category AS ENUM('poetry', 'drama', 'prose', 'non-fiction', 'media');

CREATE TABLE books (
  id SERIAL PRIMARY KEY,
  title VARCHAR(80) NOT NULL,
	isbn VARCHAR(80) NOT NULL unique,
  publisher_id BIGINT NOT NULL REFERENCES publishers (id), --' book has one publisher
	category category NOT NULL, --' book has one category
  created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE book_authors (
	book_id BIGINT REFERENCES books(id), --' book belongs to many authors
	author_id BIGINT REFERENCES authors (id) --' author has many books
);

ALTER TABLE book_authors ADD CONSTRAINT unique_book_authors UNIQUE (book_id, author_id);

CREATE TABLE reviewers (
	id SERIAL PRIMARY KEY,
	name VARCHAR(80) NOT NULL,
	email VARCHAR(80) NOT NULL
);

--' reviewer has many reviews
CREATE TABLE reviews (
	id SERIAL PRIMARY KEY,
	book_id BIGINT NOT NULL REFERENCES books (id), --' book has many reviews
	reviewer_id BIGINT NOT NULL REFERENCES reviewers (id), --' review belongs to reviewer
	content TEXT NOT NULL,
	updated_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP,
	created_at TIMESTAMP(0) WITHOUT TIME ZONE NOT NULL DEFAULT CURRENT_TIMESTAMP
);
