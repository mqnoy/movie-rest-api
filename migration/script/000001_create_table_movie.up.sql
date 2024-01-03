CREATE TABLE `movie` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `title` varchar(255) DEFAULT "",
    `description` longtext DEFAULT "",
    `rating` decimal(2,1) DEFAULT 0,
    `image` varchar(255) DEFAULT "",
    `created_at` datetime(3) DEFAULT NULL,
    `updated_at` datetime(3) DEFAULT NULL,
     PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;
