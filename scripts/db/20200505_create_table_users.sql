USE users_db;

CREATE TABLE `users_db`.`users` (
    `id` BIGINT(20) UNSIGNED AUTO_INCREMENT NOT NULL,
    `first_name` VARCHAR(50) NULL,
    `last_name` VARCHAR(50) NULL,
    `email` VARCHAR(50) NOT NULL,
    `date_created` VARCHAR(50) NULL,
    PRIMARY KEY (`id`),
    UNIQUE INDEX `email_UNIQUE` (`email` ASC)
);
