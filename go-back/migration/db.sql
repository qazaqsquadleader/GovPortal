CREATE TABLE user(
		user_id  INTEGER PRIMARY KEY AUTOINCREMENT,
		user_name TEXT UNIQUE,
		user_age INTEGER,
        user_status text 	
);

CREATE TABLE user_sessions(
    token text PRIMARY KEY,
    expiresAt TEXT,
    user_id INTEGER,
    FOREIGN KEY (user_id) REFERENCES user(user_id)
);
