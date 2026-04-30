-- Drop indexes first (required before dropping the column they reference)
ALTER TABLE `users`
  DROP INDEX `ux_users_firebase_uid`,
  DROP INDEX `ix_users_disabled_at`;

-- Then drop the columns
ALTER TABLE `users`
  DROP COLUMN `firebase_uid`,
  DROP COLUMN `disabled_at`,
  DROP COLUMN `last_seen_at`;