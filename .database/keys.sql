CREATE TABLE `keys` (
  `id_key` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `fid_account` int(10) unsigned NOT NULL,
  `key_level` tinyint(4) DEFAULT 1,
  `key_private` text NOT NULL,
  `key_public` text NOT NULL,
  `key_check` char(64) NOT NULL,
  `key_created` datetime NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id_key`),
  KEY `idx_keys_account` (`fid_account`),
  CONSTRAINT `fk_keys_account` FOREIGN KEY (`fid_account`) REFERENCES `accounts` (`id_account`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
