CREATE TABLE role
(
    id   serial       not null unique primary key,
    name varchar(255) not null
);

CREATE TABLE "user"
(
    id        serial       not null unique primary key,
    full_name varchar(50)  not null,
    login     varchar(255) not null unique,
    password  varchar(255) not null unique,
    role_id   int          not null,

    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
            REFERENCES role (id)
);

INSERT INTO role (name)
VALUES ('admin');
INSERT INTO role (name)
VALUES ('manager');
INSERT INTO role (name)
VALUES ('worker');

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Linus Torvalds', 'admin', '$2a$10$mMj3EqKXqv73Kspkipw/K.jysW3MTwxkKSRSY6qxhdXmsZAnBzxx.', 1);

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Mr.Beast', 'manager1', '$2a$10$dAdNLd1fWgoXQv.f3DOHw.GSB./fdzTa8mdkPPf3bKYS6nc69W8bO', 2);

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Worker 1', 'worker1', '$2a$10$MfIw7qhpklAsl3/GYZAuc.Vu1zinwYPsPy/OeOxssHGS1N7OXiZyW', 3);
INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Worker 2', 'worker2', '$2a$10$EQBZ4/zj8qPFMGwqZsi55eVsGE1IwyhOwPyhOXzKeqdWjP550O2v2', 3);