-- # ************************************************************
-- # Sequel Pro SQL dump
-- # Version 4541
-- #
-- # http://www.sequelpro.com/
-- # https://github.com/sequelpro/sequelpro
-- #
-- # Host: 127.0.0.1 (MySQL 5.7.20)
-- # Database: tax_calculator_db
-- # Generation Time: 2019-02-04 10:06:49 +0000
-- # ************************************************************


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;
/*!40111 SET @OLD_SQL_NOTES=@@SQL_NOTES, SQL_NOTES=0 */;


-- # Dump of table bills
-- # ------------------------------------------------------------

DROP TABLE IF EXISTS `bills`;

CREATE TABLE `bills` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(255) DEFAULT NULL,
  `tax_code` int(11) DEFAULT NULL,
  `price` float DEFAULT NULL,
  `created_at` datetime DEFAULT NULL,
  `updated_at` datetime DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `bills` WRITE;
/*!40000 ALTER TABLE `bills` DISABLE KEYS */;

INSERT INTO `bills` (`id`, `name`, `tax_code`, `price`, `created_at`, `updated_at`)
VALUES
	(1,'Lucky Stretch',2,1000,'2019-02-03 13:23:15','2019-02-03 13:23:15'),
	(2,'Big Mac',1,1000,'2019-02-03 13:23:34','2019-02-03 13:23:34'),
	(3,'Movie',3,150,'2019-02-03 13:23:51','2019-02-03 13:23:51'),
	(4,'Avenger',3,200,'2019-02-03 15:36:12','2019-02-03 15:36:12');

/*!40000 ALTER TABLE `bills` ENABLE KEYS */;
UNLOCK TABLES;


-- # Dump of table taxs
-- # ------------------------------------------------------------

DROP TABLE IF EXISTS `taxs`;

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
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

LOCK TABLES `taxs` WRITE;
/*!40000 ALTER TABLE `taxs` DISABLE KEYS */;

INSERT INTO `taxs` (`id`, `tax_code`, `type`, `condition`, `formula`, `refundable`, `created_at`, `updated_at`)
VALUES
	(1,1,'Food & Beverage','','0.1 * {price}','YES','2019-02-02 01:30:49','2019-02-02 01:30:49'),
	(2,2,'Tobacco','','10 + (0.02 * {price})','NO','2019-02-02 01:30:49','2019-02-02 01:30:49'),
	(3,3,'Entertainment','{price} >= 100','0.01 * ({price} - 100)','NO','2019-02-02 01:32:16','2019-02-02 01:32:16');

/*!40000 ALTER TABLE `taxs` ENABLE KEYS */;
UNLOCK TABLES;

/*!40111 SET SQL_NOTES=@OLD_SQL_NOTES */;
/*!40101 SET SQL_MODE=@OLD_SQL_MODE */;
/*!40014 SET FOREIGN_KEY_CHECKS=@OLD_FOREIGN_KEY_CHECKS */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
