CREATE TABLE `albums` (
                          `id` bigint NOT NULL AUTO_INCREMENT,
                          `created_at` datetime(3) DEFAULT NULL,
                          `updated_at` datetime(3) DEFAULT NULL,
                          `deleted_at` datetime(3) DEFAULT NULL,
                          `title` longtext,
                          `artist` longtext,
                          `price` double DEFAULT NULL,
                          PRIMARY KEY (`id`),
                          KEY `idx_albums_deleted_at` (`deleted_at`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci