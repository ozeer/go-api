# 国内行政地区数据
1、创建数据库及表
```
CREATE TABLE `cnarea_2023`
(
    `id`          mediumint(7) unsigned          NOT NULL AUTO_INCREMENT,
    `level`       tinyint(1) unsigned            NOT NULL COMMENT '层级',
    `parent_code` bigint(14) unsigned            NOT NULL DEFAULT '0' COMMENT '父级行政代码',
    `area_code`   bigint(14) unsigned            NOT NULL DEFAULT '0' COMMENT '行政代码',
    `zip_code`    mediumint(6) unsigned zerofill NOT NULL DEFAULT '000000' COMMENT '邮政编码',
    `city_code`   char(6)                        NOT NULL DEFAULT '' COMMENT '区号',
    `name`        varchar(200)                    NOT NULL DEFAULT '' COMMENT '名称',
    `short_name`  varchar(50)                    NOT NULL DEFAULT '' COMMENT '简称',
    `merger_name` varchar(200)                    NOT NULL DEFAULT '' COMMENT '组合名',
    `pinyin`      varchar(30)                    NOT NULL DEFAULT '' COMMENT '拼音',
    `lng`         decimal(10, 6)                 NOT NULL DEFAULT '0.000000' COMMENT '经度',
    `lat`         decimal(10, 6)                 NOT NULL DEFAULT '0.000000' COMMENT '纬度',
    PRIMARY KEY (`id`),
    UNIQUE KEY `uk_code` (`area_code`) USING BTREE,
    KEY `idx_parent_code` (`parent_code`) USING BTREE
) ENGINE = MyISAM
  DEFAULT CHARSET = utf8 COMMENT ='中国行政地区表';
```

2、导入数据
```
mysql -h127.0.0.1 -uroot -p123456 -D china_area < cnarea_2023.sql
```