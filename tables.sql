CREATE TABLE USERS (
    FullName VARCHAR(45) NOT NULL,
    ContactNo VARCHAR(45) NOT NULL,
    UserName VARCHAR(45) NOT NULL UNIQUE,
    PWord VARCHAR(45) NOT NULL,
    PRIMARY KEY (`UserName`)
);