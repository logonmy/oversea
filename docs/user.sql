
-- 管理员表
CREATE TABLE `oz_admin_user` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `user_name` varchar(20) NOT NULL DEFAULT '',
  `password` varchar(32) NOT NULL DEFAULT '',
  `salt` varchar(255) NOT NULL DEFAULT '',
  `sex` int(11) NOT NULL DEFAULT '0',
  `email` varchar(50) NOT NULL DEFAULT '',
  `phone` varchar(12)  DEFAULT NULL ,
  `last_login` datetime DEFAULT NULL,
  `last_ip` varchar(15) NOT NULL DEFAULT '',
  `status` int(11) NOT NULL DEFAULT '0',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`id`),
  UNIQUE KEY `user_name` (`user_name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 操作日志表
CREATE TABLE `oz_action_log` (
  `id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT,
  `object_type` varchar(30) NOT NULL DEFAULT '',
  `object_id` mediumint(8) unsigned NOT NULL DEFAULT '0',
  `user_id` int(11) DEFAULT NULL,
  `actor` varchar(30) NOT NULL DEFAULT '',
  `action` varchar(30) NOT NULL DEFAULT '',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `comment` text,
  `extra` text NOT NULL,
  `read` tinyint(4) DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `action` (`object_id`,`action`,`create_time`),
  KEY `read` (`object_type`,`read`) USING BTREE,
  KEY `user_id` (`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- 移民项目类型表
CREATE TABLE `oz_project_type` (
  `project_type_id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '项目类型id',
  `project_type_name` varchar(30) NOT NULL DEFAULT '' COMMENT '项目类型名称',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态，0正常 -1禁用',
  PRIMARY KEY (`id`),
  UNIQUE KEY `project_type_name` (`project_type_name`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='项目类型表';

-- 移民国家表
CREATE TABLE `oz_immigrant_nation` (
  `nation_id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '国家或地区id',
  `nation_name` varchar(30) NOT NULL DEFAULT '' COMMENT '国家或地区名称',
  `flag` VARCHAR(500) DEFAULT NULL COMMENT '国旗图标',
  `desc` text DEFAULT NULL COMMENT '简介描素',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态，0正常 -1禁用',
  PRIMARY KEY (`nation_id`),
  UNIQUE KEY `nation_name` (`nation_name`) USING BTREE,
  UNIQUE KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='移民国家表';


-- 移民项目表
CREATE TABLE `oz_immigrant_project` (
  `project_id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '项目id',
  `project_name` varchar(30) NOT NULL DEFAULT '' COMMENT '项目名称',
  `nation_id` mediumint(8) NOT NULL COMMENT '移民国家id',
  `project_type_id`  mediumint(8) NOT NULL COMMENT '移民类型id',
  `investment_amount` varchar(50) NOT NULL DEFAULT '' COMMENT '投资金额',
  `complexity` tinyint(4) DEFAULT '0' COMMENT '手续复杂度，0-简单; 1-普通; 2-困难',
  `residency_requirement` varchar(50) NOT NULL DEFAULT '' COMMENT '居住要求',
  `job_requirement` varchar(50) NOT NULL DEFAULT '' COMMENT '工作要求',
  `cycle_time` varchar(20) NOT NULL DEFAULT '' COMMENT '办理周期',
  `visa_type` varchar(20) NOT NULL DEFAULT '' COMMENT '签证类型',
  `project_desc` text  COMMENT '项目介绍',
  `apply_requirement` text COMMENT '申请条件描述',
  `policy_advantage` text  COMMENT '政策优势描述',
  `handling_procedures` text  COMMENT '办理流程描述',
  `apply_list` text  COMMENT '申请材料清单',
  `charge_list` text  COMMENT '费用清单',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态，0正常 -1禁用',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ,
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`project_id`),
  UNIQUE KEY `nation_name` (`project_name`) USING BTREE,
  UNIQUE KEY `status` (`status`) USING BTREE,
  UNIQUE KEY `project_type_id` (`project_type_id`) USING BTREE,
  UNIQUE KEY `nation_id` (`nation_id`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='移民项目表';


-- ----------------------------
-- 5、后台菜单权限表
-- ----------------------------
drop table if exists sys_menu;
create table sys_menu (
  menu_id 			int(11) 		not null auto_increment    comment '菜单ID',
  menu_name 		varchar(50) 	not null 				   comment '菜单名称',
  parent_id 		int(11) 		default 0 			       comment '父菜单ID',
  order_num 		int(4) 			default null 			   comment '显示顺序',
  url 				varchar(200) 	default ''				   comment '请求地址',
  menu_type 		char(1) 		default '' 			       comment '类型:M目录,C菜单,F按钮',
  visible 			int(1) 			default 0 				   comment '菜单状态:0显示,1隐藏',
  perms 			varchar(100) 	default '' 				   comment '权限标识',
  icon 				varchar(100) 	default '' 				   comment '菜单图标',
  create_by         varchar(64)     default ''                 comment '创建者',
  create_time 		timestamp       default current_timestamp  comment '创建时间',
  update_time 		timestamp       default current_timestamp  comment '更新时间',
  update_by 		varchar(64) 	default ''			       comment '更新者',
  remark 			varchar(500) 	default '' 				   comment '备注',
  primary key (menu_id)
) engine=innodb auto_increment=1000 default charset=utf8;


-- ----------------------------
-- 11、系统访问记录
-- ----------------------------
drop table if exists sys_logininfor;
create table sys_logininfor (
  info_id 		int(11) 	 not null auto_increment   comment '访问ID',
  login_name 	varchar(50)  default '' 			   comment '登录账号',
  ipaddr 		varchar(50)  default '' 			   comment '登录IP地址',
  browser  		varchar(50)  default '' 			   comment '浏览器类型',
  os      		varchar(50)  default '' 			   comment '操作系统',
  status 		int(1) 		 default 0 			 	   comment '登录状态 0成功 1失败',
  msg      		varchar(255) default '' 			   comment '提示消息',
  login_time 	timestamp    default current_timestamp comment '访问时间',
  primary key (info_id)
) engine=innodb auto_increment=100 default charset=utf8;


-- ----------------------------
-- 12、在线用户记录
-- ----------------------------
drop table if exists sys_user_online;
create table sys_user_online (
  sessionId 	    varchar(50)  default ''              	comment '用户会话id',
  login_name 	    varchar(50)  default '' 		 	 	comment '登录账号',
  dept_name 		varchar(50)  default '' 		 	 	comment '部门名称',
  ipaddr 		    varchar(50)  default '' 			 	comment '登录IP地址',
  browser  		    varchar(50)  default '' 			 	comment '浏览器类型',
  os      		    varchar(50)  default '' 			 	comment '操作系统',
  status      	    varchar(10)  default '' 			 	comment '在线状态on_line在线off_line离线',
  start_timestsamp 	timestamp    default current_timestamp  comment 'session创建时间',
  last_access_time  timestamp    default current_timestamp  comment 'session最后访问时间',
  expire_time 	    int(5) 		 default 0 			 	    comment '超时时间，单位为分钟',
  primary key (sessionId)
) engine=innodb default charset=utf8;


-- ----------------------------
-- 12、客户表
-- ----------------------------


INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('1', '投资移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('2', '技术移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('3', '创业移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('4', '杰出人才', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('5', '留学生移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('6', '家属团聚', '0');


INSERT INTO `oz_admin_user` (`id`, `user_name`, `password`, `salt`, `sex`, `email`, `last_login`, `last_ip`, `status`,
`create_time`, `update_time`)
VALUES
	(1,'admin','7fef6171469e80d32c0559f88b377245','',1,'admin@admin.com','2016-05-11 10:33:49','127.0.0.1',0,'2016-05-11 10:33:49','2016-05-11 10:33:49');