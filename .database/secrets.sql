CREATE TABLE `secrets` (
  `id_secret` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `fid_account` int(10) unsigned NOT NULL,
  `secret_level` tinyint(4) DEFAULT 1,
  `secret_private` text NOT NULL,
  `secret_public` text NOT NULL,
  `secret_check` char(64) NOT NULL,
  `secret_created` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id_secret`),
  UNIQUE KEY `uniq_secret_account` (`fid_account`),
  KEY `idx_keys_account` (`fid_account`),
  CONSTRAINT `fk_secret_account` FOREIGN KEY (`fid_account`) REFERENCES `accounts` (`id_account`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=17 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
