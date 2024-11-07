CREATE TABLE IF NOT EXISTS "payment_details" (
    "id" bigserial PRIMARY KEY,
    "account_name" varchar NOT NULL,
    "account_number" varchar NOT NULL,
    "bank_name" varchar NOT NULL,
    "company_id" bigint NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "payment_details" ADD FOREIGN KEY ("company_id") REFERENCES "company_details" ("id");