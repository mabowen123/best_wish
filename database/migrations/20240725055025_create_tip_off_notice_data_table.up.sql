CREATE TABLE tip_off_notice_data (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT,
  `orig_id` bigint(20) NOT NULL,
  `title` varchar(128) NOT NULL DEFAULT '',
  `content` varchar(2048) NOT NULL DEFAULT '',
  `date_time` varchar(255) NOT NULL DEFAULT '',
  `short_time` varchar(255) NOT NULL DEFAULT '',
  `shi_jian_chuo` bigint(20) unsigned NOT NULL DEFAULT '0',
  `cate_id` varchar(255) NOT NULL DEFAULT '',
  `cate_name` varchar(255) NOT NULL DEFAULT '',
  `comments` int(10) unsigned NOT NULL DEFAULT '0',
  `lou_zhu` varchar(255) NOT NULL DEFAULT '',
  `lou_zhu_reg_time` varchar(255) NOT NULL DEFAULT '',
  `url` varchar(255) NOT NULL DEFAULT '',
  `pachong_id` int(11) DEFAULT NULL,
  `yuan_url` varchar(128) NOT NULL DEFAULT '',
  `is_notice` tinyint(3) unsigned NOT NULL DEFAULT '0',
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '创建时间戳',
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间戳',
   PRIMARY KEY (`id`) USING BTREE,
   KEY `idx_createAt_isNotice` (`created_at`,`is_notice`) USING BTREE,
   KEY `uniq_origId` (`orig_id`) USING BTREE,
   KEY `idx_title` (`title`)
)COMMENT="线报通知内容";;