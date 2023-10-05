-- migrate:up
ALTER TABLE "game"."wallets" 
ADD COLUMN IF NOT EXISTS "user_ksuid" varchar NOT NULL
REFERENCES "game"."users" ("ksuid");

-- migrate:down
ALTER TABLE "game"."wallets" 
DROP COLUMN IF EXISTS "user_ksuid";
