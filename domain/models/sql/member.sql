CREATE TABLE `member` (
              `id` int unsigned NOT NULL AUTO_INCREMENT,
              `member_name` varchar(255) NOT NULL COMMENT '会员用户登录账户',
              `member_pwd` varchar(255) NOT NULL COMMENT '登录密码',
              `member_desc` varchar(255) DEFAULT NULL COMMENT '用户自我介绍',
              `read_books` int DEFAULT NULL COMMENT '阅读书籍数量',
              `register_day` varchar(32) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci DEFAULT NULL COMMENT '注册时间',
              PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='用户会员表';