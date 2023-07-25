CREATE TABLE IF NOT EXISTS `rdev_helm_repository` (
    `created_at`                    datetime DEFAULT NULL,
    `updated_at`                    datetime DEFAULT NULL,
    `id`                            varchar(64)  NOT NULL,
    `name`                          varchar(256) NOT NULL,
    `username`                      varchar(256) NOT NULL,
    `password`                      varchar(256) NOT NULL,
    `url`                           varchar(256) NOT NULL,
    `is_active`                     tinyint(1) DEFAULT '1',
    `cert_file`                     longtext DEFAULT NULL,
    `key_file`                      longtext DEFAULT NULL,
    `ca_file`                      longtext DEFAULT NULL,
    `insecure_skip_tls_verify`      boolean DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
);