CREATE TABLE doctor_appointments (
	id VARCHAR(26) PRIMARY KEY,
	doctor_id VARCHAR(36) NOT NULL REFERENCES doctors(id) ON DELETE CASCADE,
	patient_id VARCHAR(36) NOT NULL REFERENCES patients(id) ON DELETE CASCADE,
	appointment_date DATE NOT NULL,
	start_time TIME NOT NULL,
  end_time TIME NOT NULL,
	type INT NOT NULL,
	status INT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX idx_doctor_appointments_unique
ON doctor_appointments (
	doctor_id,
	appointment_date,
	start_time,
	end_time
);

CREATE TRIGGER update_doctor_appointment_timestamp
BEFORE UPDATE ON doctor_appointments
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
