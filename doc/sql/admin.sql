CREATE TABLE `admins` (
                                  `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
                                  `name` varchar(60) NOT NULL DEFAULT '' COMMENT '员工名称',
                                  `phone` varchar(12) NOT NULL COMMENT '手机号',
                                  `password` varchar(60) NOT NULL COMMENT '密码',
                                  `status` tinyint NOT NULL DEFAULT '1' COMMENT '状态 1、启用 2、关闭',
                                  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '注册时间',
                                  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '修改时间',
                                  `deleted_at` timestamp NULL DEFAULT NULL COMMENT '删除时间',
                                  PRIMARY KEY (`id`) USING BTREE,
                                  UNIQUE KEY `idx_phone_deleted_at` (`phone`,`deleted_at`) USING BTREE
) ENGINE=InnoDB COMMENT='管理员表';
