CREATE TABLE `shared` (
  `id_shared` int unsigned NOT NULL AUTO_INCREMENT,
  `fid_site` int DEFAULT NULL,
  `fid_verification` int DEFAULT NULL,
  `shared_timestamp` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_shared`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
