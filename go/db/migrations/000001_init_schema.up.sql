CREATE TABLE "users" (
  "id" varchar PRIMARY KEY,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "events" (
  "name" varchar PRIMARY KEY,
  "start_date" date NOT NULL,
  "end_date" date NOT NULL,
  "link" varchar NOT NULL,
  "description" text NOT NULL
);