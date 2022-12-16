CREATE TABLE `sys_user` (
                            `id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
                            `name` varchar(50) NOT NULL,
                            `nick_name` varchar(50) DEFAULT NULL,
                            `avatar` varchar(200) DEFAULT NULL,
                            `password` varchar(20) DEFAULT NULL,
                            `email` varchar(100) DEFAULT NULL,
                            `mobile` varchar(20) DEFAULT NULL,
                            `created_at` bigint(20) NOT NULL,
                            `updated_at` bigint(20) NOT NULL,
                            `deleted_at` bigint(20) DEFAULT NULL,
                            PRIMARY KEY (`id`),
                            UNIQUE KEY `idx_sys_user_name` (`name`)
) ENGINE=InnoDB AUTO_INCREMENT=5 DEFAULT CHARSET=utf8;