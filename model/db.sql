create database if not exists `db_apiserver` default character set utf8 ;

# 选中库
use `db_apiserver`;

# 查询是否存在该数据库
DROP table if exists `tb_users`;

create table `tb_user` (
    `id` int(20) unsigned not null auto_increment,
    `username` varchar(255) not null ,
    `password` varchar(255) not null ,
    `state` int(10) default '1' comment '用户状态 0默认 1删除',
    `createdAt` timestamp null default null,
    `updatedAt` timestamp null default null,
    `deletedAt` timestamp null default null,
    primary key (`id`),
    unique key `username` (`username`),
    key `idx_tb_users_deletedAt` (`deletedAt`)
)ENGINE=MyISAM AUTO_INCREMENT=2 DEFAULT CHARSET=utf8 comment '用户表'
/*!40101 SET character_set_client = @saved_cs_client */;

--
-- Dumping data for table `tb_users`
--

LOCK TABLES `tb_users` WRITE;
/*!40000 ALTER TABLE `tb_users` DISABLE KEYS */;
INSERT INTO `tb_users` VALUES (0,'admin','$2a$10$veGcArz47VGj7l9xN7g2iuT9TF21jLI1YGXarGzvARNdnt4inC9PG','2018-05-27 16:25:33','2018-05-27 16:25:33',NULL);
/*!40000 ALTER TABLE `tb_users` ENABLE KEYS */;
UNLOCK TABLES;