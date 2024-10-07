-- Create 'accounts' table
CREATE TABLE IF NOT EXISTS "accounts" (
                                          "id" INTEGER PRIMARY KEY AUTOINCREMENT,
                                          "owner" TEXT NOT NULL,
                                          "balance" INTEGER NOT NULL,
                                          "currency" TEXT NOT NULL,
                                          "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL
);

-- Create 'entries' table
CREATE TABLE IF NOT EXISTS "entries" (
                                         "id" INTEGER PRIMARY KEY AUTOINCREMENT,
                                         "account_id" INTEGER NOT NULL,
                                         "amount" INTEGER NOT NULL,
                                         "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                         FOREIGN KEY ("account_id") REFERENCES "accounts" ("id")
    );

-- Create 'transfers' table
CREATE TABLE IF NOT EXISTS "transfers" (
                                           "id" INTEGER PRIMARY KEY AUTOINCREMENT,
                                           "from_account_id" INTEGER NOT NULL,
                                           "to_account_id" INTEGER NOT NULL,
                                           "amount" INTEGER NOT NULL,
                                           "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                           FOREIGN KEY ("from_account_id") REFERENCES "accounts" ("id"),
    FOREIGN KEY ("to_account_id") REFERENCES "accounts" ("id")
    );

-- Indexes
CREATE INDEX IF NOT EXISTS "index_accounts_owner" ON "accounts" ("owner");
CREATE INDEX IF NOT EXISTS "index_entries_account_id" ON "entries" ("account_id");
CREATE INDEX IF NOT EXISTS "index_transfers_from_account_id" ON "transfers" ("from_account_id");
CREATE INDEX IF NOT EXISTS "index_transfers_to_account_id" ON "transfers" ("to_account_id");
CREATE INDEX IF NOT EXISTS "index_transfers_from_to_account_id" ON "transfers" ("from_account_id", "to_account_id");

-- Comments
-- PRAGMA foreign_keys = ON;  -- Enable foreign key support

-- Comments on columns
-- COMMENT ON COLUMN "entries"."amount" IS 'can be negative or positive';
-- COMMENT ON COLUMN "transfers"."amount" IS 'must be positive';

-- INSERT INTO accounts (owner, balance, currency) VALUES ('test', 10, 'USD')