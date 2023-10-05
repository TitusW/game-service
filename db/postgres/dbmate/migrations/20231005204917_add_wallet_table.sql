-- migrate:up
CREATE TABLE IF NOT EXISTS "game"."wallets" (
  "ksuid" text PRIMARY KEY NOT NULL,
  "current_amount" numeric,
  "inserted_at" timestamptz NOT NULL,
  "updated_at" timestamptz NOT NULL,
  "deleted_at" timestamptz
);

-- migrate:down
DROP TABLE IF EXISTS "game"."wallets";

