CREATE TABLE `accounts` (
  `id_account` int NOT NULL AUTO_INCREMENT,
  `account_type` varchar(4) DEFAULT NULL,
  `account_salt` varchar(256) DEFAULT NULL,
  `account_username` varchar(64) NOT NULL,
  `account_email` varchar(160) NOT NULL,
  `account_password` varchar(256) DEFAULT NULL,
  `account_security` tinyint DEFAULT 1,
  `account_public` text,
  `account_private` text,
  `account_private_hash` varchar(64) DEFAULT NULL,
  `account_status` varchar(4) DEFAULT NULL,
  `account_level` tinyint DEFAULT NULL,
  PRIMARY KEY (`id_account`),
  UNIQUE KEY `uniq_account_username` (`account_username`),
  UNIQUE KEY `uniq_account_email` (`account_email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;