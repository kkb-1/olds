CREATE DATABASE olds_user;

USE olds_user;

CREATE TABLE `olds_user` (
    `id` varchar(100) not null comment "唯一标识",
    `username` varchar(100) not null comment "唯一用户名",
    `password` varchar(100) not null comment "加密密码",
    `nickname` varchar(100) comment "昵称",
    `avatar` varchar(100) comment "头像地址",
    `status` tinyint not null comment "用户状态：-1禁用、1可用",
    `create_time` int comment "创建时间",
    `update_time` int comment "修改时间",
    primary key (`id`),
    unique (`username`)
)COMMENT='管理员用户信息表';

INSERT INTO `olds_user` (`id`, `username`, `password`, `status`) values ('root', 'root', 'e10adc3949ba59abbe56e057f20f883e', 1);
