# database

## account_user table
CREATE TABLE `account_user` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT 'account',
  `email` varchar(30) NOT NULL DEFAULT '' COMMENT 'e-mail',
  `phone` varchar(15) NOT NULL DEFAULT '' COMMENT 'phone number',
  `username` varchar(30) NOT NULL DEFAULT '' COMMENT 'user name',
  `password` varchar(32) NOT NULL DEFAULT '' COMMENT 'password',
  `create_at` int(11) NOT NULL DEFAULT '0' COMMENT 'create time',
  `create_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'create IP',
  `last_login_at` int(11) NOT NULL DEFAULT '0' COMMENT 'last login time',
  `last_login_ip_at` varchar(12) NOT NULL DEFAULT '' COMMENT 'last login ip',
  `login_times` int(11) NOT NULL DEFAULT '0' COMMENT 'login times',
  `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT 'status 1:enable, 0:disable, -1:deleted',
  PRIMARY KEY (`id`),
  KEY `idx_email` (`email`),
  KEY `idx_phone` (`phone`),
  KEY `idx_username` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='account';

insert into account_user values (12345678, 'arks@test.com', '0900111222', 'test', 'test', 1234567, '172.16.10.10', 7654321, '172.16.10.20', 7777, 1);