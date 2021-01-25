CREATE TABLE IF NOT EXISTS projects(
    id SERIAL,
    name varchar(255) NOT NULL,
    description text,
    PRIMARY KEY (id)
);
CREATE TABLE IF NOT EXISTS columns(
    id SERIAL,
    name varchar(255) NOT NULL,
    project_id int NOT NULL,
    position int NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (project_id) references projects(id) on delete cascade on update cascade
);
CREATE TABLE IF NOT EXISTS tasks(
    id SERIAL,
    name varchar(255) NOT NULL,
    description text,
    column_id int NOT NULL,
    position int NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (column_id) references columns(id) on delete cascade on update cascade
);
CREATE TABLE IF NOT EXISTS comments(
    id SERIAL,
    description text,
    task_id int NOT NULL,
    position int NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (task_id) references tasks(id) on delete cascade on update cascade
    );