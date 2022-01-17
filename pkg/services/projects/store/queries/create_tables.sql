CREATE TYPE tags AS ENUM ('UNKNOWN', 'LANGUAGE', 'ARCHITECTURE');

CREATE TABLE projects (
    id VARCHAR PRIMARY KEY DEFAULT gen_random_uuid(),
    name TEXT NOT NULL UNIQUE,
    description TEXT NOT NULL,
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    deleted_at INT NOT NULL
);

CREATE TABLE projects_tags (
    id VARCHAR PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id VARCHAR NOT NULL,
    tag tags NOT NULL,
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id)
);

CREATE TABLE code_files (
    id VARCHAR PRIMARY KEY DEFAULT gen_random_uuid(),
    project_id VARCHAR NOT NULL,
    name TEXT NOT NULL,
    content TEXT NOT NULL,
    created_at INT NOT NULL,
    updated_at INT NOT NULL,
    deleted_at INT NOT NULL,
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id)
);

INSERT INTO projects(id, name, description, created_at, updated_at, deleted_at) VALUES ('ea916764-d2fa-4004-804b-eeded947fe6e', 'project_v1', 'Project v1', 0, 0, 0);
INSERT INTO projects_tags(project_id, tag) VALUES ('ea916764-d2fa-4004-804b-eeded947fe6e', 'LANGUAGE'), ('ea916764-d2fa-4004-804b-eeded947fe6e', 'ARCHITECTURE');
INSERT INTO code_files(project_id, name, content, created_at, updated_at, deleted_at) VALUES ('ea916764-d2fa-4004-804b-eeded947fe6e', 'file_1', 'test 1', 0, 0, 0), ('ea916764-d2fa-4004-804b-eeded947fe6e', 'file_2', 'test 2', 0, 0, 0);

INSERT INTO projects(id, name, description, created_at, updated_at, deleted_at) VALUES ('9a1ab3d8-8093-4617-8284-a626d64186d4', 'project_v2', 'Project v2', 0, 0, 0);
INSERT INTO projects_tags(project_id, tag) VALUES ('9a1ab3d8-8093-4617-8284-a626d64186d4', 'UNKNOWN');
INSERT INTO code_files(project_id, name, content, created_at, updated_at, deleted_at) VALUES ('9a1ab3d8-8093-4617-8284-a626d64186d4', 'file_2', 'test 2', 0, 0, 0), ('9a1ab3d8-8093-4617-8284-a626d64186d4', 'file_3', 'test 3', 0, 0, 0);