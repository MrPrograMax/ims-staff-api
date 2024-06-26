CREATE TABLE taskStatus (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    statusId BIGINT,
    userId BIGINT,
    FOREIGN KEY (statusId) REFERENCES TaskStatus(id),
    FOREIGN KEY (userId) REFERENCES "user"(id)
);

INSERT INTO taskStatus (name)
VALUES ('Выполнена'), ('В работе'), ('Ожидает'), ('Отменена');