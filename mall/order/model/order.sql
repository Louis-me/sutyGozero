CREATE TABLE `order` (
	`id` INT(11) NOT NULL AUTO_INCREMENT,
	`productname` TEXT NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	`price` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	`unit` VARCHAR(50) NULL DEFAULT NULL COLLATE 'utf8_general_ci',
	PRIMARY KEY (`id`) USING BTREE
)
COLLATE='utf8_general_ci'
ENGINE=InnoDB
;
