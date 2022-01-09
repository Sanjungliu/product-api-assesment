CREATE TABLE product(
	id serial PRIMARY KEY,
	name varchar,
	price bigint,
	description varchar,
	quantity int,
	created_at timestamptz NOT NULL DEFAULT (NOW())
);