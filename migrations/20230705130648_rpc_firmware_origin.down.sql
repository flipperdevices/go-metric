alter table metric.flipper_rpc_info
    drop column if exists firmware_fork_name;

--migration:split

alter table metric.flipper_rpc_info
    drop column if exists firmware_git_url;
