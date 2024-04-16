CREATE DATABASE IF NOT EXISTS booking;

USE booking;

CREATE TABLE IF NOT EXISTS room
(
    id PRIMARY KEY AUTO_INCREMENT,
	name VARCHAR(100),
	capacity INT,
);

CREATE TABLE IF NOT EXISTS reservation
(
    id PRIMARY KEY AUTO_INCREMENT,
	id_salle INT,
	time_start  VARCHAR(5),
	time_end  VARCHAR(5),
	date_start  VARCHAR(10),
	date_end  VARCHAR(10),

);

INSERT INTO room
VALUES (1, "Colza", 100, False);
