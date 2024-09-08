CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users_health_data (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    userId UUID UNIQUE NOT NULL,
    age INT NOT NULL DEFAULT 0,
    height FLOAT NOT NULL DEFAULT 0.0,
    weight FLOAT NOT NULL DEFAULT 0.0,
    pulse INT NOT NULL DEFAULT 0,
    pressure VARCHAR(255) NOT NULL DEFAULT '',
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);