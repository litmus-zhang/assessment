CREATE TABLE IF NOT EXISTS "customers" (
    "id" bigserial PRIMARY KEY,
    "first_name" varchar NOT NULL,
    "last_name" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "phone_number" varchar NOT NULL,
    "address" varchar,
    "company_id" bigint NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "customers" ADD FOREIGN KEY ("company_id") REFERENCES "company_details" ("id");