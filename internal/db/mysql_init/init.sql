CREATE SCHEMA IF NOT EXISTS `payment_service` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
CREATE USER IF NOT EXISTS 'payment_service' @'%' IDENTIFIED WITH caching_sha2_password BY 'payment_service';
REVOKE ALL PRIVILEGES,
GRANT OPTION
FROM 'payment_service' @'%';
GRANT INSERT,
  SELECT,
  UPDATE,
  DELETE ON `payment_service`.* TO 'payment_service' @'%';
FLUSH PRIVILEGES;