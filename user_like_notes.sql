CREATE TABLE `user_like_notes` (
                                   `user_id` int(11) NOT NULL,
                                   `note_id` int(11) NOT NULL,
                                   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                   PRIMARY KEY (`user_id`,`note_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci