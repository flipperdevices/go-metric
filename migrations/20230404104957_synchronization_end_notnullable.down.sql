alter table metric.synchronization_end
    drop column if exists changes_count;

--migration:split

alter table metric.synchronization_end
    add column changes_count Nullable(Int32) default null;