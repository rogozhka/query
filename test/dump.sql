
DROP TABLE IF EXISTS `_x_query_test_table1`;
CREATE TABLE IF NOT EXISTS `_x_query_test_table1` (
  `test_uid` char(32) NOT NULL,
  `test_datetime` datetime NOT NULL,
  `test_expired` datetime NOT NULL,
  `test_is_active` enum('Y','N') NOT NULL DEFAULT 'Y',
  `test_number` int(11) NOT NULL,
  `test_text` mediumtext NOT NULL,
  PRIMARY KEY (`test_uid`),
  KEY `test_datetime` (`test_datetime`),
  KEY `test_expired` (`test_expired`),
  KEY `test_is_active` (`test_is_active`)
);


INSERT INTO `_x_query_test_table1` (`test_uid`, `test_datetime`, `test_expired`, `test_is_active`, `test_number`, `test_text`) VALUES
('cron', '2013-09-25 13:39:01', '2013-09-25 13:44:01', 'N', 1, 'cron:text:'),
('memset', '2013-09-25 12:57:02', '2013-09-25 13:57:02', 'Y', 2, 'memset:text:'),
('update8k', '2013-09-24 23:38:16', '2013-09-24 23:43:16', 'N', 3, 'update8k:text:'),
('test5', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 5, 'test5:text:'),
('test6', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 6, 'test6:text:'),
('test7', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 7, 'test7:text:'),
('test8', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 8, 'test8:text:'),
('test9', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 9, 'test9:text:'),
('test10', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 10, 'test10:text:'),
('test11', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 11, 'test11:text:'),
('test12', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 12, 'test12:text:'),
('test13', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 13, 'test13:text:'),
('test14', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 14, 'test14:text:'),
('test15', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 15, 'test15:text:'),
('test16', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 16, 'test16:text:'),
('test17', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 17, 'test17:text:'),
('test18', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 18, 'test18:text:'),
('test19', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 19, 'test19:text:'),
('test20', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 20, 'test20:text:'),
('test21', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 21, 'test21:text:'),
('test22', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 22, 'test22:text:'),
('test23', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 23, 'test23:text:'),
('test24', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 24, 'test24:text:'),
('test25', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 25, 'test25:text:'),
('test26', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 26, 'test26:text:'),
('test27', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 27, 'test27:text:'),
('test28', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 28, 'test28:text:'),
('test29', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 29, 'test29:text:'),
('test30', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 30, 'test30:text:'),
('test31', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 31, 'test31:text:'),
('test32', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 32, 'test32:text:'),
('test33', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 33, 'test33:text:'),
('test34', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 34, 'test34:text:'),
('test35', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 35, 'test35:text:'),
('test36', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 36, 'test36:text:'),
('test37', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 37, 'test37:text:'),
('test38', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 38, 'test38:text:'),
('test39', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 39, 'test39:text:'),
('test40', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 40, 'test40:text:'),
('test41', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 41, 'test41:text:'),
('test42', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 42, 'test42:text:'),
('test43', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 43, 'test43:text:'),
('test44', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 44, 'test44:text:'),
('test45', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 45, 'test45:text:'),
('test46', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 46, 'test46:text:'),
('test47', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 47, 'test47:text:'),
('test48', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 48, 'test48:text:'),
('test49', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 49, 'test49:text:'),
('test50', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 50, 'test50:text:'),
('test51', '2018-01-26 23:17:02', '2018-01-26 23:17:20', 'Y', 51, 'test51:text:');
