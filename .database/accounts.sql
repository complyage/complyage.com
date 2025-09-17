CREATE TABLE `accounts` (
  `id_account` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `account_salt` varchar(256) DEFAULT NULL,
  `account_identifier` varchar(160) NOT NULL,
  `account_username` varchar(64) CHARACTER SET ascii COLLATE ascii_bin NOT NULL,
  `account_password` varchar(256) DEFAULT NULL,
  `account_status` char(4) CHARACTER SET ascii COLLATE ascii_bin DEFAULT NULL,
  `account_level` tinyint(3) unsigned DEFAULT NULL,
  `account_created` timestamp NOT NULL DEFAULT current_timestamp(),
  `account_login` timestamp NULL DEFAULT current_timestamp() ON UPDATE current_timestamp(),
  `account_identity` text DEFAULT '{}',
  PRIMARY KEY (`id_account`),
  UNIQUE KEY `uniq_account_username` (`account_username`),
  UNIQUE KEY `uniq_account_identifier` (`account_identifier`),
  KEY `idx_account_status` (`account_status`),
  KEY `idx_account_level` (`account_level`),
  KEY `idx_account_created` (`account_created`),
  KEY `idx_account_login` (`account_login`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_bin;
