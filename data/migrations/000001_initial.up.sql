CREATE TABLE IF NOT EXISTS users (
    id          SERIAL NOT NULL PRIMARY KEY,
    telegram_id VARCHAR(30) NOT NULL UNIQUE,
    chat_id     VARCHAR(30) NOT NULL,
    username    VARCHAR(30) NOT NULL,
    created_at  TIMESTAMP NOT NULL DEFAULT now(),
    updated_at  TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS chats (
    id                   SERIAL NOT NULL PRIMARY KEY,
    chat_id              VARCHAR(30) NOT NULL UNIQUE,
    current_character_id INTEGER,
    status               INTEGER,
    add_substatus        INTEGER,
    update_substatus     INTEGER,
    created_at           TIMESTAMP NOT NULL DEFAULT now(),
    updated_at           TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS characters (
    id           SERIAL NOT NULL PRIMARY KEY,
    owner        VARCHAR(30) NOT NULL,
    name         VARCHAR(30) NOT NULL,
    char_level   INTEGER,
    current_hp   INTEGER,
    max_hp       INTEGER,
    armor        INTEGER,
    gold_coins   INTEGER,
    silver_coins INTEGER,
    copper_coins INTEGER,
    created_at   TIMESTAMP NOT NULL DEFAULT now(),
    updated_at   TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS abilities (
    id              SERIAL NOT NULL PRIMARY KEY,
    character_owner INTEGER,
    strength        INTEGER,
    dexterity       INTEGER,
    constitution    INTEGER,
    intelligence    INTEGER,
    wisdom          INTEGER,
    charisma        INTEGER,
    created_at      TIMESTAMP NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS skills_insights(
    id              SERIAL NOT NULL PRIMARY KEY,
    character_owner INTEGER,
    acrobatics      INTEGER,
    animalHandling  INTEGER,
    arcana          INTEGER,
    athletics       INTEGER,
    deception       INTEGER,
    history         INTEGER,
    insight         INTEGER,
    intimidation    INTEGER,
    investigation   INTEGER,
    medicine        INTEGER,
    nature          INTEGER,
    perception      INTEGER,
    performance     INTEGER,
    persuasion      INTEGER,
    religion        INTEGER,
    sleight_of_hand INTEGER,
    stealth         INTEGER,
    survival        INTEGER,
    created_at      TIMESTAMP NOT NULL DEFAULT now(),
    updated_at      TIMESTAMP NOT NULL DEFAULT now()
);
