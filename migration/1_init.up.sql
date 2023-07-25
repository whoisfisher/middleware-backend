CREATE TABLE IF NOT EXISTS `rdev_user` (
    `created_at`    datetime DEFAULT NULL,
    `updated_at`    datetime DEFAULT NULL,
    `id`            varchar(64)  NOT NULL,
    `name`          varchar(256) NOT NULL,
    `password`      varchar(256) NOT NULL,
    `is_active`     tinyint(1) DEFAULT '1',
    `is_admin`      tinyint(1) DEFAULT '1',
    `type`          varchar(256) NOT NULL,
    `role`          int NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
);

CREATE TABLE  IF NOT EXISTS  rdev_instance (
    `id`                    varchar(64) NOT NULL,
    `name`                  varchar(64) NOT NULL,
    `type`                  varchar(64) NOT NULL,
    `cluster_id`            varchar(64) NOT NULL,
    `template_id`           varchar(64) NOT NULL,
    `namespace`             varchar(64) NOT NULL,
    `count`                 int         NOT NULL,
    `request_cpu`           varchar(64) NOT NULL,
    `request_memory`        varchar(64) NOT NULL,
    `limit_cpu`             varchar(64) NOT NULL,
    `limit_memory`          varchar(64) NOT NULL,
    `volume`                varchar(64) NOT NULL,
    `status`                varchar(64) NOT NULL,
    `created_at`            datetime    DEFAULT NULL,
    `updated_at`            datetime    DEFAULT NULL,
    PRIMARY KEY (`id`)
);

CREATE TABLE  IF NOT EXISTS  rdev_user_instance (
    `user_id`               varchar(64) NOT NULL,
    `instance_id`           varchar(64) NOT NULL,
    UNIQUE KEY `user_instance_key` (`user_id`, `instance_id`)
);

CREATE TABLE IF NOT EXISTS `rdev_cluster`
(
    `created_at`            datetime     DEFAULT NULL,
    `updated_at`            datetime     DEFAULT NULL,
    `id`                    varchar(255) NOT NULL,
    `name`                  varchar(255) NOT NULL,
    `api_server`            varchar(255) DEFAULT NULL,
    `version`               varchar(255) DEFAULT NULL,
    `token`                 mediumtext   DEFAULT NULL,
    `kube_config`           longtext     DEFAULT NULL,
    `type`                  varchar(255) DEFAULT NULL,
    `status`                varchar(255) DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `name` (`name`)
    );

CREATE TABLE IF NOT EXISTS rdev_user_cluster
(
    `user_id`               varchar(64) NOT NULL,
    `cluster_id`            varchar(64) NOT NULL,
    UNIQUE KEY `user_instance_key` (`user_id`, `cluster_id`)
);

CREATE TABLE IF NOT EXISTS rdev_templates
(
    `created_at`            datetime     DEFAULT NULL,
    `updated_at`            datetime     DEFAULT NULL,
    `id`                    varchar(64) NOT NULL PRIMARY KEY,
    `name`                  varchar(64) NOT NULL,
    `base_template`         longtext NOT NULL,
    `advance_template`      longtext NOT NULL
);