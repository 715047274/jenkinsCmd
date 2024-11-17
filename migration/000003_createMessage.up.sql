-- Create 'transfers' table

CREATE TABLE IF NOT EXISTS "messages" (
                                           "id" INTEGER PRIMARY KEY AUTOINCREMENT,
                                           "event_type" TEXT NOT NULL,
                                           "payload" TEXT NOT NULL,
                                           "status" TEXT NOT NULL,
                                           "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL,
                                           "updated_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP NOT NULL

);