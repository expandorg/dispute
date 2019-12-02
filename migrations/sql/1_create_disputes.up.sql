CREATE TABLE `disputes` (
  `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
  `dispute_message` varchar(255) NOT NULL DEFAULT '',
  `resolution_messageCopy` varchar(255) NOT NULL DEFAULT '',
  `status` varchar(40) NOT NULL DEFAULT '',
  `active` tinyint(4) NOT NULL DEFAULT '1',
  `response_id` int(11) unsigned NOT NULL,
  `worker_id` int(11) unsigned NOT NULL,
  `task_id` int(11) unsigned NOT NULL,
  `job_id` int(11) unsigned NOT NULL,
  `score_id` int(11) unsigned NOT NULL,
  `verifier_id` int(11) unsigned NOT NULL,
  `created_at` TIMESTAMP NOT NULL DEFAULT NOW(),
	`updated_at` TIMESTAMP NOT NULL DEFAULT NOW() ON UPDATE NOW(),
  PRIMARY KEY (`id`)
)