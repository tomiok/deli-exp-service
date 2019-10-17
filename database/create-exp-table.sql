CREATE TABLE `product` (
  `uid` varchar(100) NOT NULL,
  `name` varchar(100) NOT NULL,
  `date` date NOT NULL,
  `city` varchar(150) DEFAULT NULL,
  `country` varchar(100) DEFAULT NULL,
  `details` varchar(400) DEFAULT NULL,
  PRIMARY KEY (`uid`)
) ENGINE=Inn