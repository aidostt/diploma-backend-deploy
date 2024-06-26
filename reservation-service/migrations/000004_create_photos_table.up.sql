CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
CREATE TABLE IF NOT EXISTS photos (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    restaurantID UUID REFERENCES restaurants(id) ON DELETE CASCADE,
    url TEXT NOT NULL UNIQUE
);
