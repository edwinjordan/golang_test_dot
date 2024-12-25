/*
 Navicat Premium Data Transfer

 Source Server         : SERVER_LARAGON
 Source Server Type    : MySQL
 Source Server Version : 50724 (5.7.24)
 Source Host           : 127.0.0.1:3306
 Source Schema         : golang_test_dot

 Target Server Type    : MySQL
 Target Server Version : 50724 (5.7.24)
 File Encoding         : 65001

 Date: 25/12/2024 21:49:08
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for ms_category
-- ----------------------------
DROP TABLE IF EXISTS `ms_category`;
CREATE TABLE `ms_category`  (
  `category_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `category_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `category_delete_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`category_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of ms_category
-- ----------------------------
INSERT INTO `ms_category` VALUES ('601c8f5480c47606456df7da91fbf20b', 'Sepatu', '0000-00-00 00:00:00');
INSERT INTO `ms_category` VALUES ('8b64e6326d6c5f1203fa2251709ba9ee', 'Topi', '0000-00-00 00:00:00');
INSERT INTO `ms_category` VALUES ('f8dc43bee4d26fad60e5427321ed18db', 'Tas Gunung', '0000-00-00 00:00:00');

-- ----------------------------
-- Table structure for tb_customer
-- ----------------------------
DROP TABLE IF EXISTS `tb_customer`;
CREATE TABLE `tb_customer`  (
  `customer_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `customer_code` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `customer_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `customer_gender` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `customer_phonenumber` varchar(30) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `customer_email` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `customer_password` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `customer_status` tinyint(1) NULL DEFAULT NULL,
  `customer_create_at` datetime NULL DEFAULT NULL,
  `customer_update_at` datetime NULL DEFAULT NULL,
  PRIMARY KEY (`customer_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_customer
-- ----------------------------
INSERT INTO `tb_customer` VALUES ('9533d4b4c417ebab4f09e54ab9c96857', 'ESC-1224-001', 'edwin yordan', 'L', '085765945759', 'edwinlaksono12@gmail.com', '$2a$10$Wgx80XXPCSA9ZZZxINTCe.A8AXPWFFxyqyR9i1naoq6FQkAXx0joK', 0, '2024-12-24 07:51:34', '2024-12-24 08:17:12');
INSERT INTO `tb_customer` VALUES ('a645a0717f682b4c4db990ccc5f19948', 'ESC-1224-002', 'Jono', 'L', '086786543222', 'jono@gmail.com', '$2a$10$15myar1kDjnpQJRs7dTopubBo25IRzXZwkN/G/f.Q0w1iyrS3TmUK', 0, '2024-12-25 10:05:05', '2024-12-25 10:05:05');

-- ----------------------------
-- Table structure for tb_customer_address
-- ----------------------------
DROP TABLE IF EXISTS `tb_customer_address`;
CREATE TABLE `tb_customer_address`  (
  `address_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `address_customer_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `address_text` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  `address_name` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `address_create_at` datetime NULL DEFAULT NULL,
  `address_update_at` datetime NULL DEFAULT NULL,
  `address_postal_code` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  PRIMARY KEY (`address_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_customer_address
-- ----------------------------
INSERT INTO `tb_customer_address` VALUES ('26bbefb00574d6f2abdcd01608fb2200', '9533d4b4c417ebab4f09e54ab9c96857', 'Jl Raya Kediri', 'Rumah sendiri', '2024-12-25 11:17:14', '2024-12-25 11:17:14', '64181');
INSERT INTO `tb_customer_address` VALUES ('9414639b9213bd1926228bf1927817d1', '9533d4b4c417ebab4f09e54ab9c96857', 'Jl Baru no 1', 'Alamat', '2024-12-24 08:59:07', '2024-12-24 08:59:07', '');
INSERT INTO `tb_customer_address` VALUES ('e9830ba03c3cf97ad88412d9dad60c7e', '9533d4b4c417ebab4f09e54ab9c96857', 'Desa Blimbing Rt 03', 'Alamat 2', '2024-12-24 09:03:13', '2024-12-24 09:03:13', '');
INSERT INTO `tb_customer_address` VALUES ('fb7b9e61f5bddbc3f80ab2c942545dba', '9533d4b4c417ebab4f09e54ab9c96857', 'Jalan Raya Pare', 'Rumah Sendiri', '2024-12-25 16:08:22', '2024-12-25 16:08:22', '64181');

-- ----------------------------
-- Table structure for tb_order
-- ----------------------------
DROP TABLE IF EXISTS `tb_order`;
CREATE TABLE `tb_order`  (
  `order_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `order_customer_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `order_total_item` int(11) NULL DEFAULT NULL,
  `order_subtotal` double NULL DEFAULT NULL,
  `order_discount` double NULL DEFAULT NULL,
  `order_total` double NULL DEFAULT NULL,
  `order_status` tinyint(1) NULL DEFAULT NULL,
  `order_create_at` datetime NULL DEFAULT NULL,
  `order_inv_number` varchar(100) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `order_notes` text CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL,
  PRIMARY KEY (`order_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_order
-- ----------------------------
INSERT INTO `tb_order` VALUES ('4717aaca0886bcae778061de2b6a1358', '9533d4b4c417ebab4f09e54ab9c96857', 3, 200000, 0, 600000, 1, '2024-12-24 15:39:00', 'ORN-241224-002', 'ok');
INSERT INTO `tb_order` VALUES ('75f5634b30fae6390416b1b59e7ffe57', '9533d4b4c417ebab4f09e54ab9c96857', 3, 200000, 0, 600000, 1, '2024-12-25 15:01:15', 'ORN-251224-003', 'ok');
INSERT INTO `tb_order` VALUES ('76f828396035d33000c7489252f262a4', '9533d4b4c417ebab4f09e54ab9c96857', 2, 200000, 0, 400000, 1, '2024-12-24 15:38:07', 'ORN-241224-001', 'ok');
INSERT INTO `tb_order` VALUES ('82e22438323a1c5a75d48e448bff818f', '9533d4b4c417ebab4f09e54ab9c96857', 3, 200000, 0, 600000, 1, '2024-12-25 15:00:26', 'ORN-251224-001', 'ok');
INSERT INTO `tb_order` VALUES ('d677fd620269859b2606fffc4eaceaf3', '9533d4b4c417ebab4f09e54ab9c96857', 3, 200000, 0, 600000, 1, '2024-12-25 15:00:39', 'ORN-251224-002', 'ok');

-- ----------------------------
-- Table structure for tb_order_detail
-- ----------------------------
DROP TABLE IF EXISTS `tb_order_detail`;
CREATE TABLE `tb_order_detail`  (
  `order_detail_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NOT NULL,
  `order_detail_parent_id` varchar(32) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `order_detail_product` varchar(255) CHARACTER SET latin1 COLLATE latin1_swedish_ci NULL DEFAULT NULL,
  `order_detail_qty` int(11) NULL DEFAULT NULL,
  `order_detail_price` double NULL DEFAULT NULL,
  `order_detail_subtotal` double NULL DEFAULT NULL,
  PRIMARY KEY (`order_detail_id`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = latin1 COLLATE = latin1_swedish_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of tb_order_detail
-- ----------------------------
INSERT INTO `tb_order_detail` VALUES ('01b5f043f7de6e33ec1e9f6a35ded6b2', 'd677fd620269859b2606fffc4eaceaf3', 'Celana', 2, 400000, 800000);
INSERT INTO `tb_order_detail` VALUES ('01b6d969844252aaf60e2e38d52996a5', '75f5634b30fae6390416b1b59e7ffe57', 'Baju', 1, 400000, 400000);
INSERT INTO `tb_order_detail` VALUES ('02a3eb8f51dfc5f614a93331b6c3e63b', '82e22438323a1c5a75d48e448bff818f', 'Baju', 1, 400000, 400000);
INSERT INTO `tb_order_detail` VALUES ('1b4988e78490e9b9b74f58a5591f0171', 'd677fd620269859b2606fffc4eaceaf3', 'Baju', 1, 400000, 400000);
INSERT INTO `tb_order_detail` VALUES ('503d1c278578145d6a53084f6f7dda49', '4717aaca0886bcae778061de2b6a1358', 'Baju', 1, 400000, 400000);
INSERT INTO `tb_order_detail` VALUES ('cf173a6bd351a8bd928cfce946a7595d', '76f828396035d33000c7489252f262a4', 'Baju', 1, 400000, 400000);
INSERT INTO `tb_order_detail` VALUES ('e20e326e9ac04e0476cbfeac4a159d3a', '75f5634b30fae6390416b1b59e7ffe57', 'Celana', 2, 400000, 800000);
INSERT INTO `tb_order_detail` VALUES ('e9dc6fc138a137f4e334861562710145', '4717aaca0886bcae778061de2b6a1358', 'Celana', 2, 400000, 800000);
INSERT INTO `tb_order_detail` VALUES ('ef05522acdac02f39527541a28ac8737', '82e22438323a1c5a75d48e448bff818f', 'Celana', 2, 400000, 800000);

SET FOREIGN_KEY_CHECKS = 1;
