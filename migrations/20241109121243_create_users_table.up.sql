-- 创建扩展以支持 UUID 生成
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- 创建 users 表
CREATE TABLE users (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    telp_number VARCHAR(255),
    email VARCHAR(255) UNIQUE NOT NULL,
    password VARCHAR(255),
    role VARCHAR(50),
    image_url VARCHAR(255),
    is_verified BOOLEAN DEFAULT FALSE,
    balance BIGINT DEFAULT 0,
    create_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    oauth_provider VARCHAR(50),
    oauth_id VARCHAR(255),
    oauth_token TEXT,
    oauth_refresh_token TEXT,
    oauth_token_expiry TIMESTAMP
);
