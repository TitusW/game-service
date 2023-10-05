-- migrate:up
ALTER TABLE "game"."users"
ALTER COLUMN "deleted_at" DROP NOT NULL;

-- migrate:down
ALTER TABLE "game"."users"
ALTER COLUMN "deleted_at" SET NOT NULL;
