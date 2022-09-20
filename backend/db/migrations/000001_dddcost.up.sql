CREATE TABLE IF NOT EXISTS dddcost(
    id SERIAL PRIMARY KEY,
    origin INT,
    destination INT,
    cost DECIMAL(10,2)
    );


INSERT INTO dddcost (origin, destination, cost) VALUES (11, 16, 1.90);
INSERT INTO dddcost (origin, destination, cost) VALUES (16, 11, 2.90);
INSERT INTO dddcost (origin, destination, cost) VALUES (11, 17, 1.70);
INSERT INTO dddcost (origin, destination, cost) VALUES (17, 11, 2.70);
INSERT INTO dddcost (origin, destination, cost) VALUES (11, 18, 0.90);
INSERT INTO dddcost (origin, destination, cost) VALUES (18, 11, 1.90);