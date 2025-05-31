CREATE TABLE doctor_schedules (
		id SERIAL PRIMARY KEY,
		doctor_id VARCHAR(36) NOT NULL REFERENCES doctors(id) ON DELETE CASCADE,
		day_of_week INT NOT NULL, -- 1= Monday, 2= Tuesday, ..., 7= Sunday
		start_time TIME NOT NULL,
		end_time TIME NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_doctor_schedule_timestamp
BEFORE UPDATE ON doctor_schedules
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
