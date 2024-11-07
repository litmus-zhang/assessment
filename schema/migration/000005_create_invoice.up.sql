CREATE TABLE IF NOT EXISTS "invoices" (
    "id" bigserial PRIMARY KEY,
    "customer_id" bigint NOT NULL,
    "name" varchar NOT NULL,
    "amount" DECIMAL(10, 2) NOT NULL DEFAULT 0,
    "due_date" timestamptz NOT NULL,
    "status" varchar NOT NULL DEFAULT 'PENDING',
    "company_id" bigint NOT NULL,
    "note" varchar,
    "discount" DECIMAL(10, 2) NOT NULL DEFAULT 0,
    "updated_at" timestamptz NOT NULL DEFAULT (now()),
    "created_at"  timestamptz NOT NULL DEFAULT (now()),
    FOREIGN KEY ("customer_id") REFERENCES "customers" (id),
    FOREIGN KEY ("company_id") REFERENCES "company_details" (id)
);

CREATE TABLE IF NOT EXISTS "items" (
    "id" bigserial PRIMARY KEY,
    "invoice_id" bigint NOT NULL,
    "name" varchar NOT NULL,
    "description" varchar NOT NULL,
    "quantity" int NOT NULL,
    "unit_price" DECIMAL(10, 2) NOT NULL,
    "total_price" DECIMAL(10, 2) GENERATED ALWAYS AS (quantity * unit_price) STORED,
    FOREIGN KEY ("invoice_id") REFERENCES "invoices" (id) ON DELETE CASCADE
);