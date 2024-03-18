ALTER SYSTEM SET max_connections = '200';
ALTER SYSTEM SET shared_buffers = '8GB';
ALTER SYSTEM SET effective_cache_size = '24GB';
ALTER SYSTEM SET maintenance_work_mem = '2GB';
ALTER SYSTEM SET checkpoint_completion_target = '0.9';
ALTER SYSTEM SET wal_buffers = '16MB';
ALTER SYSTEM SET default_statistics_target = '100';
ALTER SYSTEM SET random_page_cost = '1.1';
ALTER SYSTEM SET effective_io_concurrency = '200';
ALTER SYSTEM SET work_mem = '10485kB';
ALTER SYSTEM SET min_wal_size = '1GB';
ALTER SYSTEM SET max_wal_size = '4GB';
ALTER SYSTEM SET max_worker_processes = '16';
ALTER SYSTEM SET max_parallel_workers_per_gather = '4';
ALTER SYSTEM SET max_parallel_workers = '16';
ALTER SYSTEM SET max_parallel_maintenance_workers = '4';

CREATE SCHEMA IF NOT EXISTS message_sender_service_schema;

CREATE TABLE IF NOT EXISTS message_sender_service_schema.messages
(
    id                  UUID PRIMARY KEY,
    content                  VARCHAR(300),
    recipient_phone_number   VARCHAR(20),
    send_status              VARCHAR(10),
    messageId                UUID,
    creation_time            TIMESTAMP,
    send_time                TIMESTAMP
);

CREATE INDEX IF NOT EXISTS index_send_status_from_messages on message_sender_service_schema.messages (send_status);

CREATE TABLE IF NOT EXISTS message_sender_service_schema.config
(
    id                  UUID PRIMARY KEY,
    name               VARCHAR(30),
    status             VARCHAR(3)
);



-- INSERT messages
INSERT INTO message_sender_service_schema.messages (id, content, recipient_phone_number, send_status, messageId, creation_time, send_time)
VALUES
('f6bc843d-47ee-4a48-b0d3-c8729057e384', 'Reminder: Complete your tasks.', '+1666555444', 'SENT', NULL, '2024-03-17 19:30:00', NULL),
('5e1db6c5-4db5-4912-8e76-f32ed152fae5', 'Important: Review attached document.', '+1555666777', 'NOT_SENT', NULL, '2024-03-17 20:00:00', NULL),
('d25e1d91-7db4-4b29-9c45-42f16b9ab4d9', 'New message: Check your inbox.', '+1444333222', 'SENT', NULL, '2024-03-17 21:15:00', NULL),
('2c7bb013-b7e3-4fcf-9512-8da4f5e8f76b', 'Reminder: Tomorrow is the deadline.', '+1333222111', 'SENT', NULL, '2024-03-17 22:30:00', NULL),
('ef4be131-27e6-4c85-951f-9d50e76f4159', 'Notification: Update your profile.', '+1222111333', 'NOT_SENT', NULL, '2024-03-17 23:45:00', NULL),
('7361f2e3-8447-41c0-a2c3-5412d65f28c7', 'New message: Please respond.', '+1111000000', 'SENT', NULL, '2024-03-18 00:00:00', NULL),
('bca8f509-4b9c-479d-b214-062779f09da8', 'Reminder: Do not forget the meeting.', '+1000111100', 'SENT', NULL, '2024-03-18 01:30:00', NULL),
('f4a7e94e-8010-46c2-9f49-62d4f5d1aaf5', 'Your account needs attention.', '+1999888777', 'NOT_SENT', NULL, '2024-03-18 02:00:00', NULL),
('f9ac17ed-1b4b-41f1-9c24-37d76b52cb71', 'Hello, how are you?', '+1234567890', 'SENT', NULL, '2024-03-17 09:00:00', NULL),
('4c2f859f-baf1-44cf-a037-9a0e1832a1c4', 'This is a test message.', '+1987654321', 'SENT', NULL, '2024-03-17 09:15:00', NULL),
('e3b4fc2f-89c1-4a50-af77-1d29b7cf2d57', 'Reminder: Meeting tomorrow.', '+1555555555', 'NOT_SENT', NULL, '2024-03-17 10:30:00', NULL),
('c841f4cb-5b2b-4824-8159-6c02c0ac1cc0', 'Do not forget to submit your report.', '+1777777777', 'SENT', NULL, '2024-03-17 11:45:00', NULL),
('35bbabfc-3c63-4632-99b4-8f016e1b73aa', 'Important: Deadline approaching.', '+1666666666', 'SENT', NULL, '2024-03-17 12:00:00', NULL),
('e45b062e-774e-44da-9e3b-8464c80a42e5', 'You have a new message.', '+1888888888', 'SENT', NULL, '2024-03-17 13:30:00', NULL),
('95fe7fa9-fdf9-4e48-b155-5500e7b16833', 'Meeting reminder: 3 PM today.', '+1222222222', 'NOT_SENT', NULL, '2024-03-17 14:00:00', NULL),
('c78e6432-24ff-4e21-82d4-2e75d1aa5f8a', 'Test message for validation.', '+1333333333', 'SENT', NULL, '2024-03-17 15:15:00', NULL),
('6e74c83f-9a02-4892-9135-72a8a7c9ac7d', 'Reminder: Pay your bills.', '+1444444444', 'SENT', NULL, '2024-03-17 16:30:00', NULL),
('8e1e5a78-fb92-434f-85d2-7316664b7193', 'Your appointment is tomorrow.', '+1999999999', 'NOT_SENT', NULL, '2024-03-17 17:45:00', NULL),
('cc5e7a39-272f-4d85-b4cb-0bfcac4b7a61', 'You have a new notification.', '+1888777666', 'SENT', NULL, '2024-03-17 18:00:00', NULL),
('81c815c6-d4cc-4a3f-9161-41540196d907', 'Hey, what are you doing?', '+13567116', 'SENT', NULL, '2024-03-17 18:00:00', NULL);

INSERT INTO message_sender_service_schema.config (id, name, status) VALUES ('7d3c5c76-d41b-4f3d-95d5-7d3fbef92dc1', 'AUTO-SEND-FEATURE', 'ON');


---- Create app user for customer with just select,insert,add privileges ---
CREATE USER app_user WITH PASSWORD 'f9oOlr7x43IZ9HchxwB1JdVjY5a9KejK';
REVOKE USAGE ON SCHEMA public FROM app_user;
GRANT USAGE ON SCHEMA message_sender_service_schema TO app_user;
REVOKE all privileges on all tables in schema message_sender_service_schema from app_user;
GRANT SELECT, INSERT, UPDATE ON ALL TABLES IN SCHEMA message_sender_service_schema TO app_user;

RESET ROLE;





















