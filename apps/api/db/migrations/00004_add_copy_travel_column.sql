-- +goose Up

ALTER TABLE travels
  ADD COLUMN copied_from_travel_id BIGINT UNSIGNED NULL,
  ADD INDEX idx_travels_copied_from_travel_id (copied_from_travel_id),
  ADD CONSTRAINT fk_travels_copied_from_travel
    FOREIGN KEY (copied_from_travel_id) REFERENCES travels (id) ON DELETE SET NULL;

-- +goose Down

ALTER TABLE travels DROP FOREIGN KEY fk_travels_copied_from_travel;
ALTER TABLE travels DROP INDEX idx_travels_copied_from_travel_id;
ALTER TABLE travels DROP COLUMN copied_from_travel_id;
