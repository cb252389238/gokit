#规则表
DROP TABLE IF EXISTS `auth_rule`;
CREATE TABLE `auth_rule` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` char(80) NOT NULL DEFAULT '' COMMENT '规则名唯一标识',
    `title` char(20) NOT NULL DEFAULT ''  COMMENT '规则中文名',
    `category` tinyint(1) NOT NULL DEFAULT 0 COMMENT '规则类型，例如用户类、文章类',
    `categoryName` char(100) NOT NULL DEFAULT '' COMMENT '规则类型名称',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1正常 0禁用',
    `condition` char(100) NOT NULL DEFAULT '' COMMENT '规则表达式，为空表示存在就验证，不为空表示按照条件验证',
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;

#角色表
DROP TABLE IF EXISTS `auth_role`;
CREATE TABLE `auth_role` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
    `title` char(100) NOT NULL DEFAULT '' COMMENT '用户组中文名',
    `status` tinyint(1) NOT NULL DEFAULT '1' COMMENT '状态 1正常 0禁用',
    `rules` varchar(1024) NOT NULL DEFAULT '' COMMENT '用户组规则ID,多个则用,隔开',
    PRIMARY KEY (`id`)
) ENGINE=MyISAM  DEFAULT CHARSET=utf8;


#角色明细表
DROP TABLE IF EXISTS `auth_role_access`;
CREATE TABLE `auth_role_access` (
    `uid` bigint(20) unsigned NOT NULL DEFAULT 0 COMMENT '用户ID',
    `role_id` int(11) unsigned NOT NULL DEFAULT 0 COMMENT '角色ID',
    UNIQUE KEY `uid_role_id` (`uid`,`role_id`),
    KEY `uid` (`uid`),
    KEY `role_id` (`role_id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;