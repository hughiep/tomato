-- reverse: modify "tasks" table
ALTER TABLE `tasks` DROP FOREIGN KEY `tasks_ibfk_2`, DROP INDEX `project_id`, DROP COLUMN `project_id`;
