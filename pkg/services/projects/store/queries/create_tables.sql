CREATE TABLE projects (
    id SERIAL PRIMARY KEY,
    name VARCHAR(200) NOT NULL UNIQUE,
    description VARCHAR(500) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE code_files (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    name VARCHAR(200) NOT NULL,
    content VARCHAR(100000) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id)
);

CREATE TABLE tags (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL
);

CREATE TABLE projects_tags (
    id SERIAL PRIMARY KEY,
    project_id INT NOT NULL,
    tag_id INT NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    CONSTRAINT fk_project FOREIGN KEY(project_id) REFERENCES projects(id),
    CONSTRAINT fk_tag FOREIGN KEY(tag_id) REFERENCES tags(id)
);

CREATE INDEX project_tags_project_idx ON projects_tags(project_id);
CREATE INDEX project_tags_tag_idx ON projects_tags(tag_id);

INSERT INTO tags(name, created_at, updated_at) VALUES ('LANGUAGE', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), ('ARCHITECTURE', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO projects(name, description, created_at, updated_at) VALUES ('project_v1', 'Project v1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO projects_tags(project_id, tag_id, created_at, updated_at) VALUES (1, 1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (1, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO code_files(project_id, name, content, created_at, updated_at) VALUES (1, 'file_1', 'test 1', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (1, 'file_2', 'test 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);

INSERT INTO projects(name, description, created_at, updated_at) VALUES ('project_v2', 'Project v2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO projects_tags(project_id, tag_id, created_at, updated_at) VALUES (2, 2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);
INSERT INTO code_files(project_id, name, content, created_at, updated_at) VALUES (2, 'file_2', 'test 2', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP), (2, 'file_3', 'test 3', CURRENT_TIMESTAMP, CURRENT_TIMESTAMP);