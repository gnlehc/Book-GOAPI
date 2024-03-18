CREATE TABLE Books (
    BookID INT PRIMARY KEY,
    Title VARCHAR(255),
    Author VARCHAR(255),
    ISBN VARCHAR(20)
);

CREATE TABLE Books (
    Title VARCHAR(255),
    Author VARCHAR(255),
    ISBN VARCHAR(20)
);


SELECT * FROM Books;

DROP TABLE IF EXISTS Books;