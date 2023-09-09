CREATE TABLE `book` (
                        `id` int unsigned NOT NULL AUTO_INCREMENT COMMENT '书籍id',
                        `book_title` varchar(255) NOT NULL COMMENT '书籍标题',
                        `book_desc` varchar(255) NOT NULL COMMENT '书籍描述',
                        `author_id` int DEFAULT NULL COMMENT '作者id',
                        `chapter_total` int DEFAULT NULL COMMENT '总章节数量',
                        `copyright` varchar(255) DEFAULT NULL COMMENT '版权方 出版社',
                        `state` tinyint DEFAULT '1' COMMENT '状态 默认1 待上架 5 已上架 0 删除',
                        `category_id` int DEFAULT '1' COMMENT '分类id',
                        PRIMARY KEY (`id`),
                        KEY `idx` (`state`),
                        KEY `idx_categry` (`category_id`),
                        KEY `idx_author` (`author_id`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='书籍';