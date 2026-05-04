CREATE TABLE IF NOT EXISTS users (
    userid SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    ca TIMESTAMP DEFAULT NOW(),
    usertag VARCHAR(255),
    admin BOOLEAN DEFAULT FALSE,
    status TIMESTAMP DEFAULT NOW()
);

CREATE TABLE IF NOT EXISTS messages (
    id SERIAL PRIMARY KEY,
    useridmessage INT REFERENCES users(userid) ON DELETE CASCADE,
    roomidmessage INT NOT NULL,
    CAmessage TIMESTAMP DEFAULT NOW(),
    message TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS private_chats (
    id SERIAL PRIMARY KEY,
    user1_id INT NOT NULL REFERENCES users(userid) ON DELETE CASCADE,
    user2_id INT NOT NULL REFERENCES users(userid) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT NOW(),
    UNIQUE(user1_id, user2_id)
);