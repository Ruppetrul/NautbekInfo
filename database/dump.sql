CREATE TABLE IF NOT EXISTS user_visits (
    visit_ip VARCHAR(40),
    visit_date VARCHAR(20) NOT NULL,
    visit_count INTEGER
);

ALTER TABLE user_visits ADD UNIQUE (visit_date, visit_ip);
