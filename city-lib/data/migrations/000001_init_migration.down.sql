ALTER TABLE rentals 
DROP INDEX rentals_books_lookup;

DROP TABLE IF EXISTS rentals;
DROP TABLE IF EXISTS books;
