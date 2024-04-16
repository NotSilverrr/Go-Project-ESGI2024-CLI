CREATE DATABASE IF NOT EXISTS booking;

USE booking;

CREATE TABLE IF NOT EXISTS room (
    id INT PRIMARY KEY AUTO_INCREMENT,
    name VARCHAR(100),
    capacity INT
);

CREATE TABLE IF NOT EXISTS reservation (
    id INT PRIMARY KEY AUTO_INCREMENT,
    id_salle INT,
    time_start VARCHAR(5),
    time_end VARCHAR(5),
    date_start VARCHAR(10),
    date_end VARCHAR(10),
    FOREIGN KEY (id_salle) REFERENCES room(id)
);

INSERT INTO room(id, name, capacity)
VALUES (1, "Colza", 100);