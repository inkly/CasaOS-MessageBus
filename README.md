# CasaOS-MessageBus

The message bus of the **inkly distribution of CasaOS**, a maintained release of the project after upstream [IceWhaleTech/CasaOS](https://github.com/IceWhaleTech/CasaOS) stopped shipping in 2025. This repository descends from IceWhale's through [alvins82's fork](https://github.com/alvins82/CasaOS-MessageBus), whose one commit here fixed the setup script's fallback on Ubuntu 26.

CasaOS services do not call each other to report what happened. They publish it here once, and whoever is listening hears it.

## What it does

An **event** is something that already happened: a source ID, a name, and a flat map of string properties. `POST /v2/message_bus/event/{source_id}/{name}` publishes one. An **action** is the same shape pointed the other way — a request that a service do something — published at `POST /v2/message_bus/action/{source_id}/{name}`. Both are declared first as types (`/v2/message_bus/event_type`, `/v2/message_bus/action_type`) so a subscriber can ask for them by name. Nothing in this distribution publishes or subscribes to actions today; the routes exist because the API is symmetric.

Who publishes:

- **casaos** POSTs `casaos:system:utilization` every 5 seconds over loopback — CPU, memory, network, disks — plus `casaos:file:operate` while a file operation runs.
- **local-storage** publishes `local-storage:disk:added` and its siblings, built from kernel uevents as drives appear and disappear.
- **app-management** publishes `app:*` events — install progress, start, stop, errors — over the unix socket at `/tmp/message-bus.sock` rather than the TCP port.

Who subscribes: the dashboard, over socket.io at `/v2/message_bus/socket.io/`. Its CPU, network and disk widgets redraw on each `casaos:system:utilization`; app cards follow `app:apply-changes-begin`, `app:start-end` and the rest. A plain websocket subscription to one source is also available at `GET /v2/message_bus/event/{source_id}`.

Delivery is fire and forget. An event nobody is subscribed to is dropped, not queued, and there is no replay — the 5-second telemetry makes a missed frame cheap.

The service also owns the YSK cards on the dashboard — task and notice cards — served at `/v2/message_bus/ysk`. That is the only state it keeps of its own.

Requests carry a CasaOS JWT. The unix socket, loopback peers and websocket upgrades skip it.

## Where things live

| Path | What |
|---|---|
| `/etc/casaos/message-bus.conf` | config, written from the built-in sample on first start |
| `/var/run/casaos/message-bus.db` | registered event and action types; runtime only, and services re-register on start |
| `/var/lib/casaos/db/message-bus.db` | YSK cards and settings, kept across reboots |
| `/var/run/casaos/message-bus.url` | the TCP address the gateway and other services read |
| `/tmp/message-bus.sock` | the unix socket, same API |
| `/var/log/casaos/message-bus.log` | log |

The service picks a free loopback port at start, registers `/v2/message_bus` and `/doc/v2/message_bus` with the gateway, and writes its address to the file above. The API specification is [`api/message_bus/openapi.yaml`](api/message_bus/openapi.yaml), also served at `/doc/v2/message_bus`.

## Install

Components are not installed one at a time. The distribution ships as one bundle:

```sh
curl -fsSL https://github.com/inkly/CasaOS-Install/releases/latest/download/install.sh | sudo bash
```

What a release contains, and how it is built, is described in [CasaOS-Install](https://github.com/inkly/CasaOS-Install#readme).

## What this fork changed

**The access log no longer records host-side event publishes.** casaos's 5-second telemetry produced one JSON access-log line per publish, which systemd hands to journald: about 17 000 lines a day that buried every real request. The telemetry rate is unchanged — it is the dashboard's only realtime feed — and the log line is skipped instead. The skip is narrow: POST on the event publish route, and only from the unix socket or a loopback peer. Websocket subscribes, event-type registrations, actions and anything from the LAN stay logged ([CasaOS #2211](https://github.com/IceWhaleTech/CasaOS/issues/2211)).

Packaging changed too, without touching the service: the release workflow builds and publishes the tarballs here instead of calling upstream's pipeline, which needed credentials this fork does not have and so never produced a release.

## Development

Go 1.20 or later. The generated API code in `codegen/` is committed, so a plain checkout builds:

```sh
go build ./...
go test ./...
```

After editing `api/message_bus/openapi.yaml`, run `go generate` to regenerate `codegen/`; the release workflow does the same before building.

## Licence

Apache License 2.0 — see [LICENSE](LICENSE). The service is the work of IceWhale and its contributors; alvins82's setup-script fix for Ubuntu 26 is in this history as well. CasaOS is a mark of IceWhale, used here to say what this is a release of.
