/*
 Navicat Premium Data Transfer

 Source Server         : dev
 Source Server Type    : MySQL
 Source Server Version : 80200
 Source Host           : localhost:3306
 Source Schema         : party_affairs

 Target Server Type    : MySQL
 Target Server Version : 80200
 File Encoding         : 65001

 Date: 24/03/2024 17:45:04
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for banner
-- ----------------------------
DROP TABLE IF EXISTS `banner`;
CREATE TABLE `banner` (
  `title` varchar(128) NOT NULL COMMENT '轮播标题',
  `src` varchar(128) NOT NULL COMMENT '轮播图',
  `app` varchar(32) DEFAULT NULL COMMENT '所属应用',
  `url` varchar(128) DEFAULT NULL COMMENT '跳转链接',
  `params` varchar(128) DEFAULT NULL COMMENT '跳转参数',
  `weight` int unsigned DEFAULT NULL COMMENT '轮播权重',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `banner` varchar(32) DEFAULT NULL COMMENT '所属应用',
  `path` varchar(256) DEFAULT NULL COMMENT '跳转链接',
  PRIMARY KEY (`id`),
  KEY `idx_banner_created_at` (`created_at`),
  KEY `idx_banner_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4  COMMENT='轮播';

-- ----------------------------
-- Records of banner
-- ----------------------------
BEGIN;
INSERT INTO `banner` VALUES ('春华秋月何时了', 'e526022fed3eb9fc53ad18c7c18f1fb036aea4556ee4e45838e8ca840f545316', '', '', '', 0, 1, 1706944965, 1707115877, NULL, 'https://www.baidu.com?id=1');
INSERT INTO `banner` VALUES ('往事知多少', 'e526022fed3eb9fc53ad18c7c18f1fb036aea4556ee4e45838e8ca840f545316', '', '', '', 0, 2, 1706944991, 1707115691, NULL, 'https://www.baidu.com/');
COMMIT;

-- ----------------------------
-- Table structure for news
-- ----------------------------
DROP TABLE IF EXISTS `news`;
CREATE TABLE `news` (
  `title` varchar(128) NOT NULL COMMENT '新闻标题',
  `unit` varchar(128) NOT NULL COMMENT '发布单位',
  `cover` varchar(128) NOT NULL COMMENT '封面图片',
  `desc` varchar(128) NOT NULL COMMENT '新闻描述',
  `content` blob NOT NULL COMMENT '新闻内容',
  `read` int unsigned DEFAULT NULL COMMENT '阅读人数',
  `classify_id` int unsigned NOT NULL COMMENT '新闻分类',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_news_created_at` (`created_at`),
  KEY `idx_news_updated_at` (`updated_at`),
  KEY `fk_news_news_classify` (`classify_id`),
  CONSTRAINT `fk_news_news_classify` FOREIGN KEY (`classify_id`) REFERENCES `news_classify` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='新闻';

-- ----------------------------
-- Records of news
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for news_classify
-- ----------------------------
DROP TABLE IF EXISTS `news_classify`;
CREATE TABLE `news_classify` (
  `name` varchar(128) NOT NULL COMMENT '分类名称',
  `weight` int unsigned DEFAULT NULL COMMENT '分类权重',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_news_classify_created_at` (`created_at`),
  KEY `idx_news_classify_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='新闻分类';

-- ----------------------------
-- Records of news_classify
-- ----------------------------
BEGIN;
INSERT INTO `news_classify` VALUES ('党务新闻', 0, 1, 1706960554, 1706960554);
COMMIT;

-- ----------------------------
-- Table structure for news_comment
-- ----------------------------
DROP TABLE IF EXISTS `news_comment`;
CREATE TABLE `news_comment` (
  `content_id` int unsigned NOT NULL COMMENT '新闻ID',
  `from_uid` int unsigned NOT NULL COMMENT '发布者ID',
  `reply_uid` int unsigned DEFAULT NULL COMMENT '回复者ID',
  `text` varchar(1024) DEFAULT NULL COMMENT '回复文本',
  `images` varchar(1024) DEFAULT NULL COMMENT '回复图片',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  KEY `idx_news_comment_created_at` (`created_at`),
  KEY `fk_news_content_comments` (`content_id`),
  CONSTRAINT `fk_news_content_comments` FOREIGN KEY (`content_id`) REFERENCES `news_content` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=18 DEFAULT CHARSET=utf8mb4  COMMENT='新闻评论';

-- ----------------------------
-- Records of news_comment
-- ----------------------------
BEGIN;
INSERT INTO `news_comment` VALUES (1, 5, 0, '<p data-v-513852e0=\"\">春花秋月何时了，往事知多少，小楼左右又东风，故国不堪回首月明中</p>', NULL, 5, 1707100417);
INSERT INTO `news_comment` VALUES (1, 5, 0, '<p data-v-513852e0=\"\">春花秋月何时了，往事知多少，小楼左右又东风，故国不堪回首月明中</p><p data-v-513852e0=\"\">春花秋月何时了，往事知多少，小楼左右又东风，故国不堪回首月明中</p><p data-v-513852e0=\"\">春花秋月何时了，往事知多少，小楼左右又东风，故国不堪回首月明中</p>', NULL, 6, 1707100430);
INSERT INTO `news_comment` VALUES (1, 6, 5, '你吹牛逼', NULL, 10, 1707103210);
INSERT INTO `news_comment` VALUES (1, 6, 5, '你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼', NULL, 11, 1707103413);
INSERT INTO `news_comment` VALUES (1, 6, 5, '你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼你吹牛逼', NULL, 12, 1707103419);
COMMIT;

-- ----------------------------
-- Table structure for news_content
-- ----------------------------
DROP TABLE IF EXISTS `news_content`;
CREATE TABLE `news_content` (
  `title` varchar(128) NOT NULL COMMENT '新闻标题',
  `unit` varchar(128) NOT NULL COMMENT '发布单位',
  `cover` varchar(128) NOT NULL COMMENT '封面图片',
  `desc` varchar(128) NOT NULL COMMENT '新闻描述',
  `content` blob NOT NULL COMMENT '新闻内容',
  `read` int unsigned DEFAULT NULL COMMENT '阅读人数',
  `classify_id` int unsigned NOT NULL COMMENT '新闻分类',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `is_top` tinyint(1) DEFAULT NULL COMMENT '是否置顶',
  PRIMARY KEY (`id`),
  KEY `idx_news_content_created_at` (`created_at`),
  KEY `idx_news_content_updated_at` (`updated_at`),
  KEY `fk_news_content_classify` (`classify_id`),
  CONSTRAINT `fk_news_content_classify` FOREIGN KEY (`classify_id`) REFERENCES `news_classify` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='新闻内容';

-- ----------------------------
-- Records of news_content
-- ----------------------------
BEGIN;
INSERT INTO `news_content` VALUES ('春花秋月何时了，往事知多少，小楼左右又东风，故国不堪回首月明中', '中华新闻', '18f7421664c6f34daccf3954bca821fed209017e3ea0163495d744f6b1e06a66', '春花秋月何时了，往事知多少，小楼左右又东风，故国不堪回首月明中', 0x3C703EE698A5E88AB1E7A78BE69C88E4BD95E697B6E4BA86EFBC8CE5BE80E4BA8BE79FA5E5A49AE5B091EFBC8CE5B08FE6A5BCE5B7A6E58FB3E58F88E4B89CE9A38EEFBC8CE69585E59BBDE4B88DE5A0AAE59B9EE9A696E69C88E6988EE4B8AD3C2F703E, 44, 1, 1, 1707035332, 1707047652, 1);
COMMIT;

-- ----------------------------
-- Table structure for notice
-- ----------------------------
DROP TABLE IF EXISTS `notice`;
CREATE TABLE `notice` (
  `title` varchar(128) NOT NULL COMMENT '公告标题',
  `desc` varchar(512) NOT NULL COMMENT '公告描述',
  `unit` varchar(128) NOT NULL COMMENT '发布单位',
  `content` text NOT NULL COMMENT '轮播权重',
  `is_top` tinyint(1) DEFAULT NULL COMMENT '是否置顶',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_notice_created_at` (`created_at`),
  KEY `idx_notice_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='通知';

-- ----------------------------
-- Records of notice
-- ----------------------------
BEGIN;
INSERT INTO `notice` VALUES ('12', '中华新闻网中华新闻网中华新闻网中华新闻网', '中华新闻网', '<p>中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网中华新闻网</p>', 0, 1, 1707132637, 1707994911);
COMMIT;

-- ----------------------------
-- Table structure for notice_user
-- ----------------------------
DROP TABLE IF EXISTS `notice_user`;
CREATE TABLE `notice_user` (
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `notice_id` int unsigned NOT NULL COMMENT '通知id',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `un` (`user_id`,`notice_id`),
  KEY `idx_notice_user_created_at` (`created_at`),
  KEY `fk_notice_notice_user` (`notice_id`),
  CONSTRAINT `fk_notice_notice_user` FOREIGN KEY (`notice_id`) REFERENCES `notice` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=6 DEFAULT CHARSET=utf8mb4  COMMENT='通知用户';

-- ----------------------------
-- Records of notice_user
-- ----------------------------
BEGIN;
INSERT INTO `notice_user` VALUES (5, 1, 1, 1707203687);
COMMIT;

-- ----------------------------
-- Table structure for resource
-- ----------------------------
DROP TABLE IF EXISTS `resource`;
CREATE TABLE `resource` (
  `title` varchar(128) NOT NULL COMMENT '文档标题',
  `desc` varchar(256) NOT NULL COMMENT '文档描述',
  `url` varchar(256) NOT NULL COMMENT '文档url',
  `download_count` int unsigned DEFAULT NULL COMMENT '下载次数',
  `classify_id` int unsigned NOT NULL COMMENT '新闻分类',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_resource_created_at` (`created_at`),
  KEY `idx_resource_updated_at` (`updated_at`),
  KEY `fk_resource_resource_classify` (`classify_id`),
  CONSTRAINT `fk_resource_resource_classify` FOREIGN KEY (`classify_id`) REFERENCES `resource_classify` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4  COMMENT='资料';

-- ----------------------------
-- Records of resource
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for resource_classify
-- ----------------------------
DROP TABLE IF EXISTS `resource_classify`;
CREATE TABLE `resource_classify` (
  `name` varchar(128) NOT NULL COMMENT '分类名称',
  `weight` int unsigned DEFAULT NULL COMMENT '分类权重',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_resource_classify_created_at` (`created_at`),
  KEY `idx_resource_classify_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='资料分类';

-- ----------------------------
-- Records of resource_classify
-- ----------------------------
BEGIN;
INSERT INTO `resource_classify` VALUES ('学习', 0, 1, 1707995357, 1707995357);
COMMIT;

-- ----------------------------
-- Table structure for resource_content
-- ----------------------------
DROP TABLE IF EXISTS `resource_content`;
CREATE TABLE `resource_content` (
  `title` varchar(128) NOT NULL COMMENT '文档标题',
  `desc` varchar(256) NOT NULL COMMENT '文档描述',
  `url` varchar(256) NOT NULL COMMENT '文档url',
  `download_count` int unsigned DEFAULT NULL COMMENT '下载次数',
  `classify_id` int unsigned NOT NULL COMMENT '资源分类',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_resource_content_created_at` (`created_at`),
  KEY `idx_resource_content_updated_at` (`updated_at`),
  KEY `fk_resource_content_resource_classify` (`classify_id`),
  CONSTRAINT `fk_resource_content_resource_classify` FOREIGN KEY (`classify_id`) REFERENCES `resource_classify` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='资料';

-- ----------------------------
-- Records of resource_content
-- ----------------------------
BEGIN;
INSERT INTO `resource_content` VALUES ('1', '撒', '8d4a40044b088f91f99e8ab42756021f3795ebf5677a6c85b9a5237e042e3e4f', 0, 1, 1, 1707995520, 1707995520);
COMMIT;

-- ----------------------------
-- Table structure for task
-- ----------------------------
DROP TABLE IF EXISTS `task`;
CREATE TABLE `task` (
  `title` varchar(128) NOT NULL COMMENT '任务标题',
  `desc` varchar(256) NOT NULL COMMENT '任务公告',
  `is_update` tinyint(1) DEFAULT NULL COMMENT '是否可更新',
  `start` int unsigned NOT NULL COMMENT '开始时间',
  `end` int unsigned NOT NULL COMMENT '结束时间',
  `config` text NOT NULL COMMENT '任务配置',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  KEY `idx_task_created_at` (`created_at`),
  KEY `idx_task_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='任务';

-- ----------------------------
-- Records of task
-- ----------------------------
BEGIN;
INSERT INTO `task` VALUES ('1', '是', 1, 1706796649, 1709216034, '[{\"show\":true,\"type\":\"input\",\"field\":\"e0d255c6\",\"config\":{\"label\":\"标题\",\"require\":true,\"disabled\":false,\"max_length\":30,\"value\":\"\",\"placeholder\":\"请输入\",\"options\":[],\"required\":true}},{\"show\":true,\"type\":\"select\",\"field\":\"f2521264\",\"config\":{\"label\":\"标题\",\"required\":true,\"disabled\":false,\"value\":\"\",\"placeholder\":\"请输入\",\"max_length\":30,\"options\":[\"选择一\",\"选择二\"]}}]', 1, 1708006259, 1708355480);
COMMIT;

-- ----------------------------
-- Table structure for task_value
-- ----------------------------
DROP TABLE IF EXISTS `task_value`;
CREATE TABLE `task_value` (
  `task_id` int unsigned NOT NULL COMMENT '任务id',
  `user_id` int unsigned NOT NULL COMMENT '用户id',
  `value` text NOT NULL COMMENT '数据值',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `tu` (`task_id`,`user_id`),
  KEY `idx_task_value_created_at` (`created_at`),
  KEY `idx_task_value_updated_at` (`updated_at`),
  CONSTRAINT `fk_task_task_value` FOREIGN KEY (`task_id`) REFERENCES `task` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4  COMMENT='任务值';

-- ----------------------------
-- Records of task_value
-- ----------------------------
BEGIN;
INSERT INTO `task_value` VALUES (1, 5, '{\"e0d255c6\":\"2\",\"f2521264\":\"选择一\"}', 2, 1708354034, 1708356109);
COMMIT;

-- ----------------------------
-- Table structure for user_video_process
-- ----------------------------
DROP TABLE IF EXISTS `user_video_process`;
CREATE TABLE `user_video_process` (
  `video_id` int unsigned NOT NULL COMMENT '视频ID',
  `user_id` int unsigned NOT NULL COMMENT '用户ID',
  `finish` tinyint(1) DEFAULT '0' COMMENT '是否完成',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `time` int unsigned NOT NULL COMMENT '当前进度时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `vu` (`video_id`,`user_id`),
  KEY `idx_user_video_process_created_at` (`created_at`),
  KEY `idx_user_video_process_updated_at` (`updated_at`),
  CONSTRAINT `fk_video_content_process` FOREIGN KEY (`video_id`) REFERENCES `video_content` (`id`) ON DELETE CASCADE
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4  COMMENT='用户视频进度';

-- ----------------------------
-- Records of user_video_process
-- ----------------------------
BEGIN;
INSERT INTO `user_video_process` VALUES (1, 5, 0, 1, 1708270513, 1708356642, 75);
INSERT INTO `user_video_process` VALUES (2, 5, 1, 2, 1708271734, 1708272200, 125);
INSERT INTO `user_video_process` VALUES (1, 3, 1, 3, 1708848981, 1708850223, 125);
COMMIT;

-- ----------------------------
-- Table structure for video_classify
-- ----------------------------
DROP TABLE IF EXISTS `video_classify`;
CREATE TABLE `video_classify` (
  `name` varchar(128) NOT NULL COMMENT '分类名称',
  `cover` varchar(128) NOT NULL COMMENT '分类封面',
  `weight` int unsigned DEFAULT NULL COMMENT '分类权重',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `is_top` tinyint(1) DEFAULT NULL COMMENT '是否指定',
  `is_task` tinyint(1) DEFAULT NULL COMMENT '是否任务',
  `start` int unsigned DEFAULT NULL COMMENT '任务开始时间',
  `end` int unsigned DEFAULT NULL COMMENT '任务结束时间',
  `desc` varchar(256) NOT NULL COMMENT '分类描述',
  PRIMARY KEY (`id`),
  KEY `idx_video_classify_created_at` (`created_at`),
  KEY `idx_video_classify_updated_at` (`updated_at`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4  COMMENT='视频分类';

-- ----------------------------
-- Records of video_classify
-- ----------------------------
BEGIN;
INSERT INTO `video_classify` VALUES ('这是一个视频描述，讲述了分类内容', '18f7421664c6f34daccf3954bca821fed209017e3ea0163495d744f6b1e06a66', 1, 1, 1708159296, 1708268178, 0, 1, 1708196403, 1708588696, '这是一个视频描述，讲述了分类内容这是一个视频描述，讲述了分类内容');
COMMIT;

-- ----------------------------
-- Table structure for video_content
-- ----------------------------
DROP TABLE IF EXISTS `video_content`;
CREATE TABLE `video_content` (
  `title` varchar(128) NOT NULL COMMENT '视频标题',
  `desc` varchar(256) NOT NULL COMMENT '视频描述',
  `url` varchar(256) NOT NULL COMMENT '视频url',
  `classify_id` int unsigned NOT NULL COMMENT '资源分类',
  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `created_at` bigint DEFAULT NULL COMMENT '创建时间',
  `updated_at` bigint DEFAULT NULL COMMENT '修改时间',
  `weight` int unsigned DEFAULT NULL COMMENT '序号',
  `duration` int unsigned NOT NULL COMMENT '持续时间',
  PRIMARY KEY (`id`),
  KEY `idx_video_content_created_at` (`created_at`),
  KEY `idx_video_content_updated_at` (`updated_at`),
  KEY `fk_video_content_classify` (`classify_id`),
  CONSTRAINT `fk_video_content_classify` FOREIGN KEY (`classify_id`) REFERENCES `video_classify` (`id`),
  CONSTRAINT `fk_video_content_video_classify` FOREIGN KEY (`classify_id`) REFERENCES `video_classify` (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=3 DEFAULT CHARSET=utf8mb4  COMMENT='视频内容';

-- ----------------------------
-- Records of video_content
-- ----------------------------
BEGIN;
INSERT INTO `video_content` VALUES ('sa', '新增视频', 'f3fd49374ff69fc8f79efcd312fb45b9078e37229cc56c33cfc20e3fe1b78b58', 1, 1, 1708239402, 1708253220, 0, 125);
INSERT INTO `video_content` VALUES ('测试是哦额', '测试是哦额测试是哦额', 'f3fd49374ff69fc8f79efcd312fb45b9078e37229cc56c33cfc20e3fe1b78b58', 1, 2, 1708271718, 1708271718, 0, 125);
COMMIT;

SET FOREIGN_KEY_CHECKS = 1;
