CREATE DATABASE IF NOT EXISTS shipwar Character SET UTF8;
USE shipwar;
SET FOREIGN_KEY_CHECKS=0;

/* show create table xxxx; */

/* 自增id表 */
DROP TABLE IF EXISTS `autoid`;
CREATE TABLE `autoid` (
    `tablename` varchar(64) NOT NULL COMMENT "表名",
    `nextid` int NOT NULL COMMENT "获取下一个id",
    PRIMARY KEY (`tablename`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 游客账号表 */
DROP TABLE IF EXISTS `touristaccount`;
CREATE TABLE `touristaccount` (
    `account` varchar(64) NOT NULL COMMENT "游客账号",
    `password` varchar(64) NOT NULL COMMENT "游客密码",
    `createtime` int NOT NULL COMMENT "游客账号创建时间",
    `id` int NOT NULL COMMENT "游客唯一id",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

/* 用户基础信息表 */
DROP TABLE IF EXISTS `userinfo`;
CREATE TABLE `userinfo` (
    `id` int NOT NULL COMMENT "用户唯一id",
    `name` varchar(64) NOT NULL COMMENT "用户昵称",
    `sex` int NOT NULL COMMENT "用户性别",
    `headicon` varchar(64) NOT NULL COMMENT "用户头像",
    `gold` int NOT NULL COMMENT "金币",
    `exppool` int NOT NULL COMMENT "经验池",
    `level` int NOT NULL COMMENT "等级",
    `levelexp` int NOT NULL COMMENT "等级经验",
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
