# v0.0.7

* upd: refactor default plugin handling in prep to support more platforms
* upd: remove darwin build target
* upd: rename example config file
* upd: dependency gopsutil
* upd: support custom api ca cert loading (circonus output)
* fix: include example configuration files
* upd: dest guard for agent version
* fix: env var syntax in example conf

# v0.0.6

* upd: remove agent version tag from internal metrics (to reduce cardinality)
* add: `cua_version` metric

# v0.0.5

* add: windows default plugin placeholder
* add: rollup tag to internal metrics

# v0.0.4

* add: debug_metrics setting
* fix: copy metric for origin and origin instance
* upd: default plugins in example configuration
* fix: ldflags typo

# v0.0.3

* add: metric volume internal metric `cua_metrics_sent`
* upd: cgm debug logging to use info
* add: default plugins go to default check (cpu,mem,disk,diskio,swap,system,kernel,processes,internal)

# v0.0.2

* add: `instance_id` required input plugin setting

# v0.0.1

initial testing release
