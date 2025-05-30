CREATE TABLE IF NOT EXISTS moods (
		id SERIAL PRIMARY KEY,
		doctor_id VARCHAR(36) NOT NULL REFERENCES doctors(id) ON DELETE CASCADE,
		scale INT NOT NULL, -- Mood scale from 1 to 5
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX unique_user_per_day ON moods (doctor_id, (DATE(created_at)));
