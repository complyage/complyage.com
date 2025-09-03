CREATE TABLE `vpns` (
  `id_vpn` int NOT NULL AUTO_INCREMENT,
  `vpn_name` varchar(64) DEFAULT NULL,
  `vpn_url` varchar(512) DEFAULT NULL,
  `vpn_blurb` varchar(256) DEFAULT NULL,
  `vpn_highlights` varchar(256) DEFAULT NULL,
  `vpn_region` varchar(64) DEFAULT NULL,
  `vpn_price` varchar(16) DEFAULT NULL,
  `vpn_rating` int DEFAULT NULL,
  PRIMARY KEY (`id_vpn`),
  KEY `vpn_rating` (`vpn_rating` DESC) /*!80000 INVISIBLE */,
  KEY `vpn_price` (`vpn_price` DESC)
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
