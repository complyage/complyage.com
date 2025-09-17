CREATE TABLE `verify` (
  `id_verify` BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
  `fid_account` INT UNSIGNED NOT NULL,
  `verify_uuid` CHAR(36) NOT NULL,
  `verify_display` VARCHAR(36) NOT NULL,
  `verify_type` CHAR(4) NOT NULL,
  `verify_status` VARCHAR(4) NOT NULL DEFAULT 'PEND',
  `verify_created` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `verify_updated` DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_verify`),
  UNIQUE KEY `uniq_verify_uuid` (`verify_uuid`),
  KEY `idx_verify_fid_account` (`fid_account`),
  KEY `idx_verify_type` (`verify_type`),
  KEY `idx_verify_status` (`verify_status`),
  KEY `idx_verify_created` (`verify_created`),
  KEY `idx_verify_account_type` (`fid_account`,`verify_type`),
  CONSTRAINT `fk_verify_accounts`
    FOREIGN KEY (`fid_account`) REFERENCES `accounts` (`id_account`)
    ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;