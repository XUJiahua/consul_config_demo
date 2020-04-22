
Demo for watching config update

### Local

```
go run .
```

### Remote (Consul)

```
consul agent -dev
consul kv put consul_config_demo/config.json @config.json
```


```
go run . -remote
```
