CREATE TABLE `crypto` (
  `id_crypto` int NOT NULL AUTO_INCREMENT,
  `crypto_name` varchar(60) DEFAULT NULL,
  `crypto_address` varchar(255) DEFAULT NULL,
  PRIMARY KEY (`id_crypto`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
