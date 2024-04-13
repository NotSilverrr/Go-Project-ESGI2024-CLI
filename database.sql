CREATE DATABASE IF NOT EXISTS booking;

USE booking;

CREATE TABLE IF NOT EXISTS salles
(
    id PRIMARY KEY
	name VARCHAR(100),
	capacity INT,
	booked BOOLEAN DEFAULT False 
);

INSERT INTO salles
VALUES (1, "Colza", 100, False);

INSERT INTO salles
VALUES (2, "Eolienne", 30, False);