CREATE TABLE `vpns` (
  `id_vpn` INT NOT NULL AUTO_INCREMENT,
  `vpn_name` VARCHAR(64) DEFAULT NULL,
  `vpn_url` VARCHAR(512) DEFAULT NULL,
  `vpn_blurb` VARCHAR(256) DEFAULT NULL,
  `vpn_highlights` VARCHAR(256) DEFAULT NULL,
  `vpn_region` VARCHAR(64) DEFAULT NULL,
  `vpn_price` VARCHAR(16) DEFAULT NULL,
  `vpn_rating` INT DEFAULT NULL,
  PRIMARY KEY (`id_vpn`),
  KEY `idx_vpn_rating` (`vpn_rating` DESC),
  KEY `idx_vpn_price` (`vpn_price`),
  KEY `idx_vpn_name` (`vpn_name`),
  KEY `idx_vpn_price_rating_name` (`vpn_price`, `vpn_rating`, `vpn_name`)
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_0900_ai_ci;
