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

 Date: 07/09/2023 17:05:34
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for user_extend
-- ----------------------------
DROP TABLE IF EXISTS `user_extend`;
CREATE TABLE `user_extend` (
  `id` int NOT NULL COMMENT '主键ID',
  `uid` int DEFAULT NULL COMMENT '用户uid',
  `sex` tinyint DEFAULT '0' COMMENT '性别 0未知 1男 2女',
  `birthday` varchar(60) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT '出生日期',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

SET FOREIGN_KEY_CHECKS = 1;
