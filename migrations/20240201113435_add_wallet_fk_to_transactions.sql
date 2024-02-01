-- +goose Up
-- +goose StatementBegin
ALTER TABLE transactions
ADD COLUMN wallet_id BIGINT NOT NULL,
ADD CONSTRAINT fk_wallet
FOREIGN KEY (wallet_id) REFERENCES wallets(id) ON DELETE CASCADE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE transactions
DROP FOREIGN KEY fk_wallet;
DROP COLUMN wallet_id;
-- +goose StatementEnd
