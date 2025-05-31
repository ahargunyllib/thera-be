CREATE TABLE hospital_partners (
	id VARCHAR(26) PRIMARY KEY,
	from_hospital_id INT NOT NULL REFERENCES hospitals(id) ON DELETE CASCADE,
	to_hospital_id INT NOT NULL REFERENCES hospitals(id) ON DELETE CASCADE,
	partner_type INT NOT NULL DEFAULT(1),
	status INT NOT NULL DEFAULT(1),
	reason TEXT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TRIGGER update_hospital_partners_timestamp
BEFORE UPDATE ON hospital_partners
FOR EACH ROW
EXECUTE FUNCTION update_timestamp();

CREATE UNIQUE INDEX idx_hospital_partners_unique
ON hospital_partners (
	from_hospital_id,
	to_hospital_id,
	partner_type
);
