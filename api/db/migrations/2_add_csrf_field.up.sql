ALTER TABLE users
ADD COLUMN csrf_token varchar DEFAULT '',
ADD COLUMN session_token varchar DEFAULT '';
