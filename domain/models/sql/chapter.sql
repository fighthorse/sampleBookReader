CREATE TABLE `chapter` (
                           `id` int unsigned NOT NULL AUTO_INCREMENT,
                           `book_id` int unsigned NOT NULL,
                           `chapter_name` varchar(255) DEFAULT NULL COMMENT '章节标题',
                           `chapter_content` longtext COMMENT '章节内容',
                           `chapter_rank` int unsigned DEFAULT NULL COMMENT '章节排序',
                           PRIMARY KEY (`id`),
                           UNIQUE KEY `IDX` (`book_id`,`chapter_rank`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=7 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='书籍-章节列表';

