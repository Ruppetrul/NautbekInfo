CREATE TABLE IF NOT EXISTS user_visits (
    visit_ip VARCHAR(40),
    visit_date VARCHAR(20) NOT NULL,
    visit_count INTEGER
);

ALTER TABLE user_visits ADD UNIQUE (visit_date, visit_ip);
ALTER TABLE user_visits ADD COLUMN app VARCHAR(40);

CREATE TABLE IF NOT EXISTS user_feedback (
    visit_ip VARCHAR(40),
    visit_date VARCHAR(20) NOT NULL,
    text TEXT,
    app VARCHAR(40)
);