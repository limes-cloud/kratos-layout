SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `party_affairs`
--

-- --------------------------------------------------------

--
-- 表的结构 `banner`
--

CREATE TABLE `banner` (
                          `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                          `title` varchar(128) NOT NULL COMMENT '轮播标题',
                          `src` varchar(128) NOT NULL COMMENT '轮播封面',
                          `path` varchar(128) DEFAULT NULL COMMENT '跳转链接',
                          `weight` int(11) DEFAULT NULL COMMENT '轮播权重',
                          `status` tinyint(1) DEFAULT '0' COMMENT '轮播状态',
                          `created_at` bigint(20) DEFAULT '0' COMMENT '创建时间',
                          `updated_at` bigint(20) DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='轮播';

-- --------------------------------------------------------

--
-- 表的结构 `gorm_init`
--

CREATE TABLE `gorm_init` (
                             `id` int(10) UNSIGNED NOT NULL,
                             `init` tinyint(1) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `gorm_init`
--

INSERT INTO `gorm_init` (`id`, `init`) VALUES
    (1, 1);

-- --------------------------------------------------------

--
-- 表的结构 `information`
--

CREATE TABLE `information` (
                               `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                               `classify_id` bigint(20) UNSIGNED NOT NULL COMMENT '分类id',
                               `title` varchar(128) NOT NULL COMMENT '咨询标题',
                               `cover` varchar(128) NOT NULL COMMENT '咨询封面',
                               `description` varchar(512) NOT NULL COMMENT '咨询描述',
                               `unit` varchar(128) NOT NULL COMMENT '发布单位',
                               `content` longtext NOT NULL COMMENT '咨询内容',
                               `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
                               `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '咨询状态',
                               `read` bigint(20) NOT NULL DEFAULT '0' COMMENT '阅读人数',
                               `created_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
                               `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='咨询';

-- --------------------------------------------------------

--
-- 表的结构 `information_classify`
--

CREATE TABLE `information_classify` (
                                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                        `name` varchar(32) NOT NULL COMMENT '名称',
                                        `weight` int(11) DEFAULT '1' COMMENT '权重',
                                        `created_at` bigint(20) DEFAULT NULL COMMENT '创建时间',
                                        `updated_at` bigint(20) DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资讯分组';

--
-- 转存表中的数据 `information_classify`
--

INSERT INTO `information_classify` (`id`, `name`, `weight`, `created_at`, `updated_at`) VALUES
                                                                                            (1, '时代楷模', 0, 1710060983, 1710061243),
                                                                                            (2, '教育党员', 0, 1711955552, 1711955552),
                                                                                            (3, '风采展示', 0, 1726823660, 1726823660),
                                                                                            (4, '学习问答', 0, 1727260887, 1727260887);

-- --------------------------------------------------------

--
-- 表的结构 `notice`
--

CREATE TABLE `notice` (
                          `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                          `title` varchar(128) NOT NULL COMMENT '通知标题',
                          `description` varchar(512) NOT NULL COMMENT '通知描述',
                          `unit` varchar(128) NOT NULL COMMENT '通知单位',
                          `content` longtext NOT NULL COMMENT '通知内容',
                          `is_top` tinyint(1) NOT NULL DEFAULT '0' COMMENT '是否置顶',
                          `status` tinyint(1) NOT NULL DEFAULT '0' COMMENT '通知状态',
                          `created_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
                          `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通知';


-- --------------------------------------------------------

--
-- 表的结构 `notice_user`
--

CREATE TABLE `notice_user` (
                               `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                               `notice_id` bigint(20) UNSIGNED NOT NULL COMMENT '通知id',
                               `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
                               `is_read` tinyint(1) NOT NULL COMMENT '是否阅读',
                               `created_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '创建时间',
                               `updated_at` bigint(20) NOT NULL DEFAULT '0' COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='通知用户';

-- --------------------------------------------------------

--
-- 表的结构 `resource`
--

CREATE TABLE `resource` (
                            `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                            `title` varchar(128) NOT NULL COMMENT '资料标题',
                            `description` varchar(256) NOT NULL COMMENT '资料描述',
                            `url` varchar(256) NOT NULL COMMENT '资料url',
                            `download_count` int(10) UNSIGNED DEFAULT NULL COMMENT '下载次数',
                            `classify_id` bigint(20) UNSIGNED NOT NULL COMMENT '资源分类',
                            `created_at` bigint(20) DEFAULT NULL COMMENT '创建时间',
                            `updated_at` bigint(20) DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资料';


--
-- 表的结构 `resource_classify`
--

CREATE TABLE `resource_classify` (
                                     `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                                     `name` varchar(32) NOT NULL COMMENT '名称',
                                     `weight` int(11) DEFAULT '1' COMMENT '权重',
                                     `created_at` bigint(20) DEFAULT NULL COMMENT '创建时间',
                                     `updated_at` bigint(20) DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='资料分组';


--
-- 表的结构 `task`
--

CREATE TABLE `task` (
                        `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                        `title` varchar(128) NOT NULL COMMENT '任务标题',
                        `description` varchar(256) NOT NULL COMMENT '任务公告',
                        `is_update` tinyint(1) DEFAULT NULL COMMENT '是否可更新',
                        `start` bigint(20) UNSIGNED NOT NULL COMMENT '开始时间',
                        `end` bigint(20) UNSIGNED NOT NULL COMMENT '结束时间',
                        `config` text NOT NULL COMMENT '任务配置',
                        `created_at` bigint(20) DEFAULT NULL COMMENT '创建时间',
                        `updated_at` bigint(20) DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务';


--
-- 表的结构 `task_value`
--

CREATE TABLE `task_value` (
                              `id` bigint(20) UNSIGNED NOT NULL COMMENT '主键ID',
                              `user_id` bigint(20) UNSIGNED NOT NULL COMMENT '用户id',
                              `task_id` bigint(20) UNSIGNED NOT NULL COMMENT '任务id',
                              `value` text NOT NULL COMMENT '数据值',
                              `created_at` bigint(20) DEFAULT NULL COMMENT '创建时间',
                              `updated_at` bigint(20) DEFAULT NULL COMMENT '修改时间'
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='任务值';


--
-- 转储表的索引
--

--
-- 表的索引 `banner`
--
ALTER TABLE `banner`
    ADD PRIMARY KEY (`id`),
  ADD KEY `idx_banner_created_at` (`created_at`),
  ADD KEY `idx_banner_updated_at` (`updated_at`);

--
-- 表的索引 `gorm_init`
--
ALTER TABLE `gorm_init`
    ADD PRIMARY KEY (`id`);

--
-- 表的索引 `information`
--
ALTER TABLE `information`
    ADD PRIMARY KEY (`id`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `updated_at` (`updated_at`),
  ADD KEY `classify_id` (`classify_id`);

--
-- 表的索引 `information_classify`
--
ALTER TABLE `information_classify`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`),
  ADD KEY `updated_at` (`updated_at`),
  ADD KEY `created_at` (`created_at`);

--
-- 表的索引 `notice`
--
ALTER TABLE `notice`
    ADD PRIMARY KEY (`id`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `updated_at` (`updated_at`);

--
-- 表的索引 `notice_user`
--
ALTER TABLE `notice_user`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `notice_id` (`notice_id`,`user_id`);

--
-- 表的索引 `resource`
--
ALTER TABLE `resource`
    ADD PRIMARY KEY (`id`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `updated_at` (`updated_at`),
  ADD KEY `classify_id` (`classify_id`);

--
-- 表的索引 `resource_classify`
--
ALTER TABLE `resource_classify`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`),
  ADD KEY `updated_at` (`updated_at`),
  ADD KEY `created_at` (`created_at`);

--
-- 表的索引 `task`
--
ALTER TABLE `task`
    ADD PRIMARY KEY (`id`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `updated_at` (`updated_at`);

--
-- 表的索引 `task_value`
--
ALTER TABLE `task_value`
    ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `task_id` (`task_id`,`user_id`),
  ADD KEY `created_at` (`created_at`),
  ADD KEY `updated_at` (`updated_at`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `banner`
--
ALTER TABLE `banner`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `gorm_init`
--
ALTER TABLE `gorm_init`
    MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `information`
--
ALTER TABLE `information`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=33;

--
-- 使用表AUTO_INCREMENT `information_classify`
--
ALTER TABLE `information_classify`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=5;

--
-- 使用表AUTO_INCREMENT `notice`
--
ALTER TABLE `notice`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=2;

--
-- 使用表AUTO_INCREMENT `notice_user`
--
ALTER TABLE `notice_user`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID';

--
-- 使用表AUTO_INCREMENT `resource`
--
ALTER TABLE `resource`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=13;

--
-- 使用表AUTO_INCREMENT `resource_classify`
--
ALTER TABLE `resource_classify`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `task`
--
ALTER TABLE `task`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=6;

--
-- 使用表AUTO_INCREMENT `task_value`
--
ALTER TABLE `task_value`
    MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键ID', AUTO_INCREMENT=72;

--
-- 限制导出的表
--

--
-- 限制表 `information`
--
ALTER TABLE `information`
    ADD CONSTRAINT `information_ibfk_1` FOREIGN KEY (`classify_id`) REFERENCES `information_classify` (`id`);

--
-- 限制表 `notice_user`
--
ALTER TABLE `notice_user`
    ADD CONSTRAINT `notice_user_ibfk_1` FOREIGN KEY (`notice_id`) REFERENCES `notice` (`id`) ON DELETE CASCADE;

--
-- 限制表 `resource`
--
ALTER TABLE `resource`
    ADD CONSTRAINT `resource_ibfk_1` FOREIGN KEY (`classify_id`) REFERENCES `resource_classify` (`id`);

--
-- 限制表 `task_value`
--
ALTER TABLE `task_value`
    ADD CONSTRAINT `task_value_ibfk_1` FOREIGN KEY (`task_id`) REFERENCES `task` (`id`) ON DELETE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
