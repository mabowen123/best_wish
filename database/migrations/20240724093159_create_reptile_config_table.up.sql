CREATE TABLE reptile_config (
   `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
   `url` varchar(255) NOT NULL DEFAULT '' COMMENT '请求地址',
   `interval_time` int(10) unsigned NOT NULL DEFAULT '0' COMMENT '间隔时间(秒',
   `type` tinyint(3) unsigned NOT NULL DEFAULT '0' COMMENT '地址类型 1-线报',
   `next_time` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '下次执行时间',
   `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
   `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间戳',
   `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间戳',
   PRIMARY KEY (`id`),
   KEY `idx_nextTime` (`next_time`)
) COMMENT="线报爬虫配置";
