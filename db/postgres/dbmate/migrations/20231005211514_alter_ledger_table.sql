-- migrate:up
ALTER TABLE "game"."ledgers" 
ADD COLUMN IF NOT EXISTS "wallet_ksuid" varchar NOT NULL
REFERENCES "game"."wallets" ("ksuid");

-- migrate:down
ALTER TABLE "game"."ledgers" 
DROP COLUMN IF EXISTS "wallet_ksuid";
