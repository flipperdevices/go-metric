alter table metric.open
    add column user_session Nullable(UUID) default null;

--migration:split

alter table metric.open
    add column app_version Nullable(String) default null;

--migration:split

alter table metric.flipper_gatt_info
    add column user_session Nullable(UUID) default null;

--migration:split

alter table metric.flipper_gatt_info
    add column app_version Nullable(String) default null;


--migration:split

alter table metric.flipper_rpc_info
    add column user_session Nullable(UUID) default null;

--migration:split

alter table metric.flipper_rpc_info
    add column app_version Nullable(String) default null;


--migration:split

alter table metric.synchronization_end
    add column user_session Nullable(UUID) default null;

--migration:split

alter table metric.synchronization_end
    add column app_version Nullable(String) default null;


--migration:split

alter table metric.update_flipper_end
    add column user_session Nullable(UUID) default null;

--migration:split

alter table metric.update_flipper_end
    add column app_version Nullable(String) default null;


--migration:split

alter table metric.update_flipper_start
    add column user_session Nullable(UUID) default null;

--migration:split

alter table metric.update_flipper_start
    add column app_version Nullable(String) default null;