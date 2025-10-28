CREATE TABLE `scopes` (
      `id_scope` int(10) unsigned NOT NULL AUTO_INCREMENT,
      `fid_website` int(11) DEFAULT NULL,
      `scope_code` char(4) DEFAULT NULL,
      `scope_status` char(4) DEFAULT NULL,
      `scope_moderator` int(11) DEFAULT NULL,
      `scope_created` timestamp NULL DEFAULT current_timestamp(),
      `scope_updated` timestamp NULL DEFAULT NULL ON UPDATE current_timestamp(),
      PRIMARY KEY (`id_scope`),
      KEY `fid_website` (`fid_website`),
      KEY `scope_code` (`scope_code`),
      KEY `scope_status` (`scope_status`),
      KEY `scope_created` (`scope_created`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_uca1400_ai_ci;
