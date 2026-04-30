-- Add Firebase UID + basic account status fields
ALTER TABLE `users`
  ADD COLUMN `firebase_uid` VARCHAR(128) NOT NULL AFTER `id`,
  ADD COLUMN `disabled_at`  DATETIME NULL AFTER `updated_at`,
  ADD COLUMN `last_seen_at` DATETIME NULL AFTER `disabled_at`;

-- Enforce uniqueness so a Firebase UID maps to exactly one local user row
ALTER TABLE `users`
  ADD UNIQUE KEY `ux_users_firebase_uid` (`firebase_uid`);

-- Optional: helpful index for moderation / cleanup queries
ALTER TABLE `users`
  ADD KEY `ix_users_disabled_at` (`disabled_at`);