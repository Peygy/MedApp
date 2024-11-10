CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS doctors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    doctor_name VARCHAR(100) NOT NULL,
    specialization VARCHAR(100) NOT NULL,
    experience_years INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

/*INSERT INTO doctors (doctor_name, specialization, experience_years)
VALUES 
    ('Иван Иванов', 'Терапевт', 10),
    ('Мария Петрова', 'Хирург', 15),
    ('Алексей Смирнов', 'Кардиолог', 8),
    ('Ольга Кузнецова', 'Педиатр', 5),
    ('Дмитрий Сидоров', 'Невролог', 12);*/