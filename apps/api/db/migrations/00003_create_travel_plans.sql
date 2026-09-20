-- +goose Up

CREATE TABLE travel_plans (
    id BIGINT UNSIGNED NOT NULL AUTO_INCREMENT,
    travel_id BIGINT UNSIGNED NOT NULL,
    date DATE NOT NULL,
    place VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    sort_order INT NOT NULL,
    created_at DATETIME NOT NULL,
    updated_at DATETIME NOT NULL,
    PRIMARY KEY (id),
    INDEX idx_travel_plans_travel_id (travel_id),
    FOREIGN KEY (travel_id) REFERENCES travels(id) ON DELETE CASCADE
);

-- +goose Down

DROP TABLE travel_plans;
