alter table metric.synchronization_end
    add column changes_count Nullable(Int32) default null;