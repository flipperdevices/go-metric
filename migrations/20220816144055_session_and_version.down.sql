alter table metric.open
    drop column if exists user_session;

--migration:split

alter table metric.open
    drop column if exists app_version;

--migration:split

alter table metric.flipper_gatt_info
    drop column if exists user_session;

--migration:split

alter table metric.flipper_gatt_info
    drop column if exists app_version;


--migration:split

alter table metric.flipper_rpc_info
    drop column if exists user_session;

--migration:split

alter table metric.flipper_rpc_info
    drop column if exists app_version;


--migration:split

alter table metric.synchronization_end
    drop column if exists user_session;

--migration:split

alter table metric.synchronization_end
    drop column if exists app_version;


--migration:split

alter table metric.update_flipper_end
    drop column if exists user_session;

--migration:split

alter table metric.update_flipper_end
    drop column if exists app_version;


--migration:split

alter table metric.update_flipper_start
    drop column if exists user_session;

--migration:split

alter table metric.update_flipper_start
    drop column if exists app_version;

