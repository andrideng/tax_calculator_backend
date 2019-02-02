-- create database
CREATE DATABASE tax_calculator_db;

-- Tax Table
CREATE TABLE `taxs` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `tax_code` int(11) DEFAULT NULL,
  `type` varchar(100) DEFAULT NULL,
  `condition` varchar(255) DEFAULT NULL,
  `formula` varchar(255) DEFAULT NULL,
  `refundable` varchar(10) DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=latin1;

-- INSERT TAXS QUERY
INSERT into taxs(`tax_code`, `type`, `formula`, `refundable`, `created_at`, `updated_at`) VALUES(1, 'Food & Beverage', '0.01 * {price}', 'YES', now(), now());
INSERT into taxs(`tax_code`, `type`, `formula`, `refundable`, `created_at`, `updated_at`) VALUES(2, 'Tobacco', '10 + (0.02 * {price})', 'NO', now(), now());
INSERT into taxs(`tax_code`, `type`, `formula`, `condition`, `refundable`, `created_at`, `updated_at`) VALUES(3, 'Entertainment', '0.01 * ({price} - 100)', '0 < {price} < 100', 'NO', now(), now());

-- Bill Table
CREATE TABLE `bills` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `tax_code` int(11) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=latin1;