-- reverse: drop "payments" table
CREATE TABLE `payments` (
  `id` bigint unsigned NOT NULL AUTO_INCREMENT,
  `package_id` bigint unsigned NOT NULL COMMENT "Package ID",
  `user_id` bigint unsigned NOT NULL COMMENT "User ID",
  `status` enum('PENDING','COMPLETED','CANCELLED','FAILED') NOT NULL COMMENT "Order status",
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  INDEX `user_id` (`user_id`),
  CONSTRAINT `payments_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION
) CHARSET utf8mb4 COLLATE utf8mb4_0900_ai_ci;
-- reverse: modify "users" table
ALTER TABLE `users` DROP COLUMN `role`, DROP COLUMN `stripe_customer_id`;
