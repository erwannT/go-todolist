CREATE SCHEMA TODOLIST;

USE TODOLIST;


-- ************************************** `tasks`

CREATE TABLE `tasks`
(
 `id`    MEDIUMINT NOT NULL AUTO_INCREMENT ,
 `task` TEXT NOT NULL ,
 `date` datetime NOT NULL,
 `done` bool,

PRIMARY KEY (`id`)
);



