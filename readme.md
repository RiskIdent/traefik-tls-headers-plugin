# Traefik TLS headers plugin

[![Main workflow](https://github.com/RiskIdent/traefik-tls-headers-plugin/actions/workflows/main.yml/badge.svg)](https://github.com/RiskIdent/traefik-tls-headers-plugin/actions/workflows/main.yml)
[![Go matrix workflow](https://github.com/RiskIdent/traefik-tls-headers-plugin/actions/workflows/go-cross.yml/badge.svg)](https://github.com/RiskIdent/traefik-tls-headers-plugin/actions/workflows/go-cross.yml)

## Usage

This plugin will take TLS information from the client connection and write them to some headers.

```yaml
middlewares:
  my-middleware:
    plugin:
      tlsheaders:
        headers:
          cipher: X-TLS-Cipher
```

### Configuration

Traefik static configuration must define the module name (as is usual for Go packages).

The following declaration (given here in YAML) defines a plugin:

<details open><summary>File (YAML)</summary>

```yaml
# Static configuration

experimental:
  plugins:
    tlsheaders:
      moduleName: github.com/RiskIdent/traefik-tls-headers-plugin
      version: v0.1.0
```

</details>

<details><summary>CLI</summary>

```bash
# Static configuration

--experimental.plugins.tlsheaders.moduleName=github.com/RiskIdent/traefik-tls-headers-plugin
--experimental.plugins.tlsheaders.version=v0.1.0
```

</details>


<details><summary>Kubernetes</summary>

```yaml
# Dynamic configuration

apiVersion: traefik.io/v1alpha1
kind: Middleware
metadata:
  name: my-middleware
spec:
  plugin:
    tlsheaders:
      headers:
        cipher: X-TLS-Cipher
```

</details>

### Test locally

In order to test the plugin locally, start the printheaders application:

```bash
make start_headers_reader
```

Then start Traefik with the plugin:

```bash
make testcontainer
```

The traefik test configuration is located in the testconfig directory.

And finally, make a request to the Traefik instance:

```bash
curl https://localhost -k
```

The response should contain the header(s) you set up.

```
Host: localhost
User-Agent: curl/7.81.0
Accept: */*
X-Forwarded-For: 127.0.0.1
X-Forwarded-Host: localhost
X-Forwarded-Port: 443
X-Forwarded-Proto: https
X-Forwarded-Server: ri-t-0940
X-Real-Ip: 127.0.0.1
X-Tls-Cipher: TLS_AES_128_GCM_SHA256
Accept-Encoding: gzip
```
