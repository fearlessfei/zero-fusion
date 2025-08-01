CREATE TABLE `cron_config` (
    `id` bigint NOT NULL AUTO_INCREMENT,
    `cron_spec` varchar(30) NOT NULL DEFAULT '' COMMENT 'cron表达式',
    `task_type` varchar(50) NOT NULL DEFAULT '' COMMENT '任务类型',
    `payload` varchar(255) NOT NULL DEFAULT '' COMMENT '任务参数',
    `task_options` varchar(255) DEFAULT NULL COMMENT '任务选项',
    `created_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_task_type` (`task_type`) USING BTREE
) ENGINE=InnoDB  DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;