CREATE TABLE `product` (
                           `uid` varchar(100) NOT NULL,
                           `name` varchar(100) NOT NULL,
                           `date` date NOT NULL,
                           `city` varchar(150) DEFAULT NULL,
                           `country` varchar(100) DEFAULT NULL,
                           `details` varchar(400) DEFAULT NULL,
                           PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `tags` (
                        `uid` varchar(100) NOT NULL,
                        `csv_tags` varchar(100) NOT NULL,
                        PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

CREATE TABLE `experience_post` (
                                   `uid` varchar(100) NOT NULL,
                                   `title` varchar(100) NOT NULL,
                                   `subtitle` varchar(150) NOT NULL,
                                   `body` text NOT NULL,
                                   `date` date NOT NULL,
                                   `author_uid` varchar(150) NOT NULL,
                                   `tags_uid` varchar(100) DEFAULT NULL,
                                   `products_uid` varchar(400) DEFAULT NULL,
                                   `photo_url` varchar(100) DEFAULT NULL,
                                   PRIMARY KEY (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

ALTER TABLE experience_post ADD CONSTRAINT experience_post_FK FOREIGN KEY (products_uid) REFERENCES product(uid);
ALTER TABLE experience_post ADD CONSTRAINT experience_post_FK_1 FOREIGN KEY (tags_uid) REFERENCES tags(uid);