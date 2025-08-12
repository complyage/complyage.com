CREATE TABLE
        `verifications` (
                `id_verification` bigint NOT NULL AUTO_INCREMENT,
                `fid_account` bigint NOT NULL,
                `verification_type` varchar(4) DEFAULT NULL,
                `verification_data` longblob,
                `verification_meta` text,
                `verification_status` varchar(4) DEFAULT NULL,
                `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
                `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                PRIMARY KEY (`id_verification`),
                KEY `idx_verifications_fid_account` (`fid_account`),
                KEY `idx_verifications_type` (`verification_type`),
                CONSTRAINT `fk_verifications_accounts` FOREIGN KEY (`fid_account`) REFERENCES `accounts` (`id_account`) ON DELETE CASCADE
        ) ENGINE = InnoDB AUTO_INCREMENT = 3 DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;