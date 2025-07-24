CREATE TABLE
        `ips` (
                `id` int NOT NULL AUTO_INCREMENT,
                `start_ip` bigint NOT NULL,
                `end_ip` bigint NOT NULL,
                `city` varchar(200) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                `state` varchar(100) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                `country` varchar(5) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                `postal_code` varchar(20) COLLATE utf8mb4_unicode_ci DEFAULT NULL,
                `latitude` float (10, 6) DEFAULT NULL,
                `longitude` float (10, 6) DEFAULT NULL,
                PRIMARY KEY (`id`),
                KEY `idx_ip_range` (`start_ip`, `end_ip`),
                KEY `idx_ip_start` (`start_ip`),
                KEY `idx_ip_lookup` (`start_ip`, `end_ip`)
        ) ENGINE = InnoDB AUTO_INCREMENT = 2915986 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;