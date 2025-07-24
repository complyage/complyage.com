CREATE TABLE `verifications` (
  `id_verification` int NOT NULL,
  `fid_account` varchar(256) DEFAULT NULL,
  `verification_type` varchar(4) DEFAULT NULL,
  `verification_data` text,
  `verification_meta` text,
  `verifcation_status` varchar(4) DEFAULT NULL,
  PRIMARY KEY (`id_verification`),
  KEY `fid_account` (`fid_account`) /*!80000 INVISIBLE */,
  KEY `verification_type` (`verification_type`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
