CREATE TABLE `transactions` (
  `id_transaction` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `transaction_method` varchar(6) DEFAULT NULL,
  `transaction_merchant` varchar(16) DEFAULT NULL,
  `transaction_amount` decimal(18,8) NOT NULL,
  `transaction_id` varchar(64) DEFAULT NULL,
  `transaction_timestamp` timestamp NOT NULL DEFAULT current_timestamp(),
  PRIMARY KEY (`id_transaction`),
  UNIQUE KEY `uniq_transaction_id` (`transaction_id`),
  KEY `idx_transaction_method` (`transaction_method`),
  KEY `idx_transaction_merchant` (`transaction_merchant`),
  KEY `idx_transaction_timestamp` (`transaction_timestamp`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;
