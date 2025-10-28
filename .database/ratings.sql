CREATE TABLE `ratings` (
  `id_rating` INT UNSIGNED NOT NULL AUTO_INCREMENT,
  `fid_vpn` INT NOT NULL,
  `fid_user` INT DEFAULT NULL,
  `rating` TINYINT DEFAULT NULL,
  `comment` VARCHAR(512) DEFAULT NULL,
  PRIMARY KEY (`id_rating`),
  KEY `idx_ratings_fid_vpn` (`fid_vpn`),
  KEY `idx_ratings_fid_user` (`fid_user`),
  CONSTRAINT `fk_ratings_vpn`
    FOREIGN KEY (`fid_vpn`)
    REFERENCES `vpns` (`id_vpn`)
    ON DELETE CASCADE
    ON UPDATE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_0900_ai_ci;
