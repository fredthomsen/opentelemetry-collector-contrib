# SNMP Receiver

This receiver fetches metrics from a device running a SNMP agent.

Supported pipeline types: `metrics`

> :construction: This receiver is in **Alpha**. Configuration fields and metric data model are subject to change.

## Prerequisites

This receiver has been built to support the following SNMP versions:

- v3
- v2c

## Configuration


| Parameter | Default | Type | Notes |
| --- | --- | --- | --- |
| version |  | String | SNMP version |
| targets |  | List[String] | IP Address + Port |
| community |  | String | SNMP community string |
| oids |  | List[String] | SNMP oids to collect |
| username |  | String | Required |
| password |  | String | Required |
| collection_interval | 2m | Duration | This receiver collects metrics on an interval. Valid time units are `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h` |

### Example Configuration

```yaml
receivers:
  snmp:
    hosts: [device1:161, device2:161]
    oids: []
    collection_interval: 5m
```

## Metrics

Details about the metrics produced by this receiver can be found in [metadata.yaml](./metadata.yaml) with further documentation in [documentation.md](./documentation.md)
