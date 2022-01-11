CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL UNIQUE
);

CREATE TABLE movements (
    id SERIAL PRIMARY KEY,
    user_id INT NOT NULL,
    description TEXT NOT NULL,
    CONSTRAINT fk_user FOREIGN KEY(user_id) REFERENCES users(id)
);

INSERT INTO users (name) VALUES ('Joana Andrade Lopes');
INSERT INTO movements (user_id, description) VALUES (1, 'tolipas');