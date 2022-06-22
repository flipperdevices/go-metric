CREATE TABLE "app_open"
(
    uuid     UUID,
    platform Int32,
    time     DateTime,
    target   Int32
) Engine = MergeTree() ORDER BY (time)

--migration:split

CREATE TABLE "flipper_gatt_info"
(
    uuid            UUID,
    platform        Int32,
    time            DateTime,
    flipper_version String
) Engine = MergeTree() ORDER BY (time)

--migration:split

CREATE TABLE "flipper_rpc_info"
(
    uuid                 UUID,
    platform             Int32,
    time                 DateTime,
    sd_card_is_available UInt8,
    internal_free_bytes  Int64,
    internal_total_bytes Int64,
    external_free_bytes  Int64,
    external_total_bytes Int64
) Engine = MergeTree() ORDER BY (time)

--migration:split

CREATE TABLE "synchronization_end"
(
    uuid                    UUID,
    platform                Int32,
    time                    DateTime,
    sub_ghz_count           Int32,
    rfid_count              Int32,
    nfc_count               Int32,
    infrared_count          Int32,
    i_button_count          Int32,
    synchronization_time_ms Int64
) Engine = MergeTree() ORDER BY (time)

--migration:split

CREATE TABLE "update_flipper_end"
(
    uuid          UUID,
    platform      Int32,
    time          DateTime,
    update_from   String,
    update_to     String,
    update_id     Int64,
    update_status Int32
) Engine = MergeTree() ORDER BY (time)

--migration:split

CREATE TABLE "update_flipper_start"
(
    uuid        UUID,
    platform    Int32,
    time        DateTime,
    update_from String,
    update_to   String,
    update_id   Int64
) Engine = MergeTree() ORDER BY (time)
