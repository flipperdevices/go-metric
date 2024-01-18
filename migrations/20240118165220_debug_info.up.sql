CREATE TABLE "debug_info_metric"
(
    uuid     UUID,
    platform Int32,
    time     DateTime,
    app_version Nullable(String),
    user_session Nullable(UUID),

    key      String default '',
    value    String default ''
) Engine = MergeTree() ORDER BY (time)

--migration:split