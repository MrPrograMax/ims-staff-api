CREATE TABLE role
(
    id          serial      not null unique primary key,
    name       varchar(255) not null
);

CREATE TABLE "user"
(
    id            serial       not null unique primary key,
    full_name     varchar(50)  not null,
    login         varchar(255) not null unique,
    password      varchar(255) not null unique,
    role_id int not null,

    CONSTRAINT fk_role
        FOREIGN KEY (role_id)
            REFERENCES role(id)
);

INSERT INTO role (name) VALUES ('admin');
INSERT INTO role (name) VALUES ('manager');
INSERT INTO role (name) VALUES ('worker');

INSERT INTO "user" (full_name, login, password, role_id)
    VALUES ('Linus Torvalds', 'admin', '123', 1);

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Mr.Beast', 'manager1', '345', 2);

INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Cat1', 'worker1', '678', 3);
INSERT INTO "user" (full_name, login, password, role_id)
VALUES ('Cat2', 'worker2', '912', 3);
