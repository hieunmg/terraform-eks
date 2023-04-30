DROP TABLE IF EXISTS "public"."sessions";
CREATE TABLE "public"."sessions" (
  "id" uuid PRIMARY KEY,
  "account_id" int NOT NULL,
  "refresh_token" varchar(100) NOT NULL,
  -- "user_agent" varchar(100) NOT NULL,
  -- "client_ip" varchar(100) NOT NULL,
  "is_blocked" boolean NOT NULL DEFAULT false,
  "expired_at" timestamptz NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

DROP TABLE IF EXISTS "public"."accounts";
CREATE TABLE "public"."accounts" (
  "id" SERIAL PRIMARY KEY,
  "username" varchar(100) NOT NULL,
  "full_name" varchar(100) NOT NULL,
  "salt" varchar(100),
  "password" varchar(100) NOT NULL,
  "status" int2 NOT NULL DEFAULT 1,
  "avatar" json NULL,
  "birthday" timestamptz,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  "updated_at" timestamptz NOT NULL DEFAULT (now())
);

INSERT INTO "accounts" ("username", "password", "full_name", "salt") 
VALUES ('hieuaws', '$2a$10$AgUakVmIxV2RT6.28v4saOZORAIwOIY7qnYwR4GGtuU/t8rEC5cPK
', 'Hieu Minh', 'caELHU');


ALTER TABLE "sessions" ADD FOREIGN KEY ("account_id") REFERENCES "accounts" ("id");
