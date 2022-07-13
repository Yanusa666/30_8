--    Схема БД для информационной системы отслеживания выполнения задач.

DROP TABLE IF EXISTS tasks_labels, tasks, labels, users;

CREATE TABLE users
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE labels
(
    id   SERIAL PRIMARY KEY,
    name TEXT NOT NULL
);

CREATE TABLE tasks
(
    id          SERIAL PRIMARY KEY,
    opened      BIGINT NOT NULL               DEFAULT extract(epoch from now()),
    closed      BIGINT                        DEFAULT 0,
    author_id   INTEGER REFERENCES users (id) DEFAULT 0,
    assigned_id INTEGER REFERENCES users (id) DEFAULT 0,
    title       TEXT,
    content     TEXT
);

CREATE TABLE tasks_labels
(
    task_id  INTEGER REFERENCES tasks (id),
    label_id INTEGER REFERENCES labels (id)
);

INSERT INTO users (id, name)
VALUES (1, 'Anna'),
       (2, 'Ben'),
       (3, 'Coul'),
       (4, 'David');

INSERT INTO tasks (id, author_id, assigned_id, content)
VALUES (1, 1, 2, 'Создавать новые задачи'),
       (2, 2, 1, 'Получать список всех задач'),
       (3, 3, 1, 'Получать список задач по автору'),
       (4, 2, 4, 'Получать список задач по метке'),
       (5, 4, 3, 'Обновлять задачу по id'),
       (6, 3, 2, 'Удалять задачу по id.');

INSERT INTO labels (id, name)
VALUES (1, 'Срочная'),
       (2, 'Желательная'),
       (3, 'Незначительная');

INSERT INTO tasks_labels (task_id,label_id)
VALUES (1, 1),
       (2, 1),
       (3, 2),
       (4, 3),
       (5, 3),
       (6, 3);