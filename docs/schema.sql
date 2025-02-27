-- SQL dump generated using DBML (dbml.dbdiagram.io)
-- Database: PostgreSQL
-- Generated at: 2025-02-27T02:52:31.551Z

CREATE TYPE "gender" AS ENUM (
  'nam',
  'nữ',
  'khác'
);

CREATE TABLE "users" (
  "id" serial PRIMARY KEY,
  "username" varchar UNIQUE NOT NULL,
  "hashed_password" varchar NOT NULL,
  "full_name" varchar NOT NULL,
  "is_verify" boolean NOT NULL DEFAULT false,
  "email" varchar,
  "gender" gender,
  "licence" varchar,
  "dob" timestamp,
  "active" bool NOT NULL DEFAULT true,
  "avatar" varchar(255),
  "updated_at" timestamptz,
  "password_changed_at" timestamptz NOT NULL DEFAULT '0001-01-01 00:00:00Z',
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "sessions" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "refresh_token" varchar NOT NULL,
  "user_agent" varchar NOT NULL,
  "client_ip" varchar NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expires_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "verifies" (
  "id" serial PRIMARY KEY,
  "username" varchar NOT NULL,
  "email" varchar NOT NULL,
  "secret_code" varchar NOT NULL,
  "is_used" bool NOT NULL DEFAULT false,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "expired_at" timestamptz NOT NULL DEFAULT ((now() + interval '15 minutes'))
);

ALTER TABLE "sessions" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;

ALTER TABLE "verifies" ADD FOREIGN KEY ("username") REFERENCES "users" ("username") ON DELETE CASCADE;
