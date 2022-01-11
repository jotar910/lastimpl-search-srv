CREATE TYPE tags AS ENUM ('UNKNOWN', 'LANGUAGE', 'ARCHITECTURE');

CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL
);

CREATE TABLE projects_tags (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    tag tags NOT NULL,
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id)
);

CREATE TABLE code_files (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    name TEXT NOT NULL,
    content TEXT NOT NULL,
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id)
);

INSERT INTO projects(name, description) VALUES ('project_v1', 'Project v1');
INSERT INTO projects_tags(project_id, tag) VALUES (1, 'LANGUAGE');
INSERT INTO projects_tags(project_id, tag) VALUES (1, 'ARCHITECTURE');
INSERT INTO code_files(project_id, name, content) VALUES (1, 'file_1', 'test 1');
INSERT INTO code_files(project_id, name, content) VALUES (1, 'file_2', 'test 2');

INSERT INTO projects(name, description) VALUES ('project_v2', 'Project v2');
INSERT INTO projects_tags(project_id, tag) VALUES (2, 'UNKNOWN');
INSERT INTO code_files(project_id, name, content) VALUES (2, 'file_1', 'test 1');
INSERT INTO code_files(project_id, name, content) VALUES (2, 'file_2', 'test 2');
INSERT INTO code_files(project_id, name, content) VALUES (2, 'file_3', 'test 3');