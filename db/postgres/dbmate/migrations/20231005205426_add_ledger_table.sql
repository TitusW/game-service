-- migrate:up
CREATE TABLE IF NOT EXISTS "game"."ledgers" (
  "ksuid" text PRIMARY KEY NOT NULL,
  "category" varchar NOT NULL,
  "amount" numeric NOT NULL,
  "inserted_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

-- migrate:down
DROP TABLE IF EXISTS "game"."ledgers";
