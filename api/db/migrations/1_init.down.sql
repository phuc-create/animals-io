-- Migration: Revert Initial Database Setup

-- Drop foreign key constraints
ALTER TABLE "group" DROP CONSTRAINT IF EXISTS "group_group_admin_id_fkey";
ALTER TABLE "post_comments" DROP CONSTRAINT IF EXISTS "post_comments_post_id_fkey";
ALTER TABLE "post_images" DROP CONSTRAINT IF EXISTS "post_images_post_id_fkey";
ALTER TABLE "follows" DROP CONSTRAINT IF EXISTS "follows_followed_user_id_fkey";
ALTER TABLE "follows" DROP CONSTRAINT IF EXISTS "follows_following_user_id_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "user_posts";
ALTER TABLE "group_members" DROP CONSTRAINT IF EXISTS "group_members_group_id_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_group_id_fkey";

-- Drop tables
DROP TABLE IF EXISTS "group_members";
DROP TABLE IF EXISTS "group";
DROP TABLE IF EXISTS "post_comments";
DROP TABLE IF EXISTS "post_images";
DROP TABLE IF EXISTS "posts";
DROP TABLE IF EXISTS "users";
DROP TABLE IF EXISTS "follows";
