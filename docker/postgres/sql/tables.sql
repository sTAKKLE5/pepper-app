-- Plants Table
CREATE TABLE plants (
    id SERIAL PRIMARY KEY,
	species VARCHAR(255) NOT NULL CHECK (species IN ('annuum', 'chinense', 'baccatum', 'pubescens', 'frutescens')),
    cultivar VARCHAR(255),
    planting_date TIMESTAMP,
    is_cross BOOLEAN
);

-- Crosses Table
CREATE TABLE crosses (
    id SERIAL PRIMARY KEY,
    plant_id INTEGER REFERENCES plants(id) ON DELETE SET NULL,
    mother_id INTEGER REFERENCES plants(id) ON DELETE SET NULL,
    father_id INTEGER REFERENCES plants(id) ON DELETE SET NULL,
    generation INTEGER
);

-- Logs Table
CREATE TABLE notes (
    id SERIAL PRIMARY KEY,
    plant_id INTEGER REFERENCES plants(id) ON DELETE CASCADE,
    height INTEGER,
    leaves INTEGER,
    buds INTEGER,
    notes TEXT
);
