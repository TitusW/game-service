-- migrate:up
CREATE TABLE IF NOT EXISTS "game"."user_bank_accounts" (
  "ksuid" varchar(255) PRIMARY KEY NOT NULL,
  "user_ksuid" varchar(255) NOT NULL,
  "bank_account_name" varchar(255) NOT NULL,
  "bank_name" varchar(255) NOT NULL,
  "account_number" varchar(255) NOT NULL
);

ALTER TABLE "game"."user_bank_accounts" ADD FOREIGN KEY ("user_ksuid") REFERENCES "game"."users" ("ksuid");

-- migrate:down
DROP TABLE IF EXISTS "game"."user_bank_accounts"
