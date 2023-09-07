/*
 Navicat Premium Data Transfer

 Source Server         : 39.103.213.44
 Source Server Type    : MySQL
 Source Server Version : 80100
 Source Host           : 39.103.213.44:3306
 Source Schema         : demo

 Target Server Type    : MySQL
 Target Server Version : 80100
 File Encoding         : 65001

 Date: 07/09/2023 17:05:24
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `id` int NOT NULL COMMENT '主键ID',
  `uid` int DEFAULT NULL COMMENT '用户uid',
  `username` varchar(120) COLLATE utf8mb4_general_ci NOT NULL COMMENT '用户名',
  `phone` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '手机号',
  `password` char(8) COLLATE utf8mb4_general_ci NOT NULL COMMENT '密码',
  `email` varchar(120) COLLATE utf8mb4_general_ci NOT NULL COMMENT '邮箱',
  `ip` varchar(60) COLLATE utf8mb4_general_ci NOT NULL COMMENT '注册IP',
  `created_at` datetime DEFAULT NULL COMMENT '注册时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
