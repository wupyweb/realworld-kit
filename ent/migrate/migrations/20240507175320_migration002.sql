-- Disable the enforcement of foreign-keys constraints
PRAGMA foreign_keys = off;
-- Drop "user_favorites" table
DROP TABLE `user_favorites`;
-- Create "favorites" table
CREATE TABLE `favorites` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `favorited_at` datetime NOT NULL, `user_id` integer NOT NULL, `article_id` integer NOT NULL, CONSTRAINT `favorites_users_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE NO ACTION, CONSTRAINT `favorites_articles_article` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE NO ACTION);
-- Create index "favorite_user_id_article_id" to table: "favorites"
CREATE UNIQUE INDEX `favorite_user_id_article_id` ON `favorites` (`user_id`, `article_id`);
-- Enable back the enforcement of foreign-keys constraints
PRAGMA foreign_keys = on;
