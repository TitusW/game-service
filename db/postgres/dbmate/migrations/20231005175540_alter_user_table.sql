-- migrate:up
ALTER TABLE "game"."users"
ADD CONSTRAINT "unique_email" UNIQUE ("email")

-- migrate:down
ALTER TABLE "game"."users" 
DROP CONSTRAINT "unique_email" UNIQUE ("email")
