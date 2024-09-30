-- modify "users" table
ALTER TABLE `users` MODIFY COLUMN `stripe_customer_id` varchar(255) NULL COMMENT "Stripe customer ID";
