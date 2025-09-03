CREATE TABLE`vpn_ratings` (
      `id_rating` int unsigned NOT NULL AUTO_INCREMENT,
      `fid_vpn` int DEFAULT NULL,
      `fid_user` int DEFAULT NULL,
      `rating` tinyint DEFAULT NULL,
      `comment` varchar(512) DEFAULT NULL,
      PRIMARY KEY (`id_rating`),
      KEY `fid_vpn` (`fid_vpn`),
      KEY `fid_user` (`fid_user`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;