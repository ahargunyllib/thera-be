DROP INDEX IF EXISTS idx_hospital_partners_unique;

DROP TRIGGER IF EXISTS update_hospital_partners_timestamp ON hospital_partners;

DROP TABLE IF EXISTS hospital_partners;
