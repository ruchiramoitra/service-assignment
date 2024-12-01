
CREATE TABLE IF NOT EXISTS services (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT
);


CREATE TABLE IF NOT EXISTS versions (
    id SERIAL PRIMARY KEY,
    service_id INT REFERENCES services(id) ON DELETE CASCADE,
    name VARCHAR(100) NOT NULL
);


DO $$ 
DECLARE
    service_id INT;
BEGIN
    FOR i IN 1..5 LOOP
        INSERT INTO services(name, description) 
        VALUES ('Service ' || i, 'Description: ' || i)
        RETURNING id INTO service_id;
        
        FOR j IN 1..3 LOOP
            INSERT INTO versions(service_id, name)
            VALUES (service_id, 'Version ' || i || '.' || j);
        END LOOP;
    END LOOP;
END $$;
