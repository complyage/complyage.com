CREATE TABLE `bills` (
  `id_bill` int(11) NOT NULL AUTO_INCREMENT,
  `bill_transaction` varchar(64) DEFAULT NULL,
  `bill_type` varchar(16) DEFAULT NULL,
  `bill_vendor` varchar(16) DEFAULT NULL,
  `bill_amount` float(7,2) DEFAULT NULL,
  `bill_timestamp` timestamp NULL DEFAULT NULL,
  `bill_meta` text DEFAULT NULL,
  PRIMARY KEY (`id_bill`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
