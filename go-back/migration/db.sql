CREATE TABLE IF NOT EXISTS user(
		userId  INTEGER PRIMARY KEY AUTOINCREMENT,
		username TEXT UNIQUE,
		password TEXT,
        email TEXT UNIQUE,
        user_status text 	
);

CREATE TABLE IF NOT EXISTS userSessions(
    token text PRIMARY KEY,
    expiresAt TEXT,
    userid INTEGER,
    FOREIGN KEY (userid) REFERENCES user(userId)
);

CREATE TABLE IF NOT EXISTS studentList{
    StudentId int PRIMARY KEY AUTOINCREMENT,
    StudentFirstName text,
    StudentLastName text,
    StudentTrainingCourse text
};