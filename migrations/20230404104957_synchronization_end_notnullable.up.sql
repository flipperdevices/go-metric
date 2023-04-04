alter table metric.synchronization_end
    drop column if exists changes_count;

--migration:split

alter table metric.synchronization_end
    add column changes_count Int32 default -2;