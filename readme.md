# NFACU

[![GitMoji](https://img.shields.io/badge/Gitmoji-%F0%9F%8E%A8%20-FFDD67.svg)](https://gitmoji.dev)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)
![Lines Of Code](https://img.shields.io/tokei/lines/github.com/UltiRequiem/nfacu?color=blue&label=Total%20Lines)
[![Go Report Card](https://goreportcard.com/badge/github.com/UltiRequiem/nfacu)](https://goreportcard.com/report/github.com/UltiRequiem/nfacu)

> .NET Framework App Config Updater

Example Configuration file

```json
[
  {
    "path": "C:\\repos\\gato\\src\\GATO.Listener\\App.config",
    "settings": {
      "DialerEventTopic.Name": "dialer-events-topic-elias",
      "Dialers": "Noble-1, Genesys-1",
      "Dialers:Genesys-1:Campaigns": "TC1, AMA, HMA",
      "Groups:Extensionless": "1, 16, 40"
    }
  },
  {
    "path": "C:\\repos\\gato\\src\\GATO.Noble.Worker\\App.config",
    "settings": {
      "DialerEventTopic.Name": "dialer-events-topic-elias"
    }
  },
  {
    "path": "C:\\repos\\gato\\src\\GATO.Genesys.Worker\\App.config",
    "settings": {
      "DialerEventTopic.Name": "dialer-events-topic-elias"
    }
  }
]
```

Then run

```
nfacu
```

That will modify the properties of the `App.config` files in the array with the
specified settings.

By default nfacu search for a `nfacu.json` file, but it can also receive another
file as parameter

Example

```
nfacu hulk.json
```

Check [`DotnetSample/`](./DotnetSample/) for a better understanding of this
tool.

Made at request of [@neosaile](https://github.com/neosaile).

## Install

Using

```sh
go install github.com/UltiRequiem/nfacu@latest
```

Or use a binary from
[releases](https://github.com/UltiRequiem/nfacu/releases/latest).

## License

Released under the MIT license.
