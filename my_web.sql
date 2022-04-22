/*
 Navicat Premium Data Transfer

 Source Server         : my-web
 Source Server Type    : MySQL
 Source Server Version : 50734
 Source Host           : 119.29.18.159:3306
 Source Schema         : my_web

 Target Server Type    : MySQL
 Target Server Version : 50734
 File Encoding         : 65001

 Date: 22/04/2022 11:17:23
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for admin_user
-- ----------------------------
DROP TABLE IF EXISTS `admin_user`;
CREATE TABLE `admin_user`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `avatar` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `password` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `token` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `expiration_at` datetime(0) NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 2 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of admin_user
-- ----------------------------
INSERT INTO `admin_user` VALUES (1, 'admin', 'https://p9-passport.byteacctimg.com/img/user-avatar/2c7703e223b37e028d27cb904683e015~300x300.image', '6f84523218386b90fb548b4e5b22a887', 'ETM01GV4870TLIFB8LR61WYHNPI7GPTL', '2021-11-30 07:54:49');

-- ----------------------------
-- Table structure for url_list
-- ----------------------------
DROP TABLE IF EXISTS `url_list`;
CREATE TABLE `url_list`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `type_id` int(10) UNSIGNED NULL DEFAULT NULL,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  `url` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 80 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of url_list
-- ----------------------------
INSERT INTO `url_list` VALUES (1, 1, 'JavaScript', 'www.baidu.com');
INSERT INTO `url_list` VALUES (2, 1, 'TypeScript', 'www.baidu.com');
INSERT INTO `url_list` VALUES (3, 1, 'Html', 'www.baidu.com');
INSERT INTO `url_list` VALUES (4, 1, 'CSS', 'www.baidu.com');
INSERT INTO `url_list` VALUES (5, 1, '小程序', 'www.baidu.com');
INSERT INTO `url_list` VALUES (6, 1, 'Vite', 'www.baidu.com');
INSERT INTO `url_list` VALUES (7, 1, 'WebPack', 'www.baidu.com');
INSERT INTO `url_list` VALUES (8, 1, 'React', 'www.baidu.com');
INSERT INTO `url_list` VALUES (9, 1, 'Angular', 'www.baidu.com');
INSERT INTO `url_list` VALUES (10, 1, 'JQuery', 'www.baidu.com');
INSERT INTO `url_list` VALUES (11, 2, 'Golang', 'www.baidu.com');
INSERT INTO `url_list` VALUES (12, 2, 'Python', 'www.baidu.com');
INSERT INTO `url_list` VALUES (13, 2, 'PHP', 'www.baidu.com');
INSERT INTO `url_list` VALUES (14, 2, 'Java', 'www.baidu.com');
INSERT INTO `url_list` VALUES (15, 2, 'C', 'www.baidu.com');
INSERT INTO `url_list` VALUES (16, 2, 'C++', 'www.baidu.com');
INSERT INTO `url_list` VALUES (17, 2, 'C#', 'www.baidu.com');
INSERT INTO `url_list` VALUES (18, 2, 'VB', 'www.baidu.com');
INSERT INTO `url_list` VALUES (19, 2, 'Perl', 'www.baidu.com');
INSERT INTO `url_list` VALUES (20, 2, 'Delphi', 'www.baidu.com');
INSERT INTO `url_list` VALUES (21, 2, 'Swift', 'www.baidu.com');
INSERT INTO `url_list` VALUES (22, 2, 'Pascal', 'www.baidu.com');
INSERT INTO `url_list` VALUES (23, 2, 'Rust', 'www.baidu.com');
INSERT INTO `url_list` VALUES (24, 2, 'Scratch', 'www.baidu.com');
INSERT INTO `url_list` VALUES (25, 2, '易语言', 'www.baidu.com');
INSERT INTO `url_list` VALUES (26, 2, 'BASIC', 'www.baidu.com');
INSERT INTO `url_list` VALUES (27, 2, 'Ruby', 'www.baidu.com');
INSERT INTO `url_list` VALUES (28, 3, 'Linux', 'www.baidu.com');
INSERT INTO `url_list` VALUES (29, 3, '储存', 'www.baidu.com');
INSERT INTO `url_list` VALUES (30, 3, 'K8s', 'www.baidu.com');
INSERT INTO `url_list` VALUES (31, 3, 'Docker', 'www.baidu.com');
INSERT INTO `url_list` VALUES (32, 3, 'Jekins', 'www.baidu.com');
INSERT INTO `url_list` VALUES (33, 3, 'DevOps', 'www.baidu.com');
INSERT INTO `url_list` VALUES (34, 3, 'KVM', 'www.baidu.com');
INSERT INTO `url_list` VALUES (35, 3, 'Shell', 'www.baidu.com');
INSERT INTO `url_list` VALUES (36, 3, 'Ansible', 'www.baidu.com');
INSERT INTO `url_list` VALUES (37, 3, 'Nginx', 'www.baidu.com');
INSERT INTO `url_list` VALUES (38, 3, 'GitLab', 'www.baidu.com');
INSERT INTO `url_list` VALUES (39, 4, 'Seleium', 'www.baidu.com');
INSERT INTO `url_list` VALUES (40, 4, 'jmeter', 'www.baidu.com');
INSERT INTO `url_list` VALUES (41, 4, 'appium', 'www.baidu.com');
INSERT INTO `url_list` VALUES (42, 4, 'Mysql', 'www.baidu.com');
INSERT INTO `url_list` VALUES (43, 4, 'monkey runner', 'www.baidu.com');
INSERT INTO `url_list` VALUES (44, 4, 'SoapUI', 'www.baidu.com');
INSERT INTO `url_list` VALUES (45, 5, '面试', 'www.baidu.com');
INSERT INTO `url_list` VALUES (46, 5, '排序', 'www.baidu.com');
INSERT INTO `url_list` VALUES (47, 5, '查找', 'www.baidu.com');
INSERT INTO `url_list` VALUES (48, 5, '树', 'www.baidu.com');
INSERT INTO `url_list` VALUES (49, 5, '图', 'www.baidu.com');
INSERT INTO `url_list` VALUES (50, 5, '数据结构', 'www.baidu.com');
INSERT INTO `url_list` VALUES (51, 5, '动态规划', 'www.baidu.com');
INSERT INTO `url_list` VALUES (52, 5, '堆栈', 'www.baidu.com');
INSERT INTO `url_list` VALUES (53, 5, '矩阵', 'www.baidu.com');
INSERT INTO `url_list` VALUES (54, 6, '计算机视觉', 'www.baidu.com');
INSERT INTO `url_list` VALUES (55, 6, '语音识别', 'www.baidu.com');
INSERT INTO `url_list` VALUES (56, 6, 'Tensorflow', 'www.baidu.com');
INSERT INTO `url_list` VALUES (57, 6, 'Keras', 'www.baidu.com');
INSERT INTO `url_list` VALUES (58, 6, 'PyTorch', 'www.baidu.com');
INSERT INTO `url_list` VALUES (59, 6, 'Caffe', 'www.baidu.com');
INSERT INTO `url_list` VALUES (60, 6, 'PaddlePaddle', 'www.baidu.com');
INSERT INTO `url_list` VALUES (61, 7, 'Arm', 'www.baidu.com');
INSERT INTO `url_list` VALUES (62, 7, 'STM32', 'www.baidu.com');
INSERT INTO `url_list` VALUES (63, 7, 'MCU', 'www.baidu.com');
INSERT INTO `url_list` VALUES (64, 7, 'PCB', 'www.baidu.com');
INSERT INTO `url_list` VALUES (65, 7, 'FPGA', 'www.baidu.com');
INSERT INTO `url_list` VALUES (66, 7, '嵌入式', 'www.baidu.com');
INSERT INTO `url_list` VALUES (67, 7, '单片机', 'www.baidu.com');
INSERT INTO `url_list` VALUES (68, 7, 'DSP', 'www.baidu.com');
INSERT INTO `url_list` VALUES (69, 7, 'ADAS', 'www.baidu.com');
INSERT INTO `url_list` VALUES (70, 8, 'DDD', 'www.baidu.com');
INSERT INTO `url_list` VALUES (71, 8, '分布式', 'www.baidu.com');
INSERT INTO `url_list` VALUES (72, 8, 'Sass', 'www.baidu.com');
INSERT INTO `url_list` VALUES (73, 8, '负载均衡', 'www.baidu.com');
INSERT INTO `url_list` VALUES (74, 8, '中台', 'www.baidu.com');
INSERT INTO `url_list` VALUES (75, 8, '云原生', 'www.baidu.com');
INSERT INTO `url_list` VALUES (76, 9, '渗透测试', 'www.baidu.com');
INSERT INTO `url_list` VALUES (77, 9, '软件逆向', 'www.baidu.com');
INSERT INTO `url_list` VALUES (78, 9, '病毒分析', 'www.baidu.com');
INSERT INTO `url_list` VALUES (79, 9, '漏洞扫描', 'http://www.baidu.com');

-- ----------------------------
-- Table structure for url_type
-- ----------------------------
DROP TABLE IF EXISTS `url_type`;
CREATE TABLE `url_type`  (
  `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT,
  `name` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 10 CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of url_type
-- ----------------------------
INSERT INTO `url_type` VALUES (1, '前端');
INSERT INTO `url_type` VALUES (2, '后端');
INSERT INTO `url_type` VALUES (3, '运维');
INSERT INTO `url_type` VALUES (4, '测试');
INSERT INTO `url_type` VALUES (5, '算法');
INSERT INTO `url_type` VALUES (6, 'AI');
INSERT INTO `url_type` VALUES (7, '硬件');
INSERT INTO `url_type` VALUES (8, '架构');
INSERT INTO `url_type` VALUES (9, '安全');

SET FOREIGN_KEY_CHECKS = 1;
