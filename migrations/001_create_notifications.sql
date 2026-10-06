CREATE TABLE notifications (

    id SERIAL PRIMARY KEY,

    user_id VARCHAR(100) NOT NULL,

    channel VARCHAR(50) NOT NULL,

    message TEXT NOT NULL,

    status VARCHAR(50) NOT NULL,

    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP

);


