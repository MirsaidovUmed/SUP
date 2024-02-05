CREATE TABLE Roles
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE Statuses
(
    id   SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL
);

CREATE TABLE users
(
    id          SERIAL PRIMARY KEY,
    first_name  VARCHAR(255) NOT NULL,
    second_name VARCHAR(255) NOT NULL,
    email       VARCHAR(255) NOT NULL,
    password    VARCHAR(255) NOT NULL,
    role_id     INT REFERENCES Roles (id),
    created_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE projects
(
    id          SERIAL PRIMARY KEY,
    name        VARCHAR(255) NOT NULL,
    description TEXT         NOT NULL,
    status_id   INTEGER REFERENCES statuses (id),
    manager_id  INTEGER REFERENCES users (id),
    start_date  DATE,
    dead_line   DATE
);

CREATE TABLE projects_participant
(
    id             SERIAL PRIMARY KEY,
    participant_id INTEGER REFERENCES users (id),
    project_id     INTEGER REFERENCES projects (id)
);

CREATE TABLE tasks
(
    id            SERIAL PRIMARY KEY,
    title         VARCHAR(255) NOT NULL,
    description   TEXT         NOT NULL,
    status_id     INT REFERENCES Statuses (id),
    controller_id INTEGER REFERENCES users (id),
    executor_id   INTEGER REFERENCES users (id),
    project_id    INTEGER REFERENCES projects (id),
    start_date    DATE,
    dead_line     DATE
);

INSERT INTO Roles (name)
VALUES ('Product Owner'),
       ('Project Manager'),
       ('Developer'),
       ('Tester'),
       ('Analyst');

INSERT INTO Statuses (name)
VALUES ('New'),
       ('In Progress'),
       ('In Review'),
       ('Done'),
       ('On Hold'),
       ('Cancelled');
