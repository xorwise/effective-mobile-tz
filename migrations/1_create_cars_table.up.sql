CREATE TABLE IF NOT EXISTS cars (
    reg_num TEXT PRIMARY KEY,
    mark TEXT NOT NULL,
    model TEXT NOT NULL,
    year INTEGER NOT NULL,
    owner_name TEXT NOT NULL,
    owner_surname TEXT NOT NULL,
    owner_patronymic TEXT NOT NULL
);
