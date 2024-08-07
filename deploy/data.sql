/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 80300
 Source Host           : localhost:3306
 Source Schema         : poverty

 Target Server Type    : MySQL
 Target Server Version : 80300
 File Encoding         : 65001

 Date: 07/08/2024 13:58:05
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for activity
-- ----------------------------
DROP TABLE IF EXISTS `activity`;
CREATE TABLE `activity` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `title` varchar(128) NOT NULL COMMENT '活动标题',
                            `cover` varchar(128) NOT NULL COMMENT '活动封面',
                            `description` varchar(512) NOT NULL COMMENT '活动描述',
                            `unit` varchar(128) NOT NULL COMMENT '发布单位',
                            `content` longtext NOT NULL COMMENT '活动内容',
                            `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
                            `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '活动状态',
                            `read` bigint NOT NULL DEFAULT '0' COMMENT '阅读人数',
                            `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
                            `updated_at` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
                            PRIMARY KEY (`id`),
                            KEY `created_at` (`created_at`),
                            KEY `updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='活动';

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                          `title` varchar(128) NOT NULL COMMENT '轮播标题',
                          `src` varchar(128) NOT NULL COMMENT '轮播封面',
                          `path` varchar(128) DEFAULT NULL COMMENT '跳转链接',
                          `weight` int DEFAULT NULL COMMENT '轮播权重',
                          `status` tinyint(1) DEFAULT '0' COMMENT '轮播状态',
                          `created_at` bigint DEFAULT '0' COMMENT '创建时间',
                          `updated_at` bigint DEFAULT '0' COMMENT '修改时间',
                          PRIMARY KEY (`id`),
                          KEY `idx_banner_created_at` (`created_at`),
                          KEY `idx_banner_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='轮播';

-- ----------------------------
-- Table structure for chat_record
-- ----------------------------
DROP TABLE IF EXISTS `chat_record`;
CREATE TABLE `chat_record` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                               `user_id` bigint NOT NULL COMMENT '用户id',
                               `session_id` varchar(128) NOT NULL COMMENT '会话id',
                               `message` text NOT NULL COMMENT '会话信息',
                               `type` char(12) NOT NULL COMMENT '信息类型',
                               `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
                               PRIMARY KEY (`id`),
                               KEY `created_at` (`created_at`),
                               KEY `session_id` (`session_id`,`created_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='会话记录';

-- ----------------------------
-- Table structure for gorm_init
-- ----------------------------
DROP TABLE IF EXISTS `gorm_init`;
CREATE TABLE `gorm_init` (
                             `id` int unsigned NOT NULL AUTO_INCREMENT,
                             `init` tinyint(1) DEFAULT NULL,
                             PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4 ;

-- ----------------------------
-- Table structure for information
-- ----------------------------
DROP TABLE IF EXISTS `information`;
CREATE TABLE `information` (
                               `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                               `title` varchar(128) NOT NULL COMMENT '咨询标题',
                               `cover` varchar(128) NOT NULL COMMENT '咨询封面',
                               `description` varchar(512) NOT NULL COMMENT '咨询描述',
                               `unit` varchar(128) NOT NULL COMMENT '发布单位',
                               `content` longtext NOT NULL COMMENT '咨询内容',
                               `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
                               `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '咨询状态',
                               `read` bigint NOT NULL DEFAULT '0' COMMENT '阅读人数',
                               `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
                               `updated_at` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
                               PRIMARY KEY (`id`),
                               KEY `created_at` (`created_at`),
                               KEY `updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='咨询';

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice` (
                          `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                          `title` varchar(128) NOT NULL COMMENT '通知标题',
                          `description` varchar(512) NOT NULL COMMENT '通知描述',
                          `unit` varchar(128) NOT NULL COMMENT '通知单位',
                          `content` longtext NOT NULL COMMENT '通知内容',
                          `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
                          `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '通知状态',
                          `created_at` bigint NOT NULL DEFAULT '0' COMMENT '创建时间',
                          `updated_at` bigint NOT NULL DEFAULT '0' COMMENT '修改时间',
                          PRIMARY KEY (`id`),
                          KEY `idx_notice_created_at` (`created_at`),
                          KEY `idx_notice_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='通知';

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
                            `id` bigint unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                            `title` varchar(128) NOT NULL COMMENT '文件标题',
                            `src` varchar(128) NOT NULL COMMENT '文件地址',
                            `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '文件状态',
                            `created_at` bigint DEFAULT '0' COMMENT '创建时间',
                            `updated_at` bigint DEFAULT '0' COMMENT '修改时间',
                            PRIMARY KEY (`id`),
                            KEY `created_at` (`created_at`),
                            KEY `updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8mb4  COMMENT='文件';

SET FOREIGN_KEY_CHECKS = 1;
