CREATE TABLE patients (
	id VARCHAR(36) PRIMARY KEY,
	full_name VARCHAR(255) NOT NULL,
	id_number VARCHAR(50) NOT NULL UNIQUE,
	phone_number VARCHAR(20) NULL,
	address TEXT NOT NULL,
	date_of_birth DATE NOT NULL,
	gender INT NOT NULL,
	height DECIMAL(5, 2) NOT NULL, -- Height in cm
	weight DECIMAL(5, 2) NOT NULL, -- Weight in kg
	blood_type INT NOT NULL,
	allergies TEXT NULL,
	medical_record_number VARCHAR(50) NOT NULL UNIQUE,
	hospital_id INT NOT NULL REFERENCES hospitals(id) ON DELETE CASCADE,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_patient_timestamp
BEFORE UPDATE ON patients
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
