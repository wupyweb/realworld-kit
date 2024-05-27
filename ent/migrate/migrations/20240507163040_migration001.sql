-- Create "articles" table
CREATE TABLE `articles` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `title` text NOT NULL, `description` text NOT NULL, `body` text NOT NULL, `slug` text NOT NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL, `user_id` integer NULL, CONSTRAINT `articles_users_articles` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "articles_slug_key" to table: "articles"
CREATE UNIQUE INDEX `articles_slug_key` ON `articles` (`slug`);
-- Create "comments" table
CREATE TABLE `comments` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `body` text NOT NULL, `created_at` datetime NOT NULL, `article_id` integer NULL, `user_id` integer NULL, CONSTRAINT `comments_articles_comments` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE SET NULL, CONSTRAINT `comments_users_comments` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create "tags" table
CREATE TABLE `tags` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `name` text NOT NULL, `user_id` integer NULL, CONSTRAINT `tags_users_tags` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE SET NULL);
-- Create index "tags_name_key" to table: "tags"
CREATE UNIQUE INDEX `tags_name_key` ON `tags` (`name`);
-- Create "users" table
CREATE TABLE `users` (`id` integer NOT NULL PRIMARY KEY AUTOINCREMENT, `username` text NOT NULL, `password` text NOT NULL, `email` text NOT NULL, `bio` text NULL, `image` text NULL, `created_at` datetime NOT NULL, `updated_at` datetime NOT NULL);
-- Create index "users_email_key" to table: "users"
CREATE UNIQUE INDEX `users_email_key` ON `users` (`email`);
-- Create "article_tags" table
CREATE TABLE `article_tags` (`article_id` integer NOT NULL, `tag_id` integer NOT NULL, PRIMARY KEY (`article_id`, `tag_id`), CONSTRAINT `article_tags_article_id` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE, CONSTRAINT `article_tags_tag_id` FOREIGN KEY (`tag_id`) REFERENCES `tags` (`id`) ON DELETE CASCADE);
-- Create "user_favorites" table
CREATE TABLE `user_favorites` (`user_id` integer NOT NULL, `article_id` integer NOT NULL, PRIMARY KEY (`user_id`, `article_id`), CONSTRAINT `user_favorites_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_favorites_article_id` FOREIGN KEY (`article_id`) REFERENCES `articles` (`id`) ON DELETE CASCADE);
-- Create "user_following" table
CREATE TABLE `user_following` (`user_id` integer NOT NULL, `follower_id` integer NOT NULL, PRIMARY KEY (`user_id`, `follower_id`), CONSTRAINT `user_following_user_id` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`) ON DELETE CASCADE, CONSTRAINT `user_following_follower_id` FOREIGN KEY (`follower_id`) REFERENCES `users` (`id`) ON DELETE CASCADE);
