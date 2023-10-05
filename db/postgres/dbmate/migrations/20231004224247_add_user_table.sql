-- migrate:up
CREATE SCHEMA IF NOT EXISTS "game";

CREATE TABLE IF NOT EXISTS "game"."users" (
  "ksuid" varchar(255) PRIMARY KEY NOT NULL,
  "email" varchar(255),
  "password" varchar(255),
  "inserted_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz NOT NULL
);

-- migrate:down
DROP TABLE IF EXISTS "game"."users";

