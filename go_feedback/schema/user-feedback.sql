CREATE SCHEMA IF NOT EXISTS rizon_db;

-- Table: role
CREATE TABLE rizon_db.user
(
    id        BIGSERIAL PRIMARY KEY,
    userId    UUID UNIQUE  NOT NULL DEFAULT gen_random_uuid(),
    name      VARCHAR(100) NOT NULL,
    email     VARCHAR(100) NOT NULL,
    deviceId  VARCHAR(100),
    status    VARCHAR(100),
    createdAt TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    isDeleted BOOLEAN      NOT NULL DEFAULT false,
    deletedAt TIMESTAMP,
    CONSTRAINT unique_user_id_per_email UNIQUE (userId, email)
);


CREATE TABLE rizon_db.feedback
(
    id            BIGSERIAL PRIMARY KEY,
    userId        UUID UNIQUE  NOT NULL REFERENCES rizon_db.user (userId),
    comments      VARCHAR(255) NOT NULL,
    isLeaveReview BOOLEAN      NOT NULL DEFAULT false,
    isEnjoying    BOOLEAN      NOT NULL DEFAULT false,
    createdAt     TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT fk_feedback_user
        FOREIGN KEY (userId)
            REFERENCES rizon_db.user (userId)
            ON DELETE CASCADE
)
