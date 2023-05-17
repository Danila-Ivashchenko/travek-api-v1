USE travek_db;

CREATE TABLE country_timezones(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    country_id BIGINT NOT NULL,
    timezone_id BIGINT NOT NULL
);

CREATE TABLE relation(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_country BIGINT NOT NULL,
    second_country BIGINT NOT NULL,
    possibility_to_stay_forever BOOL NOT NULL,
    internal_passport BOOL NOT NULL
);

CREATE TABLE timezone(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    zone SMALLINT NOT NULL
);

CREATE TABLE country(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    language VARCHAR(255) NOT NULL,
    description TEXT
);

CREATE TABLE road(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    first_country BIGINT NOT NULL,
    second_country BIGINT NOT NULL,
    transport VARCHAR(255) NOT NULL,
    time_hours BIGINT NOT NULL
);

CREATE TABLE tags(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE country_tags(
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT PRIMARY KEY,
    country_id BIGINT UNSIGNED NOT NULL,
    tag_id BIGINT UNSIGNED NOT NULL
);

INSERT INTO `country_tags` VALUES (4,7,1),(5,7,8),(6,7,9),(7,8,1),(8,8,2),(9,8,8),(10,8,9),(11,9,2),(12,9,6),(13,9,7),(14,10,1),(15,10,9),(16,8,10),(17,11,1),(18,11,2),(19,11,5),(20,11,10);
INSERT INTO `country` VALUES (7,'Russia','Russian','Russia (Russian Federation) is the largest state in the world, occupying 1/6 of the entire land area of the Earth.This is a country with a long history, rich culture and generous nature.Some monuments of architecture, history and culture on the territory of the Russian Federation are included in the list of UNESCO World Heritage Sites.'),(8,'Japan','Japanese','Japan, located in East Asia, is known for its unique blend of ancient traditions and modern innovations. Its archipelago consists of four main islands and numerous smaller ones, all surrounded by beautiful seas. Japan is renowned for its rich cultural heritage, including traditional arts such as calligraphy and flower arrangement, as well as its cuisine, which includes sushi, ramen, and many other delicious dishes'),(9,'China','Chinese','China, located in East Asia, is the world\'s most populous country with a rich history and culture that dates back thousands of years. It is known for its iconic landmarks such as the Great Wall of China, the Forbidden City, and the Terracotta Army. China is also famous for its cuisine, which includes dishes such as Peking duck, dumplings, and hot pot.'),(10,'Canada','English','Canada is a North American country known for its vast landscapes, diverse culture, and friendly people. It is the second-largest country in the world by land area and has a population of over 37 million people. Canada is known for its natural beauty, including the Rocky Mountains, Niagara Falls, and the Northern Lights. The country is also famous for its multiculturalism, with a mix of Indigenous, French, British.'),(11,'USA','English','Canada is a North American country known for its vast landscapes, diverse culture, and friendly people. It is the second-largest country in the world by land area and has a population of over 37 million people. Canada is known for its natural beauty, including the Rocky Mountains, Niagara Falls, and the Northern Lights. The country is also famous for its multiculturalism, with a mix of Indigenous, French, British.');
INSERT INTO `country_timezones` VALUES (10,3,42),(11,3,43),(12,3,44),(13,3,45),(14,3,46),(15,3,47),(16,3,48),(17,3,50),(18,3,51),(19,4,42),(20,4,43),(21,4,44),(22,4,45),(23,4,46),(24,4,47),(25,4,48),(26,4,50),(27,4,51),(28,5,42),(29,5,43),(30,5,44),(31,5,45),(32,5,46),(33,5,47),(34,5,48),(35,5,50),(36,5,51),(37,6,33),(38,6,34),(39,6,35),(40,6,36),(41,7,42),(42,7,43),(43,7,44),(44,7,45),(45,7,46),(46,7,47),(47,7,48),(48,7,50),(49,7,51),(50,8,49),(51,9,45),(52,9,46),(53,9,47),(54,9,48),(55,10,32),(56,10,33),(57,10,34),(58,11,30),(59,11,31),(60,11,32),(61,11,33),(62,11,34),(63,11,35);
INSERT INTO `relation` VALUES (1,7,11,0,1),(2,7,8,1,1),(3,7,10,0,1),(4,7,9,0,1);
INSERT INTO `road` VALUES (1,7,8,'plain',8),(2,7,11,'plain',12),(3,7,10,'plain',12),(4,7,9,'plain',5),(5,7,9,'train',32),(6,7,9,'car',30);
INSERT INTO `tags` VALUES (1,'beautiful nature'),(2,'high technology'),(3,'extreme tourism'),(4,'sea'),(5,'warm climate'),(6,'historical'),(7,'delicious food'),(8,'mountains'),(9,'relax'),(10,'ocean');
INSERT INTO `timezone` VALUES (28,-12),(29,-11),(30,-10),(31,-9),(32,-8),(33,-7),(34,-6),(35,-5),(36,-4),(37,-3),(38,-2),(39,-1),(40,0),(41,1),(42,2),(43,3),(44,4),(45,5),(46,6),(47,7),(48,8),(49,9),(50,10),(51,11),(52,12),(53,13),(54,14);
