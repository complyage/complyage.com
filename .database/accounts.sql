CREATE TABLE `accounts` (
  `id_account` int NOT NULL AUTO_INCREMENT,
  `account_type` varchar(4) DEFAULT NULL,
  `account_salt` varchar(256) DEFAULT NULL,
  `account_username` varchar(60) NOT NULL,
  `account_email` varchar(160) DEFAULT NULL,
  `account_password` varchar(256) DEFAULT NULL,
  `account_private` text,
  `account_public` text,
  `account_status` varchar(4) DEFAULT NULL,
  `account_level` tinyint DEFAULT NULL,
  `account_advanced` tinyint DEFAULT NULL,
  PRIMARY KEY (`id_account`),
  KEY `account_email` (`account_email`),
  KEY `idx_accounts_account_email` (`account_email`),
  KEY `account_username` (`account_username`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
