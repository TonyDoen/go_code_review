CREATE DATABASE IF NOT EXISTS demo  DEFAULT CHARSET utf8mb4 COLLATE utf8mb4_unicode_ci;

use demo;

CREATE TABLE `demo` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` VARCHAR(24) NOT NULL,
  `desc` VARCHAR(128) NOT NULL,
  `create_time` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `IDX_NAME` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='demo';

INSERT INTO demo.demo (`id`, `name`, `desc`, `create_time`) VALUES (1, 'zhd', 'man2333', '2006-08-31 05:19:50');
