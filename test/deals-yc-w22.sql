/*
 Navicat Premium Data Transfer

 Source Server         : MDB Local
 Source Server Type    : MySQL
 Source Server Version : 101002
 Source Host           : localhost:3306
 Source Schema         : deals-yc-w22

 Target Server Type    : MySQL
 Target Server Version : 101002
 File Encoding         : 65001

 Date: 30/05/2024 13:51:18
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for profile
-- ----------------------------
DROP TABLE IF EXISTS `profile`;
CREATE TABLE `profile`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `User_Id` bigint(20) NULL DEFAULT NULL,
  `Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Birthdate` date NULL DEFAULT NULL,
  `Gender` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `Address` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `BuildingNumber` varchar(5) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Regency_Id` bigint(20) NULL DEFAULT NULL,
  `Postalcode` varchar(6) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `PhoneNumber` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `WhatsappNumber` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ProfileImgSrc` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `LastGenerateView` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `idx_Profile_DeletedAt`(`DeletedAt`) USING BTREE,
  INDEX `fk_Profile_User`(`User_Id`) USING BTREE,
  CONSTRAINT `fk_Profile_User` FOREIGN KEY (`User_Id`) REFERENCES `user` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of profile
-- ----------------------------
INSERT INTO `profile` VALUES (1, '2024-05-30 13:28:11.002', '2024-05-30 13:28:11.002', NULL, 1, 'andika bagus', NULL, 1, '', NULL, 0, NULL, NULL, NULL, NULL, '0000-00-00 00:00:00.000');
INSERT INTO `profile` VALUES (2, '2024-05-30 13:29:07.431', '2024-05-30 13:29:07.431', NULL, 2, 'andika. igit', NULL, 1, '', NULL, 0, NULL, NULL, NULL, NULL, '0000-00-00 00:00:00.000');
INSERT INTO `profile` VALUES (3, '2024-05-30 13:29:37.424', '2024-05-30 13:29:37.424', NULL, 3, 'zainal. gham', NULL, 1, '', NULL, 0, NULL, NULL, NULL, NULL, '0000-00-00 00:00:00.000');
INSERT INTO `profile` VALUES (4, '2024-05-30 13:30:30.837', '2024-05-30 13:30:30.837', NULL, 4, 'fransiska ginting', NULL, 0, '', NULL, 0, NULL, NULL, NULL, NULL, '0000-00-00 00:00:00.000');
INSERT INTO `profile` VALUES (5, '2024-05-30 13:31:12.373', '2024-05-30 13:31:12.373', NULL, 5, 'leon edewards', NULL, 1, '', NULL, 0, NULL, NULL, NULL, NULL, '0000-00-00 00:00:00.000');
INSERT INTO `profile` VALUES (6, '2024-05-30 13:32:57.859', '2024-05-30 13:32:57.859', NULL, 6, 'sinta dewi', NULL, 0, '', NULL, 0, NULL, NULL, NULL, NULL, '0000-00-00 00:00:00.000');
INSERT INTO `profile` VALUES (7, '2024-05-30 13:34:10.551', '2024-05-30 13:34:10.551', NULL, 8, 'dewinta akrain', NULL, 0, '', NULL, 0, NULL, NULL, NULL, NULL, '0000-00-00 00:00:00.000');
INSERT INTO `profile` VALUES (8, '2024-05-30 13:33:59.209', '2024-05-30 13:34:10.550', NULL, 8, 'gigih santoso', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (9, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 9, 'ichan rachman', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (10, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 10, 'federik simanjuntak', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (11, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 11, 'lina inverse', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (12, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 12, 'luna moonfang', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (13, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 13, 'muerta avalaz', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (14, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 14, 'edgar druid', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (15, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 15, 'vio salman', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (16, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 16, 'santika dewi', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (17, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 17, 'bagus agung', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (18, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 18, 'tika sulaiman', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (19, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 19, 'wardata buyana', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (20, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 20, 'haikal rizki', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (21, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 21, 'mimin utamin', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (22, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 22, 'antero raya', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (23, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 23, 'wanda hamidah', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (24, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 24, 'ingga sarkosank', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (25, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 25, 'tegan andika', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (26, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 26, 'sulistia eka', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (27, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 27, 'wardono handoko', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (28, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 28, 'joko waluyo', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (29, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 29, 'pujo styo', NULL, 1, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);
INSERT INTO `profile` VALUES (30, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 30, 'sundari tamyis', NULL, 0, NULL, NULL, NULL, NULL, NULL, NULL, NULL, NULL);

-- ----------------------------
-- Table structure for subscription
-- ----------------------------
DROP TABLE IF EXISTS `subscription`;
CREATE TABLE `subscription`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `User_Id` bigint(20) NULL DEFAULT NULL,
  `Type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ExpiredDate` datetime(3) NULL DEFAULT NULL,
  `PaymentMethod_Id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `idx_Subscription_DeletedAt`(`DeletedAt`) USING BTREE,
  INDEX `fk_Subscription_User`(`User_Id`) USING BTREE,
  CONSTRAINT `fk_Subscription_User` FOREIGN KEY (`User_Id`) REFERENCES `user` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for subscriptionlog
-- ----------------------------
DROP TABLE IF EXISTS `subscriptionlog`;
CREATE TABLE `subscriptionlog`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `User_Id` bigint(20) NULL DEFAULT NULL,
  `Type` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ExpiredDate` datetime(3) NULL DEFAULT NULL,
  `PaymentMethod_Id` bigint(20) NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `idx_SubscriptionLog_DeletedAt`(`DeletedAt`) USING BTREE,
  INDEX `fk_SubscriptionLog_User`(`User_Id`) USING BTREE,
  CONSTRAINT `fk_SubscriptionLog_User` FOREIGN KEY (`User_Id`) REFERENCES `user` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Table structure for user
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `Name` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Password` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Status` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  `LoginAttemptCount` bigint(20) NULL DEFAULT NULL,
  `LastSuccessLogin` datetime(3) NULL DEFAULT NULL,
  `LastAllowdLogin` datetime(3) NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  UNIQUE INDEX `uni_User_Name`(`Name`) USING BTREE,
  UNIQUE INDEX `uni_User_Email`(`Email`) USING BTREE,
  INDEX `idx_User_DeletedAt`(`DeletedAt`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 31 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES (1, '2024-05-30 13:26:45.082', '2024-05-30 13:28:11.000', NULL, 'andika.bagus', 'andika.bagus@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (2, '2024-05-30 13:28:54.927', '2024-05-30 13:29:07.430', NULL, 'andika.sigit', 'andika.sigit@gmail.com', '$2a$10$2RnGDwIn.yaF4gySvVmorOlt8wKoY3oM9spYZiTnlXBUjQ487AOSC', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (3, '2024-05-30 13:29:22.601', '2024-05-30 13:29:37.423', NULL, 'zainal.igham', 'zainal.igham@gmail.com', '$2a$10$J5g9ubxroQDi/6wKSNlGLepe4yCFMGR72arw8dXJbUvJpoTHg0IKq', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (4, '2024-05-30 13:30:18.315', '2024-05-30 13:30:30.835', NULL, 'fransiska.ginting', 'fransiska.ginting@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (5, '2024-05-30 13:30:59.025', '2024-05-30 13:31:12.372', NULL, 'leon.edewards', 'leon.edewards@gmail.com', '$2a$10$OSuNN62NFTbqYy28vuJhn.oYh69GFqGjoRjo780PabsmUruUGtksS', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (6, '2024-05-30 13:32:41.904', '2024-05-30 13:32:57.857', NULL, 'sinta.dewi', 'sinta.dewi@gmail.com', '$2a$10$eTdWGTokly8B9lXyMlU5O.mmZVGbooyPla4ZHw43.Lll6RJclS3bO', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (7, '2024-05-30 13:33:14.480', '2024-05-30 13:33:14.480', NULL, 'dewinta.akrain', 'dewinta.akrain@gmail.com', '$2a$10$kUoxk1605VXPyk1zKtfPve.WUD26cm.kI5M8ngmpUHwhkDEtYhf3m', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (8, '2024-05-30 13:33:59.209', '2024-05-30 13:34:10.550', NULL, 'gigih.santoso', 'gigih.santoso@gmail.com', '$2a$10$psCWWqvdt1U7bJEcd4WlhO2d9sQPKGIwHU0iRitR5z.EJxsOuRf/u', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (9, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'ichan.rachman', 'ichan.rachman@gmail.com', '$2a$10$psCWWqvdt1U7bJEcd4WlhO2d9sQPKGIwHU0iRitR5z.EJxsOuRf/u', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (10, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'federik.simanjuntak', 'federik.simanjuntak@gmail.com', '$2a$10$psCWWqvdt1U7bJEcd4WlhO2d9sQPKGIwHU0iRitR5z.EJxsOuRf/u', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (11, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'lina.inverse', 'lina.inverse@gmail.com', '$2a$10$psCWWqvdt1U7bJEcd4WlhO2d9sQPKGIwHU0iRitR5z.EJxsOuRf/u', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (12, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'luna.moonfang', 'luna.moonfang@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (13, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'muerta.avalaz', 'muerta.avalaz@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (14, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'edgar.druid', 'edgar.druid@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (15, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'vio.salman', 'vio.salman@gmail.com', '$2a$10$J5g9ubxroQDi/6wKSNlGLepe4yCFMGR72arw8dXJbUvJpoTHg0IKq', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (16, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'santika.dewi', 'santika.dewi@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (17, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'bagus.agung', 'bagus.agung@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (18, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'tika.sulaiman', 'tika.sulaiman@gmail.com', '$2a$10$J5g9ubxroQDi/6wKSNlGLepe4yCFMGR72arw8dXJbUvJpoTHg0IKq', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (19, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'wardata.buyana', 'wardata.buyana@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (20, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'haikal.rizki', 'haikal.rizki@gmail.com', '$2a$10$J5g9ubxroQDi/6wKSNlGLepe4yCFMGR72arw8dXJbUvJpoTHg0IKq', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (21, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'mimin.utamin', 'mimin.utamin@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (22, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'antero.raya', 'antero.raya@gmail.com', '$2a$10$2RnGDwIn.yaF4gySvVmorOlt8wKoY3oM9spYZiTnlXBUjQ487AOSC', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (23, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'wanda.hamidah', 'wanda.hamidah@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (24, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'ingga.sarkosank', 'ingga.sarkosank@gmail.com', '$2a$10$2RnGDwIn.yaF4gySvVmorOlt8wKoY3oM9spYZiTnlXBUjQ487AOSC', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (25, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'tegan.andika', 'tegan.andika@gmail.com', '$2a$10$AJBFH8WO5eSWjP85DxsDwOeb/zZpmY0BBvbk.tKe0heTLfcfOswQ6', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (26, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'sulistia.eka', 'sulistia.eka@gmail.com', '$2a$10$eTdWGTokly8B9lXyMlU5O.mmZVGbooyPla4ZHw43.Lll6RJclS3bO', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (27, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'wardono.handoko', 'wardono.handoko@gmail.com ', '$2a$10$eTdWGTokly8B9lXyMlU5O.mmZVGbooyPla4ZHw43.Lll6RJclS3bO', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (28, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'joko.waluyo', 'joko.waluyo@gmail.com ', '$2a$10$eTdWGTokly8B9lXyMlU5O.mmZVGbooyPla4ZHw43.Lll6RJclS3bO', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (29, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'pujo.styo', 'pujo.styo@gmail.com', '$2a$10$psCWWqvdt1U7bJEcd4WlhO2d9sQPKGIwHU0iRitR5z.EJxsOuRf/u', 1, 0, NULL, NULL);
INSERT INTO `user` VALUES (30, '2024-05-30 13:33:59.209', '2024-05-30 13:33:59.209', NULL, 'sundari.tamyis', 'sundari.tamyis@gmail.com', '$2a$10$psCWWqvdt1U7bJEcd4WlhO2d9sQPKGIwHU0iRitR5z.EJxsOuRf/u', 1, 0, NULL, NULL);

-- ----------------------------
-- Table structure for usertoken
-- ----------------------------
DROP TABLE IF EXISTS `usertoken`;
CREATE TABLE `usertoken`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `User_Email` varchar(100) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Type` varchar(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `Token` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  `ExpiredAt` datetime(3) NULL DEFAULT NULL,
  `AttemptCount` tinyint(3) UNSIGNED NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `idx_UserToken_DeletedAt`(`DeletedAt`) USING BTREE,
  INDEX `unique`(`User_Email`) USING BTREE,
  CONSTRAINT `fk_UserToken_User` FOREIGN KEY (`User_Email`) REFERENCES `user` (`Email`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 9 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of usertoken
-- ----------------------------
INSERT INTO `usertoken` VALUES (7, '2024-05-30 13:33:14.482', '2024-05-30 13:33:14.482', NULL, 'dewinta.akrain@gmail.com', 'ConfirmByEmail', '2c6adcf5-6292-4a8b-8c24-de26e5eea24f', '2024-06-02 13:33:14.481', 0);

-- ----------------------------
-- Table structure for viewresult
-- ----------------------------
DROP TABLE IF EXISTS `viewresult`;
CREATE TABLE `viewresult`  (
  `Id` bigint(20) NOT NULL AUTO_INCREMENT,
  `CreatedAt` datetime(3) NULL DEFAULT NULL,
  `UpdatedAt` datetime(3) NULL DEFAULT NULL,
  `DeletedAt` datetime(3) NULL DEFAULT NULL,
  `Viewer_Profile_Id` bigint(20) NULL DEFAULT NULL,
  `Target_Profile_Id` bigint(20) NULL DEFAULT NULL,
  `Result` varchar(10) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NULL DEFAULT NULL,
  PRIMARY KEY (`Id`) USING BTREE,
  INDEX `idx_ViewResult_DeletedAt`(`DeletedAt`) USING BTREE,
  INDEX `fk_ViewResult_Viewer_Profile`(`Viewer_Profile_Id`) USING BTREE,
  INDEX `fk_ViewResult_Target_Profile`(`Target_Profile_Id`) USING BTREE,
  CONSTRAINT `fk_ViewResult_Target_Profile` FOREIGN KEY (`Target_Profile_Id`) REFERENCES `profile` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `fk_ViewResult_Viewer_Profile` FOREIGN KEY (`Viewer_Profile_Id`) REFERENCES `profile` (`Id`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB AUTO_INCREMENT = 1 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;

SET FOREIGN_KEY_CHECKS = 1;
