CREATE TABLE "follows" (
  "id" INTEGER PRIMARY KEY,
  "following_user_id" INTEGER,
  "followed_user_id" INTEGER,
  "created_at" timestamp
);

CREATE TABLE "users" (
  "id" INTEGER PRIMARY KEY,
  "username" varchar,
  "password" varchar,
  "role" varchar,
  "created_at" timestamp
);

CREATE TABLE "posts" (
  "id" INTEGER PRIMARY KEY,
  "title" varchar,
  "body" text,
  "user_id" INTEGER NOT NULL,
  "status" varchar,
  "like" INTEGER DEFAULT 0,
  "group_id" INTEGER,
  "created_at" timestamp
);

CREATE TABLE "post_images" (
  "id" INTEGER PRIMARY KEY,
  "post_id" INTEGER NOT NULL,
  "created_at" INTEGER NOT NULL,
  "uploaded_at" INTEGER,
  "file_size_bytes" INTEGER,
  "is_file_uploaded" INTEGER NOT NULL,
  "width" INTEGER NOT NULL,
  "height" INTEGER NOT NULL,
  "is_deleted" INTEGER NOT NULL
);

CREATE TABLE "post_comments" (
  "id" INTEGER PRIMARY KEY,
  "post_id" INTEGER NOT NULL,
  "comment" text,
  "created_at" timestamp
);

CREATE TABLE "group" (
  "id" INTEGER PRIMARY KEY,
  "group_name" varchar NOT NULL,
  "group_admin_id" INTEGER NOT NULL,
  "group_description" varchar,
  "total_members" INTEGER DEFAULT 1,
  "is_public" bool DEFAULT true,
  "is_active" bool DEFAULT true,
  "created_at" timestamp
);

CREATE TABLE "group_members" (
  "group_id" INTEGER NOT NULL,
  "group_member_id" INTEGER NOT NULL
);

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("group_id") REFERENCES "group" ("id");

ALTER TABLE "group_members" ADD FOREIGN KEY ("group_id") REFERENCES "group" ("id");

ALTER TABLE "posts" ADD CONSTRAINT "user_posts" FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("following_user_id") REFERENCES "users" ("id");

ALTER TABLE "follows" ADD FOREIGN KEY ("followed_user_id") REFERENCES "users" ("id");

ALTER TABLE "post_images" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "post_comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "group" ADD FOREIGN KEY ("group_admin_id") REFERENCES "users" ("id");
