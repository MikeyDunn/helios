CREATE TABLE `accounts` (
  `account_id` int(11) unsigned NOT NULL,
  `username` varchar(32) NOT NULL,
  `created` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `modified` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
  PRIMARY KEY (`account_id`),
  KEY `idx_username` (`username`)
);

INSERT INTO `accounts` (`account_id`, `username`) VALUES (1, 'justinbieber');
INSERT INTO `accounts` (`account_id`, `username`) VALUES (2, 'taylorswift');

CREATE TABLE `account_follower_counts` (
  `account_id` int(11) unsigned NOT NULL,
  `date` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `follower_count` int(11) unsigned NOT NULL,
  `following_count` int(11) unsigned NOT NULL,
  `post_count` int(11) unsigned NOT NULL,
  UNIQUE KEY `account_id_date_idx` (`account_id`,`date`)
);

CREATE TABLE `account_post_stats` (
  `account_id` int(11) unsigned NOT NULL,
  `post_id` bigint(20) unsigned NOT NULL,
  `post_date` datetime NOT NULL,
  `comment_count` int(11) unsigned NOT NULL,
  `like_count` int(11) unsigned NOT NULL,
  UNIQUE KEY `account_id` (`account_id`,`post_id`)
);

CREATE TABLE `account_engagement_aggregations` (
  `account_id` int(11) unsigned NOT NULL,
  `date` date NOT NULL,
  `avg_comments` int(11) unsigned NOT NULL,
  `avg_likes` int(11) unsigned NOT NULL,
  PRIMARY KEY (`account_id`),
  UNIQUE KEY `account_id` (`account_id`,`date`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
