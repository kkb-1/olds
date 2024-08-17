# 就是后续的open_id是属于监护人的，而uid是老人的，只有user表是属于自己的

CREATE DATABASE `intelligent_systems`;

USE `intelligent_systems`;

CREATE TABLE `user` (
    `open_id` char(255) NOT NULL COMMENT '微信生成用户唯一标识符',
    `uid` char(20) NOT NULL COMMENT '系统生成用户唯一标识符',
    PRIMARY KEY (`open_id`),
    UNIQUE KEY `uid` (`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

insert  into `user`(`open_id`,`uid`) values
                                         ('ozlII7UTrfBvNFOTUNidmmp-KDvs','81759CFA2421000'),
                                         ('ozlII7aq_qVlb2VilZRzh7PBUrds','818B0BD2D021000');

CREATE TABLE `details` (
                           `open_id` char(125) NOT NULL,
                           `uid` char(20) NOT NULL,
                           `phone` varchar(25) DEFAULT NULL COMMENT '手机号',
                           `role` int DEFAULT NULL COMMENT '角色（1为监护人，2为老人）',
                           `height` float DEFAULT NULL COMMENT '身高',
                           `weight` float DEFAULT NULL COMMENT '体重',
                           `age` int DEFAULT NULL COMMENT '年龄',
                           `sex` int DEFAULT NULL COMMENT '性别（1为男、2为女）',
                           `smoke` tinyint(1) DEFAULT NULL COMMENT 'true为抽',
                           `drink` tinyint(1) DEFAULT NULL COMMENT 'true为喝酒',
                           `exercise` tinyint(1) DEFAULT NULL COMMENT 'true为有锻炼',
                           PRIMARY KEY (`open_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;


insert  into `details`(`open_id`,`uid`,`phone`,`role`,`height`,`weight`,`age`,`sex`,`smoke`,`drink`,`exercise`) values
                                                                                                                    ('ozlII7aq_qVlb2VilZRzh7PBUrds','818B0BD2D021000','18664815016',2,180,0,66,1,1,1,1),
                                                                                                                    ('ozlII7UTrfBvNFOTUNidmmp-KDvs','81759CFA2421000','18664815016',1,0,0,0,NULL,0,0,0);

CREATE TABLE `binds` (
                         `open_id` char(125) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '监护人的id',
                         `uid` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci NOT NULL COMMENT '被绑定老人的id',
                         `note` varchar(40) DEFAULT NULL COMMENT '备注老人',
                         `confirm` tinyint(1) DEFAULT NULL COMMENT '是否确认绑定（true为确认）',
                         PRIMARY KEY (`open_id`,`uid`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci;

/*Data for the table `binds` */

insert  into `binds`(`open_id`,`uid`,`note`,`confirm`) values
    ('ozlII7UTrfBvNFOTUNidmmp-KDvs','818B0BD2D021000','父亲',1);