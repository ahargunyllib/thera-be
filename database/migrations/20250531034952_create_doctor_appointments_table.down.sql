DROP TRIGGER IF EXISTS update_doctor_appointment_timestamp ON doctor_appointments;

DROP INDEX IF EXISTS idx_doctor_appointments_unique;

DROP TABLE IF EXISTS doctor_appointments;
