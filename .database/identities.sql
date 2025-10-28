CREATE TABLE identities (
    id_identity           INT UNSIGNED NOT NULL AUTO_INCREMENT,
    fid_account           INT UNSIGNED NOT NULL,
    identity_type         CHAR(4)      NOT NULL,
    identity_verified     TINYINT(1)   NOT NULL DEFAULT 0,
    identity_age          TINYINT      DEFAULT -1,
    identity_dob          DATE         DEFAULT '2070-01-01',
    identity_display      VARCHAR(32)  DEFAULT NULL,
    identity_verification VARCHAR(36)  UNIQUE DEFAULT NULL,
    identity_created      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP,
    identity_updated      DATETIME     NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (id_identity),
    KEY idx_identity_account (fid_account),
    CONSTRAINT fk_identity_account FOREIGN KEY (fid_account)
        REFERENCES accounts(id_account)
        ON DELETE CASCADE
        ON UPDATE CASCADE
) ENGINE=InnoDB
  DEFAULT CHARSET=utf8mb4
  COLLATE=utf8mb4_0900_ai_ci;
  