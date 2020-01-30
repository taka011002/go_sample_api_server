CREATE TABLE api_request_logs (
    id INTEGER auto_increment PRIMARY KEY,
    user_id VARCHAR(255),
    method VARCHAR(255),
    path VARCHAR(255),
    params VARCHAR(255)
);