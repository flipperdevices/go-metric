alter table metric.flipper_rpc_info
    add column firmware_fork_name Nullable(String) default null;

--migration:split

alter table metric.flipper_rpc_info
    add column firmware_git_url Nullable(String) default null;
