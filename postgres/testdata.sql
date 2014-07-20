drop table if exists test;
create table test(
    id BIGINT NOT NULL PRIMARY KEY,
    point BIGINT NOT NULL
);

drop function if exists replace(BIGINT, BIGINT);
create function replace(key BIGINT, data BIGINT) RETURNS VOID AS
$$
BEGIN
    UPDATE test SET point = data WHERE id = key;
    IF found THEN
        RETURN;
    ELSE
        INSERT INTO test(id,point) VALUES (key, data);
    END IF;
END;
$$
LANGUAGE plpgsql;
