CREATE TABLE `verification_types` (
      `id_verification_type` int NOT NULL AUTO_INCREMENT,
      `verification_code` varchar(4) DEFAULT NULL,
      `verification_description` varchar(60) DEFAULT NULL,
      `verification_level` tinyint DEFAULT '1',
      PRIMARY KEY (`id_verification_type`)
) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;