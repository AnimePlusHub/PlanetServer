drop database if exists `PlanetMsg`;
create database `PlanetMsg` default character set utf8mb4 collate utf8mb4_unicode_ci;

use `PlanetMsg`;

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for friend
-- ----------------------------
DROP TABLE IF EXISTS `friend`;
CREATE TABLE `friend` (
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `friend_id` int(11) NOT NULL COMMENT '好友id',
  `group_id` int(11) NOT NULL COMMENT '分组id',
  `remark` varchar(20) NOT NULL COMMENT '好友备注',
  `create_time` date NOT NULL COMMENT '添加时间',
  `add_way` varchar(20) NOT NULL COMMENT '添加方式',
  PRIMARY KEY (`user_id`,`friend_id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  KEY `friend_id` (`friend_id`) USING BTREE,
  KEY `group_id` (`group_id`) USING BTREE,
  CONSTRAINT `friend_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `friend_ibfk_2` FOREIGN KEY (`friend_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  CONSTRAINT `friend_ibfk_3` FOREIGN KEY (`group_id`) REFERENCES `friend_group` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for friend_group
-- ----------------------------
DROP TABLE IF EXISTS `friend_group`;
CREATE TABLE `friend_group` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '分组id',
  `user_id` int(11) NOT NULL COMMENT '用户id',
  `group_name` varchar(20) NOT NULL COMMENT '分组名称(我的好友、朋友。。。)',
  PRIMARY KEY (`id`) USING BTREE,
  KEY `user_id` (`user_id`) USING BTREE,
  CONSTRAINT `friend_group_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '用户id主键',
  `account` varchar(20) NOT NULL COMMENT '用户账号',
  `nick_name` varchar(20) NOT NULL COMMENT '用户昵称',
  `pwd` varchar(72) NOT NULL COMMENT '用户密码',
  `email` varchar(25) DEFAULT NULL COMMENT '邮箱',
  `phone` char(11) DEFAULT NULL COMMENT '手机号',
  `pic_url` varchar(72) NOT NULL COMMENT '用户头像',
  `birth_day` date DEFAULT NULL COMMENT '生日',
  `create_time` date NOT NULL COMMENT '注册时间',
  `status` tinyint(4) DEFAULT NULL COMMENT '用户状态（在线、离线。。）',
  `last_login_time` datetime DEFAULT NULL COMMENT '最后登陆时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

SET FOREIGN_KEY_CHECKS = 1;
