CREATE TABLE taskStatus (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL
);

CREATE TABLE task (
    id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    status_id BIGINT NOT NULL,
    user_id BIGINT,
    FOREIGN KEY (status_id) REFERENCES TaskStatus(id),
    FOREIGN KEY (user_id) REFERENCES "user"(id)
);

INSERT INTO taskStatus (name)
VALUES ('Ожидает'), ('Выполнена'), ('В работе'), ('Отменена');