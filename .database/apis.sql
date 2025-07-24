CREATE TABLE `apis` (
  `id_api` int NOT NULL AUTO_INCREMENT,
  `fid_site` varchar(45) DEFAULT NULL,
  `api_private` varchar(128) DEFAULT NULL,
  `api_public` varchar(128) DEFAULT NULL,
  `api_status` varchar(4) DEFAULT NULL,
  `api_requirements` text,
  `api_territories` text,
  PRIMARY KEY (`id_api`),
  KEY `fid_site` (`fid_site`),
  KEY `api_private` (`api_private`),
  KEY `api_public` (`api_public`),
  KEY `api_status` (`api_status`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
