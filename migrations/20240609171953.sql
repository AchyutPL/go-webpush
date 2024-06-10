-- Create "subscriptions" table
CREATE TABLE "public"."subscriptions" (
  "id" bigserial NOT NULL,
  "created_at" timestamptz NULL,
  "updated_at" timestamptz NULL,
  "deleted_at" timestamptz NULL,
  "endpoint" text NULL,
  "expiration_time" bigint NULL,
  "keys" bytea NULL,
  "test" text NULL,
  PRIMARY KEY ("id")
);
-- Create index "idx_subscriptions_deleted_at" to table: "subscriptions"
CREATE INDEX "idx_subscriptions_deleted_at" ON "public"."subscriptions" ("deleted_at");
