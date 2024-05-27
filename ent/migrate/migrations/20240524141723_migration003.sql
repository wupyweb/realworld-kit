-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Create "new_tags" table
CREATE TABLE `new_tags` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL, `user_tags` integer NULL, CONSTRAINT `tags_users_tags` FOREIGN KEY (`user_tags`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Copy rows from old table "tags" to new temporary table "new_tags"
INSERT INTO `new_tags` (`id`, `name`, `user_tags`) SELECT `id`, `name`, `user_id` FROM `tags`;
-- Drop "tags" table after copying rows
DROP TABLE `tags`;
-- Rename temporary table "new_tags" to "tags"
ALTER TABLE `new_tags` RENAME TO `tags`;
-- Create index "tags_name_key" to table: "tags"
CREATE UNIQUE INDEX `tags_name_key` ON `tags` (`name`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
