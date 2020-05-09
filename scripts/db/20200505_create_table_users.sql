USE users_db;

CREATE TABLE `users_db`.`users` (
    `id` BIGINT(20) UNSIGNED AUTO_INCREMENT NOT NULL,
    `first_name` VARCHAR(50) NULL,
    `last_name` VARCHAR(50) NULL,
    `email` VARCHAR(50) NOT NULL,
    `date_created` DATETIME NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `email_UNIQUE` (`email` ASC)
);

ALTER TABLE `users_db`.`users`
ADD COLUMN `status` VARCHAR(50) NOT NULL AFTER `email`,
ADD COLUMN `password` VARCHAR(32) NOT NULL AFTER `status`;
