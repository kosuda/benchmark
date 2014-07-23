drop table if exists test;
create table test(
    id VARCHAR(12) NOT NULL PRIMARY KEY,
    point BIGINT NOT NULL
) ENGINE=MEMORY;
