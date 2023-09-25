CREATE TABLE `notes` (
                         `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
                         `title` varchar(150) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                         `category_id` int(11) unsigned NOT NULL,
                         `status` int(3) DEFAULT NULL,
                         `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                         `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                         PRIMARY KEY (`id`),
                         UNIQUE KEY `id` (`id`),
                         KEY `category` (`category_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci