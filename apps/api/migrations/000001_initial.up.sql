-- initial migration: schema_migrations tracking is managed automatically by go-migrate.
-- Add your first table definitions below.

CREATE TABLE IF NOT EXISTS `users` (
  `id`         BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `created_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` DATETIME        NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
