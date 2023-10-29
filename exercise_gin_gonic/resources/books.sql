CREATE TABLE books (
                       id SERIAL PRIMARY KEY,
                       title VARCHAR(255) NOT NULL,
                       isbn VARCHAR(255) NOT NULL,
                       language VARCHAR(255) NOT NULL,
                       publisher VARCHAR(255) NOT NULL,
                       num_pages INTEGER NOT NULL
);

INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Lord of the Rings', '978-0544003415', 'English', 'Mariner Books', 1216);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Hobbit', '978-0544003415', 'English', 'Mariner Books', 310);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Silmarillion', '978-0544003415', 'English', 'Mariner Books', 488);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Children of Hurin', '978-0544003415', 'English', 'Mariner Books', 352);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Fall of Gondolin', '978-0544003415', 'English', 'Mariner Books', 352);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Unfinished Tales', '978-0544003415', 'English', 'Mariner Books', 480);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The History of Middle-earth', '978-0544003415', 'English', 'Mariner Books', 480);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Lord of the Rings: A Readers Companion', '978-0544003415', 'English', 'Mariner Books', 480);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Harry Potter & The Philosophers Stone', '978-0747532699', 'English', 'Bloomsbury', 223);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Harry Potter & The Chamber of Secrets', '978-0747538483', 'English', 'Bloomsbury', 251);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Harry Potter & The Prisoner of Azkaban', '978-0747546243', 'English', 'Bloomsbury', 317);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Harry Potter & The Goblet of Fire', '978-0747546243', 'English', 'Bloomsbury', 636);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Harry Potter & The Order of the Phoenix', '978-0747546243', 'English', 'Bloomsbury', 766);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Harry Potter & The Half-Blood Prince', '978-0747546243', 'English', 'Bloomsbury', 607);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Harry Potter & The Deathly Hallows', '978-0747546243', 'English', 'Bloomsbury', 607);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Hitchhikers Guide to the Galaxy', '978-0345391803', 'English', 'Pan Books', 224);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('The Restaurant at the End of the Universe', '978-0345391803', 'English', 'Pan Books', 224);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Life, the Universe and Everything', '978-0345391803', 'English', 'Pan Books', 224);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('So Long, and Thanks for All the Fish', '978-0345391803', 'English', 'Pan Books', 224);
INSERT INTO books (title, isbn, language, publisher, num_pages) VALUES ('Mostly Harmless', '978-0345391803', 'English', 'Pan Books', 224);