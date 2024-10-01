-- modify "tasks" table
ALTER TABLE `tasks` ADD COLUMN `project_id` bigint unsigned NULL COMMENT "Project ID", ADD INDEX `project_id` (`project_id`), ADD CONSTRAINT `tasks_ibfk_2` FOREIGN KEY (`project_id`) REFERENCES `projects` (`id`) ON UPDATE NO ACTION ON DELETE NO ACTION;
