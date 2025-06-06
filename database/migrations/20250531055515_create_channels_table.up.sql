CREATE TABLE channels (
    id VARCHAR(36) PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
		doctor_id VARCHAR(36) NOT NULL REFERENCES doctors(id) ON DELETE CASCADE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_channels_timestamp
BEFORE UPDATE ON channels
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();
