CREATE TABLE "author" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL
);

CREATE TABLE "book" (
    "id" UUID PRIMARY KEY,
    "name" VARCHAR NOT NULL,
    "price" NUMERIC NOT NULL,
    "author_id" uuid not null REFERENCES author(id)
);


