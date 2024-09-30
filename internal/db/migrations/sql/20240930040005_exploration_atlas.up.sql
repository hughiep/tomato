-- modify "users" table
ALTER TABLE `users` ADD COLUMN `stripe_customer_id` varchar(255) NOT NULL COMMENT "Stripe customer ID", ADD COLUMN `role` enum('Free','Premium') NOT NULL DEFAULT "Free" COMMENT "User role";
-- drop "payments" table
DROP TABLE `payments`;
