#规则表
DROP TABLE IF EXISTS `yf_auth_rule`;
CREATE TABLE `yf_auth_rule` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(80) NOT NULL DEFAULT '' COMMENT '规则名唯一标识',
    `title` char(20) NOT NULL DEFAULT ''  COMMENT '规则中文名',
    `category` tinyint(1) NOT NULL DEFAULT 0 COMMENT '规则类型 1前端规则权限 2后端规则权限',
    `categoryName` char(100) NOT NULL DEFAULT '' COMMENT '规则类型描述',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1正常 0禁用',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_name` (`name`)
) ENGINE=innodb  DEFAULT CHARSET=utf8mb4;

#角色表
DROP TABLE IF EXISTS `yf_auth_role`;
CREATE TABLE `yf_auth_role` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `title` char(100) NOT NULL DEFAULT '' COMMENT '角色中文名',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1正常 0禁用',
    `rules` varchar(1024) NOT NULL DEFAULT '' COMMENT '角色拥有规则ID,多个则用,隔开',
    PRIMARY KEY (`id`)
) ENGINE=innodb  DEFAULT CHARSET=utf8mb4;


#角色明细表
DROP TABLE IF EXISTS `yf_auth_role_access`;
CREATE TABLE `yf_auth_role_access` (
    `user_id` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
    `role_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '角色ID',
    UNIQUE KEY `uni_user_id_role_id` (`user_id`,`role_id`),
    KEY `ind_user_id` (`user_id`),
    KEY `ind_role_id` (`role_id`)
) ENGINE=innodb DEFAULT CHARSET=utf8mb4;