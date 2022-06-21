create table if not exists app_open
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime
)
    engine = MergeTree ORDER BY uuid
        SETTINGS index_granularity = 8192;

create table if not exists experimental_open_fm
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists experimental_open_screenstreaming
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists flipper_gatt_info
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime,
    flipper_version
        String
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists flipper_rpc_info
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime,
    sd_card_is_available
        UInt8,
    internal_free_bytes
        Int64,
    internal_total_bytes
        Int64,
    external_free_bytes
        Int64,
    external_total_bytes
        Int64
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists open_edit
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists open_emulate
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists open_save_key
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists open_share
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists synchronization_end
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime,
    sub_ghz_count
        Int32,
    rfid_count
        Int32,
    nfc_count
        Int32,
    infrared_count
        Int32,
    i_button_count
        Int32,
    synchronization_time_ms
        Int64
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists update_flipper_end
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime,
    update_from
        String,
    update_to
        String,
    update_id
        Int64,
    update_status
        Int32
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;

create table if not exists update_flipper_start
(
    uuid
        UUID,
    platform
        Int32,
    time
        DateTime,
    update_from
        String,
    update_to
        String,
    update_id
        Int64
)
    engine = MergeTree ORDER BY time
        SETTINGS index_granularity = 8192;