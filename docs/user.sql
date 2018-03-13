CREATE TABLE `admin_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `salt` varchar(255) NOT NULL DEFAULT '',
  `sex` int(11) NOT NULL DEFAULT '0',
  `email` varchar(50) NOT NULL DEFAULT '',
  `last_login` datetime DEFAULT NULL,
  `last_ip` varchar(15) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

CREATE TABLE `session` (
	`session_key` char(64) NOT NULL,
	`session_data` blob,
	`session_expiry` int(11) unsigned NOT NULL,
	PRIMARY KEY (`session_key`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;

CREATE TABLE `action_log` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `action` varchar(20) NOT NULL DEFAULT '',
  `actor` varchar(20) NOT NULL DEFAULT '',
   `user_id` int(11) DEFAULT 0,
  `extra` varchar(1000) NOT NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

INSERT INTO `admin_user` (`id`, `user_name`, `password`, `salt`, `sex`, `email`, `last_login`, `last_ip`, `status`, `create_time`, `update_time`)
VALUES
	(1,'admin','7fef6171469e80d32c0559f88b377245','',1,'admin@admin.com','2016-05-11 10:33:49','127.0.0.1',0,'2016-05-11 10:33:49','2016-05-11 10:33:49');