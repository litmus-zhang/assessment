CREATE TABLE IF NOT EXISTS "company_details" (
    "id" bigserial PRIMARY KEY,
    "name" varchar NOT NULL,
    "address" varchar NOT NULL,
    "phone_number" varchar NOT NULL,
    "email" varchar UNIQUE NOT NULL,
    "owned_by" int NOT NULL,
    "created_at"  timestamptz NOT NULL DEFAULT (now())
);

ALTER TABLE "company_details" ADD FOREIGN KEY ("owned_by") REFERENCES "users" ("id");