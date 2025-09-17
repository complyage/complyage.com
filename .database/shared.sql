CREATE TABLE `shared` (
      `id_shared` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `fid_site` int(11) DEFAULT NULL,
      `shared_type` char(4) DEFAULT NULL,
      `shared_verification` varchar(36) DEFAULT NULL,
      `shared_timestamp` timestamp NULL DEFAULT current_timestamp(),
      PRIMARY KEY (`id_shared`),
      KEY `fid_site` (`fid_site`),
      KEY `shared_type` (`shared_type`),
      KEY `shared_verification` (`shared_verification`),
      KEY `shared_timestamp` (`shared_timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
