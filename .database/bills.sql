CREATE TABLE `myapp`.`bills` (
  `id_bill` INT NOT NULL,
  `bill_transaction` VARCHAR(64) NULL,
  `bill_type` VARCHAR(16) NULL,
  `bill_vendor` VARCHAR(16) NULL,
  `bill_amount` FLOAT(7,2) NULL,
  `bill_timestamp` TIMESTAMP NULL,
  `bill_meta` TEXT NULL,
  PRIMARY KEY (`id_bill`));
