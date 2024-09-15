CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS doctors (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    doctor_name VARCHAR(100) NOT NULL,
    specialization VARCHAR(100) NOT NULL,
    experience_years INT NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

INSERT INTO doctors (doctor_name, specialization, experience_years) 
VALUES 
    ('Иванов Иван Иванович', 'Кардиолог', 15),
    ('Петров Петр Петрович', 'Невролог', 10),
    ('Сидоров Сидор Сидорович', 'Педиатр', 8),
    ('Кузнецов Николай Николаевич', 'Ортопед', 12),
    ('Смирнова Анна Сергеевна', 'Дерматолог', 5)
ON CONFLICT (specialization) DO NOTHING;
