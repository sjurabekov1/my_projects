CREATE TYPE "tax_type" AS ENUM ('Fixed', 'Percentage');

CREATE TYPE "transaction_status" AS ENUM ('pending', 'completed', 'failed', 'canceled');

-- serial's max - 2 147 483 648
CREATE TABLE IF NOT EXISTS "bank_account" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" serial NOT NULL,
    "bank_account_name" VARCHAR,
    "bic_code" VARCHAR,
    "iban_code" VARCHAR,
    "swift_code" VARCHAR,
    "region" VARCHAR,
    "registration_number" bigint,
    "address" VARCHAR,
    "bank_address" VARCHAR,
    "phone_number" VARCHAR,
    "payment_reference" VARCHAR,
    "description" VARCHAR,
    "salesforce_id" VARCHAR UNIQUE,
    "owner_name" VARCHAR,
    "transactions_reference" VARCHAR,
    "receiving_party_bank_details" VARCHAR,
    "currency" INTEGER NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "currencies" (
    "id" serial PRIMARY KEY NOT NULL,
    "name" VARCHAR UNIQUE NOT NULL,
    "icon" VARCHAR,
    "active" BOOLEAN
);

CREATE TABLE IF NOT EXISTS "bank_account_currency" (
    "id" serial PRIMARY KEY NOT NULL,
    "bank_account_id" VARCHAR,
    "currency_id" INTEGER,
    "created_at" TIME
);

CREATE TABLE IF NOT EXISTS "company" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" serial NOT NULL,
    "company_name" VARCHAR,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "currency_convertor" (
    "id" serial PRIMARY KEY NOT NULL,
    "primary_currency" INTEGER NOT NULL,
    "secondary_currency" INTEGER NOT NULL,
    "amount" double precision,
    CONSTRAINT fk_primary_currency_id FOREIGN KEY ("primary_currency") REFERENCES "currencies" ("id"),
    CONSTRAINT fk_secondary_currency_id FOREIGN KEY ("secondary_currency") REFERENCES "currencies" ("id"),
    CONSTRAINT uq_primary_currency_secondary_currency UNIQUE ("primary_currency", "secondary_currency")
);

CREATE TABLE IF NOT EXISTS "tax" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" serial NOT NULL,
    "type" tax_type,
    "value" numeric,
    "operation_id" VARCHAR,
    "company_id" uuid,
    "currency" VARCHAR,
    "min_fee" numeric,
    "broker_id" uuid,
    "partner_id" uuid,
    "discount" numeric,
    "broker_type" tax_type,
    "broker_commission" numeric,
    "partner_commission" numeric,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "comission_group_tax_rule" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" serial NOT NULL,
    "comission_group_id" uuid,
    "tax_id" uuid
);

CREATE TABLE IF NOT EXISTS "comission_group" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" serial NOT NULL,
    "name" VARCHAR
);

CREATE TABLE IF NOT EXISTS "broker" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" serial NOT NULL,
    "name" VARCHAR,
    "legal_addrress" VARCHAR,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "partner" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" serial NOT NULL,
    "name" VARCHAR,
    "user_type" VARCHAR,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "main_graph" (
    "investor_id" VARCHAR,
    "taum" VARCHAR,
    "pocket_id" uuid,
    "symbol_id" VARCHAR
);

CREATE TABLE IF NOT EXISTS "pocket" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "o_id" bigserial NOT NULL,
    "name" VARCHAR,
    "user_id" VARCHAR,
    "company_id" uuid,
    "type" INTEGER,
    "created_at" TIME,
    "title" text,
    "icon" text,
    "salesforce_id" VARCHAR,
    "color" VARCHAR
);

CREATE TABLE IF NOT EXISTS "pocket_type" (
    "id" SERIAL PRIMARY KEY,
    "name" VARCHAR,
    "icon_id" VARCHAR,
    "required" BOOLEAN
);

CREATE TABLE IF NOT EXISTS "operation" (
    "id" VARCHAR PRIMARY KEY,
    "operation_name" VARCHAR
);

