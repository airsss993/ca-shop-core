CREATE TABLE users
(
    id         uuid PRIMARY KEY,
    name       varchar(50) NOT NULL,
    password   text        NOT NULL,
    created_at timestamp   NOT NULL DEFAULT now()
);

CREATE TABLE carts
(
    user_id     uuid PRIMARY KEY,
    total_price int       NOT NULL DEFAULT 0,
    updated_at  timestamp NOT NULL DEFAULT now(),

    FOREIGN KEY (user_id) REFERENCES users (id)
);

CREATE TABLE cart_items
(
    user_id  uuid        NOT NULL,
    sku      varchar(50) NOT NULL,
    price    int         NOT NULL,
    quantity int         NOT NULL DEFAULT 1,
    added_at timestamp   NOT NULL DEFAULT now(),

    PRIMARY KEY (user_id, sku),
    FOREIGN KEY (user_id) REFERENCES carts (user_id)
);