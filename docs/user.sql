
-- 管理员表
CREATE TABLE `sys_admin_user` (
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
CREATE TABLE `sys_action_log` (
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

ALTER TABLE `sys_action_log`
ADD COLUMN `contact`  int NULL COMMENT '客户联系人' AFTER `read`,
ADD COLUMN `customer`  int NULL COMMENT '客户' AFTER `contact`,
ADD INDEX `customer` (`customer`) USING BTREE ;


-- 移民项目类型表
CREATE TABLE `oz_project_type` (
  `project_type_id` mediumint(8) unsigned NOT NULL AUTO_INCREMENT COMMENT '项目类型id',
  `project_type_name` varchar(30) NOT NULL DEFAULT '' COMMENT '项目类型名称',
  `status` tinyint(4) DEFAULT '0' COMMENT '状态，0正常 -1禁用',
  PRIMARY KEY (`project_type_id`),
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
-- 12、客户表 crm_customer
-- ----------------------------
drop table if exists crm_customer;
CREATE TABLE `crm_customer` (
  `cust_id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `name` varchar(64) NOT NULL COMMENT '客户姓名',
  `source` int(16) DEFAULT NULL COMMENT '客户来源',
  `level` int(16) DEFAULT NULL COMMENT '客户等级',
  `website` varchar(256) DEFAULT NULL COMMENT '客户个人网站地址',
  `mobile` varchar(64) NOT NULL COMMENT '客户手机号码',
  `tel` varchar(256) DEFAULT NULL COMMENT '客户电话号码',
  `fax` varchar(256) DEFAULT NULL COMMENT '客户传真',
  `email` varchar(256) DEFAULT NULL COMMENT '客户邮箱地址',
  `wechat` VARCHAR (256) DEFAULT NULL COMMENT '微信',
  `status` int(10) DEFAULT 0 COMMENT '客户状态: 0-正常，1-禁用',
  `sex` int(2) DEFAULT 0 comment '性别: 0-未知，1-男， 2-女',
  `native_place` varchar(10) comment '籍贯',
  `address` varchar(255) comment '家庭住址',
  `idcard` VARCHAR(30) DEFAULT NULL comment '身份证号',
  `capital` varchar(30) DEFAULT NULL  comment '资本描素',
  `intro` text COMMENT '客户简介',
  `create_by` int(11) NOT NULL DEFAULT '0' COMMENT '创建者',
  `assign_to` int(11) NOT NULL DEFAULT '0' COMMENT '指派给',
  `assign_status` int(11) NOT NULL DEFAULT '0' COMMENT '指派状态: 0-未指派，1-已指派，2-无需指派',
  `assign_time` timestamp NULL DEFAULT NULL COMMENT '指派日期',
  `contacted_date` datetime   DEFAULT NULL  COMMENT '拜访日期',
  `next_date` datetime  DEFAULT NULL COMMENT '下次拜访日期',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`cust_id`),
  KEY `name` (`name`),
  KEY `assign_to` (`assign_to`) USING BTREE,
  KEY `create_by` (`create_by`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 comment '客户表';

ALTER TABLE `crm_customer`
ADD COLUMN `link_address`  varchar(255) NULL COMMENT '联系地址' AFTER `address`;
ALTER TABLE `crm_customer`
ADD COLUMN `birthday`  date DEFAULT NULL COMMENT '生日' AFTER `address`;
ALTER TABLE `crm_customer`
ADD COLUMN `age`  int(4) NULL COMMENT '年龄' AFTER `address`;

ALTER TABLE `crm_customer`
ADD COLUMN `qq`  VARCHAR(20) NULL COMMENT 'qq' AFTER `wechat`;


ALTER TABLE `crm_customer`
ADD COLUMN `intension`  varchar(255) NULL AFTER `capital`;

-- ----------------------------
-- 联系人表 linkman
-- ----------------------------
DROP TABLE IF EXISTS `crm_linkman`;
CREATE TABLE `crm_linkman` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `cust_id` int(11) DEFAULT '0' COMMENT '客户id',
  `name` varchar(255) DEFAULT NULL COMMENT '联系人姓名',
  `job` varchar(255) DEFAULT NULL COMMENT '职业',
  `call` varchar(255) DEFAULT NULL COMMENT '固定电话',
  `phone` varchar(255) DEFAULT NULL COMMENT '手机号',
  `qq` varchar(255) DEFAULT NULL COMMENT 'QQ',
  `email` varchar(255) DEFAULT NULL COMMENT '邮箱',
  `main` text,
  `sex` tinyint(4) DEFAULT NULL,
  `address` varchar(255) DEFAULT NULL,
  `intro` varchar(255) DEFAULT NULL,
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `create_time` (`create_time`) USING BTREE,
  KEY `name` (`name`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8 COMMENT='联系人信息表';


ALTER TABLE `crm_linkman`
ADD COLUMN `wechat`  VARCHAR(20) NULL COMMENT '微信' AFTER `qq`;

-- ----------------------------
-- 12、客户地址表
-- ----------------------------
drop table if exists crm_customer_addr;
create table if not exists crm_customer_addr (
    id int auto_increment primary key comment '主键编号',
    customer_id int not null comment '客户表主键编号',
    `name` varchar(32) not null comment '姓名',
    country varchar(32) not null default '' comment '国家',
    province varchar(32) not null default '' comment '省份名称',
    city varchar(32) not null default '' comment '城市名称',
    area varchar(32) not null default '' comment '地区名称',
    street varchar(200) not null default '' comment '街道',
    zip varchar(8) not null default '' comment '邮政编码',
    telphone varchar(32) not null default '' comment '电话号码',
    mobile varchar(32) not null default '' comment '手机号码',
    is_default char(1) not null default 0 comment '是否默认地址',
    sort smallint not null default 0 comment '排序',
    create_time timestamp not null default CURRENT_TIMESTAMP comment '创建时间',
    INDEX idx_customer_id (customer_id),
    INDEX idxu_telphone_street (telphone,street)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8 comment '客户地址表';

-- ----------------------------
-- 13、系统配置表
-- ----------------------------
drop table if exists sys_config;
CREATE TABLE `sys_config` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键编号',
  `config_key` varchar(60) NOT NULL COMMENT '配置名称',
  `config_value` varchar(255) DEFAULT NULL COMMENT '姓名',
  `config_group` int(11) NOT NULL DEFAULT '0' COMMENT '配置分组',
  PRIMARY KEY (`id`),
  UNIQUE KEY `config_key` (`config_group`,`config_key`) USING BTREE
) ENGINE=MyISAM DEFAULT CHARSET=utf8 COMMENT='系统配置表';

-- ----------------------------
-- 14、文章表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `cms_article` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL COMMENT '管理员ID',
  `tid` int(11) NOT NULL DEFAULT '0' COMMENT '模板ID',
  `title` varchar(80) NOT NULL DEFAULT '' COMMENT '标题',
  `sub_title` varchar(80) NOT NULL DEFAULT '' COMMENT '副标题',
  `color` char(24) NOT NULL DEFAULT '' COMMENT '标题颜色',
  `font` char(24) NOT NULL DEFAULT '' COMMENT '标题加粗',
  `thumb` varchar(255) NOT NULL DEFAULT '' COMMENT '图片地址',
  `content` text NOT NULL COMMENT '内容',
  `copy_from` varchar(100) NOT NULL DEFAULT '' COMMENT '来源',
  `keywords` varchar(100) NOT NULL DEFAULT '' COMMENT '关键字',
  `description` varchar(250) NOT NULL COMMENT '描述',
  `relation` varchar(255) NOT NULL DEFAULT '' COMMENT '相关文章',
  `page_type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '分页方式',
  `max_char_per_page` mediumint(6) NOT NULL DEFAULT '0' COMMENT '分页字符数',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否发布 0:不发布 1:发布',
  `hits` tinyint(5) NOT NULL DEFAULT '0' COMMENT '点击数',
  `is_comment` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否允许评论',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  KEY `uid` (`uid`),
  KEY `create_time` (`create_time`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='内容表' AUTO_INCREMENT=0 ;


-- ----------------------------
-- 15、文章分类表
-- ----------------------------

CREATE TABLE IF NOT EXISTS `cms_category_article_rel` (
  `id` smallint(5) unsigned NOT NULL AUTO_INCREMENT COMMENT '栏目id',
  `pid` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '父id',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '类别 0:栏目 1:单网页',
  `name` varchar(30) NOT NULL DEFAULT '' COMMENT '栏目名称',
  `enname` varchar(30) NOT NULL DEFAULT '' COMMENT '栏目英文名称',
  `desc` mediumtext NOT NULL COMMENT '描述',
  `url` varchar(100) NOT NULL COMMENT '链接地址',
  `hits` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '点击数量',
  `setting` mediumtext NOT NULL COMMENT '栏目配置',
  `order` smallint(5) unsigned NOT NULL DEFAULT '0' COMMENT '排序',
  `is_menu` tinyint(1) unsigned NOT NULL DEFAULT '1' COMMENT '是否显示，1 显示',
  PRIMARY KEY (`id`),
  KEY `module` (`pid`,`order`,`id`),
  KEY `type` (`type`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='内容分类表' AUTO_INCREMENT=17 ;

-- ----------------------------
-- 16、文章分类关系表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `cms_category_article_rel` (
  `id` smallint(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '栏目id',
  `cid` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '栏目id category表相对应id',
  `aid` int(11) unsigned NOT NULL DEFAULT '0' COMMENT '文章id',
  `tid` int(5) unsigned NOT NULL DEFAULT '0' COMMENT '模板id',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '是否发布 0:不发布 1:发布',
  `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶 0:不置顶 1:置顶',
  PRIMARY KEY (`id`),
  KEY `module` (`cid`,`aid`),
   KEY `status` (`status`)
) ENGINE=InnoDB  DEFAULT CHARSET=utf8 COMMENT='内容分类关系表' AUTO_INCREMENT=17 ;

-- ----------------------------
-- 17、前端用户表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `cms_users` (
  `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_login` varchar(60) NOT NULL DEFAULT '' COMMENT '用户名',
  `user_pass` varchar(64) NOT NULL DEFAULT '' COMMENT '登录密码；sp_password加密',
  `user_nicename` varchar(50) NOT NULL DEFAULT '' COMMENT '用户美名',
  `user_email` varchar(100) NOT NULL DEFAULT '' COMMENT '登录邮箱',
  `user_url` varchar(100) NOT NULL DEFAULT '' COMMENT '用户个人网站',
  `avatar` varchar(255) DEFAULT NULL COMMENT '用户头像，相对于upload/avatar目录',
  `sex` smallint(1) DEFAULT '0' COMMENT '性别；0：保密，1：男；2：女',
  `birthday` date DEFAULT '2000-01-01' COMMENT '生日',
  `signature` varchar(255) DEFAULT NULL COMMENT '个性签名',
  `last_login_ip` varchar(16) DEFAULT NULL COMMENT '最后登录ip',
  `last_login_time` datetime DEFAULT NULL  COMMENT '最后登录时间',
  `validation_type`    varchar(50) comment '验证类型(用户激活,重置密码,邮箱激活)',
  `validation_key`    varchar(100) comment '验证KEY',
  `salt`               varchar(32) comment '加密混淆码',
  `user_status` int(11) NOT NULL DEFAULT '1' COMMENT '用户状态 0：禁用； 1：正常 ；2：未验证',
  `mobile` varchar(20) NOT NULL DEFAULT '' COMMENT '手机号',
  `qq_openid`          varchar(64) comment 'qq openid',
  `weibo_uid`          varchar(64) comment 'weibo uid',
  `weixin_openid`      varchar(64) comment 'weixin openid',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  unique KEY `user_login` (`user_login`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 COMMENT='用户表';

-- ----------------------------
-- 17、用户留言表
-- ----------------------------
CREATE TABLE IF NOT EXISTS `cms_guestbook` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `full_name` varchar(50) NOT NULL COMMENT '留言者姓名',
  `email` varchar(100) NOT NULL COMMENT '留言者邮箱',
  `title` varchar(255) DEFAULT NULL COMMENT '留言标题',
  `msg` text NOT NULL COMMENT '留言内容',
  `createtime` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '留言时间',
  `reply_msg` VARCHAR(255) NOT NULL DEFAULT '' comment '回复信息',
  `replier` VARCHAR(50) NOT NULL DEFAULT '' comment '答复者',
  `reply_time` timestamp DEFAULT  NULL comment '答复时间',
  `is_email_reply` SMALLINT(2) DEFAULT  '0'  COMMENT '是否同时发送答复邮件，1：发送，0：不发送',
  `status` smallint(2) NOT NULL DEFAULT '1' COMMENT '留言状态，1：正常，0：删除',
  PRIMARY KEY (`id`)
) ENGINE=innodb DEFAULT CHARSET=utf8 COMMENT='留言表';

-- ----------------------------
-- 18、友情链接
-- ----------------------------
CREATE TABLE `cms_link` (
  `id` int(5) NOT NULL AUTO_INCREMENT COMMENT '标识ID',
  `ftype` tinyint(1) NOT NULL DEFAULT '0' COMMENT '0:友情链接 1:合作单位',
  `title` varchar(30) NOT NULL DEFAULT '' COMMENT '标题',
  `url` varchar(150) NOT NULL DEFAULT '' COMMENT '链接地址',
  `cover` varchar(255) NOT NULL DEFAULT '0' COMMENT '封面图片ID',
  `descrip` varchar(255) NOT NULL DEFAULT '' COMMENT '备注信息',
  `sort` int(10) NOT NULL DEFAULT '0' COMMENT '排序',
  `hits` tinyint(7) NOT NULL DEFAULT '0' COMMENT '点击率',
  `update_time` int(10) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `uid` int(7) NOT NULL DEFAULT '0' COMMENT '用户ID ',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 0：禁用； 1：正常 ；',
  PRIMARY KEY (`id`),
  KEY `status` (`status`) USING BTREE
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COMMENT='友情链接';

-- ----------------------------
-- 19、广告表
-- ----------------------------
DROP TABLE IF EXISTS `cms_ad`;
CREATE TABLE `cms_ad` (
  `id` int(11) UNSIGNED  NOT NULL AUTO_INCREMENT COMMENT '主键',
  `place_id` int(11) DEFAULT NULL COMMENT '广告位ID',
  `title` varchar(150) DEFAULT NULL COMMENT '广告名称',
  `cover` VARCHAR(255) NOT NULL  COMMENT '广告图片',
  `photolist` varchar(20) NOT NULL COMMENT '辅助图片',
  `url` varchar(150) DEFAULT NULL COMMENT '广告链接',
  `listurl` varchar(255) DEFAULT NULL COMMENT '辅助链接',
  `background` varchar(150) DEFAULT NULL COMMENT '广告背景',
  `content` text COMMENT '广告描述',
  `begin_time` timestamp  DEFAULT NULL COMMENT '开始时间',
  `end_time` timestamp  DEFAULT NULL  COMMENT '结束时间',
  `create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `update_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '广告位状态0：禁用； 1：正常 ；',
  PRIMARY KEY(`id`),
  KEY `time_sq`(`begin_time`,`end_time`),
  KEY `status`(`status`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='广告表';

-- ----------------------------
-- 20、广告位
-- ----------------------------

DROP TABLE IF EXISTS `cms_ad_place`;
CREATE TABLE `cms_ad_place` (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `title` varchar(150) DEFAULT NULL COMMENT '广告位名称',
  `name` varchar(20) NOT NULL COMMENT '调用名称',
  `show_type` int(11) NOT NULL DEFAULT '5' COMMENT '广告位类型',
  `show_num` int(11) NOT NULL DEFAULT '5' COMMENT '显示条数',
  `start_time` int(11) NOT NULL DEFAULT '0' COMMENT '开始时间',
  `end_time` int(11) NOT NULL DEFAULT '0' COMMENT '结束时间',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '创建时间',
  `update_time` int(11) NOT NULL DEFAULT '0' COMMENT '更新时间',
  `template` varchar(150) DEFAULT NULL COMMENT '广告位模板',
  `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '广告位状态 0：禁用； 1：正常 ；',
   PRIMARY KEY(`id`),
   KEY `create_time`(`create_time`),
   KEY `status`(`status`)
) ENGINE=INNODB DEFAULT CHARSET=utf8 COMMENT='广告位表';

-- ----------------------------
-- 21、用户等级
-- ----------------------------
CREATE TABLE `cms_user_level` (
  `level_id` smallint(4) unsigned NOT NULL AUTO_INCREMENT COMMENT '表id',
  `level_name` varchar(30) DEFAULT NULL COMMENT '头衔名称',
  `sort` int(3) DEFAULT '0' COMMENT '排序',
  `bomlimit` int(5) DEFAULT '0' COMMENT '积分下限',
  `toplimit` int(5) DEFAULT '0' COMMENT '积分上限',
  PRIMARY KEY (`level_id`)
) ENGINE=INNODB AUTO_INCREMENT=10 DEFAULT CHARSET=utf8 COMMENT='用户等级';


-- ----------------------------
-- 22、用户收藏表
-- ----------------------------
DROP TABLE IF EXISTS `cms_user_favorite`;
CREATE TABLE `cms_user_favorite` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `uid` int(11) NOT NULL DEFAULT '0' COMMENT '收藏用户UID',
  `collect` int(11) NOT NULL COMMENT '收藏的内容ID编号',
  `type` tinyint(1) NOT NULL DEFAULT '0' COMMENT '1文章 2图书',
  `create_at` int(10) NOT NULL DEFAULT '0',
  PRIMARY KEY (`id`),
  KEY `uid`(`uid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='用户收藏表';

-- ----------------------------
-- 23、客户来源
-- ----------------------------
DROP TABLE IF EXISTS `crm_customer_source`;
CREATE TABLE `crm_customer_source` (
  `sid` int(11) NOT NULL AUTO_INCREMENT,
  `source` VARCHAR(20) NOT NULL  COMMENT '来源名称',
  `comment` text  COMMENT '备注',
  PRIMARY KEY (`sid`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8 COMMENT='客户来源表';


INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('1', '投资移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('2', '技术移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('3', '创业移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('4', '杰出人才', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('5', '留学生移民', '0');
INSERT INTO `oz_project_type` (`project_type_id`, `project_type_name`, `status`) VALUES ('6', '家属团聚', '0');


INSERT INTO `sys_admin_user` (`id`, `user_name`, `password`, `salt`, `sex`, `email`, `last_login`, `last_ip`, `status`,
`create_time`, `update_time`)
VALUES
	(1,'admin','21232f297a57a5a743894a0e4a801fc3','',1,'admin@admin.com','2016-05-11 10:33:49','127.0.0.1',0,'2016-05-11 10:33:49','2016-05-11 10:33:49');


ALTER TABLE `crm_customer`
ADD COLUMN `follow_status`  int(11) NULL DEFAULT 0 COMMENT '跟进状态: 0-未跟进，1-已跟进，2-无需跟进' AFTER `assign_to`;


INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('5', '营销QQ', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('6', '电话', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('7', '微信服务号', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('8', '微信订阅号', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('9', 'QQ群', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('10', '客户介绍', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('11', '同事介绍', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('12', '网站预约', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('13', '网站评估', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('14', '数据挖掘', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('15', '互动游戏', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('16', '项目宣传', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('17', '其他渠道', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('18', '移动咨询', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('19', '调研', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('20', '微信群', NULL);
INSERT INTO `crm_customer_source` (`sid`, `source`, `comment`) VALUES ('21', '百度搜索', NULL);