CREATE TABLE IF NOT EXISTS "order" (
    "id" uuid PRIMARY KEY,
    "o_id" bigserial NOT NULL,
    "investor_id" VARCHAR,
    "symbol_id" VARCHAR,
    "place_time" TIMESTAMP,
    "current_modification_id" uuid,
    "exante_account_id" VARCHAR,
    "username" VARCHAR,
    "status" VARCHAR,
    "last_update" VARCHAR,
    "reason" VARCHAR,
    "limit_price" VARCHAR,
    "quantity" VARCHAR,
    "side" VARCHAR,
    "salesforce_order_id" VARCHAR,
    "pocket_id" uuid,
    "currency" VARCHAR,
    "future_pocket_id" uuid,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP DEFAULT NULL
);

CREATE TABLE IF NOT EXISTS "transaction" (
    "id" uuid PRIMARY KEY NOT NULL DEFAULT uuid_generate_v4(),
    "parent_id" uuid,
    "o_id" bigserial NOT NULL,
    "comment" VARCHAR,
    "user_id" VARCHAR,
    "pocket_id" uuid,
    "operation_id" VARCHAR,
    "amount" numeric,
    "currency" VARCHAR,
    "status" transaction_status,
    "order_id" uuid,
    "bank_account" VARCHAR,
    "salesforce_id" VARCHAR,
    "future_pocket_id" uuid,
    "created_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "updated_at" TIMESTAMP NOT NULL DEFAULT NOW(),
    "deleted_at" TIMESTAMP DEFAULT NULL
);

ALTER TABLE "bank_account"
ADD FOREIGN KEY ("currency") REFERENCES "currencies" ("id");

ALTER TABLE "bank_account_currency"
ADD FOREIGN KEY ("bank_account_id") REFERENCES "bank_account" ("salesforce_id");

ALTER TABLE "bank_account_currency"
ADD FOREIGN KEY ("currency_id") REFERENCES "currencies" ("id");

ALTER TABLE "tax"
ADD FOREIGN KEY ("operation_id") REFERENCES "operation" ("id");

ALTER TABLE "tax"
ADD FOREIGN KEY ("broker_id") REFERENCES "broker" ("id");

ALTER TABLE "tax"
ADD FOREIGN KEY ("partner_id") REFERENCES "partner" ("id");

ALTER TABLE "comission_group_tax_rule"
ADD FOREIGN KEY ("tax_id") REFERENCES "tax" ("id");

ALTER TABLE "main_graph"
ADD FOREIGN KEY ("pocket_id") REFERENCES "pocket" ("id");

ALTER TABLE "pocket"
ADD FOREIGN KEY ("company_id") REFERENCES "company" ("id");

ALTER TABLE "pocket"
ADD FOREIGN KEY ("type") REFERENCES "pocket_type" ("id");

ALTER TABLE "order"
ADD FOREIGN KEY ("pocket_id") REFERENCES "pocket" ("id");

ALTER TABLE "order"
ADD FOREIGN KEY ("currency") REFERENCES "currencies" ("name");

ALTER TABLE "order"
ADD FOREIGN KEY ("future_pocket_id") REFERENCES "pocket" ("id");

ALTER TABLE "transaction"
ADD FOREIGN KEY ("parent_id") REFERENCES "transaction" ("id");

ALTER TABLE "transaction"
ADD FOREIGN KEY ("pocket_id") REFERENCES "pocket" ("id");

ALTER TABLE "transaction"
ADD FOREIGN KEY ("operation_id") REFERENCES "operation" ("id");

ALTER TABLE "transaction"
ADD FOREIGN KEY ("currency") REFERENCES "currencies" ("name");

ALTER TABLE "transaction"
ADD FOREIGN KEY ("order_id") REFERENCES "order" ("id");

ALTER TABLE "transaction"
ADD FOREIGN KEY ("future_pocket_id") REFERENCES "pocket" ("id");

ALTER TABLE "comission_group_tax_rule"
ADD FOREIGN KEY ("comission_group_id") REFERENCES "comission_group" ("id");