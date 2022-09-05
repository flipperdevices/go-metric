CREATE TABLE "subghz_provisioning"
(
    uuid            UUID,
    platform        Int32,
    time            DateTime,
    app_version Nullable(String),
    user_session Nullable(UUID),

    region_network  String default '',
    region_sim1     String default '',
    region_sim2     String default '',
    region_ip       String default '',
    region_system   String default '',
    region_provided String default '',
    is_roaming      UInt8  default 0
) Engine = MergeTree() ORDER BY (time)

--migration:split