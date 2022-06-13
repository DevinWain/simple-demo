/*
Navicat MySQL Data Transfer

Source Server         : localhost_3306
Source Server Version : 80016
Source Host           : localhost:3306
Source Database       : simple_demo

Target Server Type    : MYSQL
Target Server Version : 80016
File Encoding         : 65001

Date: 2022-06-13 12:35:08
*/

SET FOREIGN_KEY_CHECKS=0;

-- ----------------------------
-- Table structure for `user`
-- ----------------------------
DROP TABLE IF EXISTS `user`;
CREATE TABLE `user` (
  `userid` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `username` varchar(191) NOT NULL,
  `password` longtext NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `token` longtext NOT NULL,
  PRIMARY KEY (`userid`),
  UNIQUE KEY `username` (`username`)
) ENGINE=InnoDB AUTO_INCREMENT=4 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of user
-- ----------------------------
INSERT INTO `user` VALUES ('1', 'zhanglei', '87a38998227cbbc23dcad51cd7f76ab2', '2022-06-13 10:52:13.126', '2022-06-13 10:52:13.126', '0000-00-00 00:00:00.000', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.vuf-9fTfYSAjn0SyOh-55vG6csUSJJsE6UaQHDKoJus');
INSERT INTO `user` VALUES ('2', 'devin', 'e10adc3949ba59abbe56e057f20f883e', '2022-06-13 10:54:13.569', '2022-06-13 10:54:13.569', '0000-00-00 00:00:00.000', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.vuf-9fTfYSAjn0SyOh-55vG6csUSJJsE6UaQHDKoJus');
INSERT INTO `user` VALUES ('3', 'test', 'e10adc3949ba59abbe56e057f20f883e', '2022-06-13 10:54:34.419', '2022-06-13 10:54:34.419', '0000-00-00 00:00:00.000', 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.vuf-9fTfYSAjn0SyOh-55vG6csUSJJsE6UaQHDKoJus');

-- ----------------------------
-- Table structure for `video`
-- ----------------------------
DROP TABLE IF EXISTS `video`;
CREATE TABLE `video` (
  `video_id` bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  `user_id` bigint(20) unsigned NOT NULL,
  `title` longtext NOT NULL,
  `created_at` bigint(20) DEFAULT NULL,
  `play_url` longtext,
  `cover_url` longtext,
  PRIMARY KEY (`video_id`),
  KEY `fk_user_videos` (`user_id`),
  CONSTRAINT `fk_user_videos` FOREIGN KEY (`user_id`) REFERENCES `user` (`userid`)
) ENGINE=InnoDB AUTO_INCREMENT=11 DEFAULT CHARSET=utf8;

-- ----------------------------
-- Records of video
-- ----------------------------
INSERT INTO `video` VALUES ('1', '1', 'my first', '1655084877', '3a47609f70dde7424a72635f0a3ac2be.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('2', '1', 'my second', '1655090257', '7de56f312545cde64d43e790c06acc81.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('3', '1', 'my third', '1655094370', '9ddecd33ec57f06b624caa7a62e58f0a.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('4', '1', 'my fourth', '1655094488', '54edcc6ea75a5317e8d7fdd053001341.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('5', '2', 'hi one', '1655094555', '92ea485fbedc3b294ab5308a47b30934.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('6', '2', 'hi two', '1655094590', '761feeb8f2adf6ddec787091ad615091.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('7', '2', 'hi three', '1655094680', '5889f16269f5cc7e7da432e5a14ffa14.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('8', '3', 'hey man1', '1655094712', '6889aed365f7470edfff41ed3bb7a978.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('9', '3', '22222', '1655094752', '087619f1697d3708944ca2028bd53236.mp4', 'bear.jpg');
INSERT INTO `video` VALUES ('10', '3', '3333', '1655094785', '558404624fd279d368781193fb58b840.mp4', 'bear.jpg');
