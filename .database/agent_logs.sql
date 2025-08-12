CREATE TABLE `agent_logs`.`requests` (
  `id_request` BIGINT NOT NULL AUTO_INCREMENT,
  `fid_site` INT NULL,
  `request_level` TINYINT NULL,
  `request_id` VARCHAR(64) NULL,
  `request_timestamp` TIMESTAMP NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id_request`));
