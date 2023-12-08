# Python BLE Scan Sample

Simple Python sample demonstrating BLE scanning based on the Python Bleak
library. The sample code can be found in sample/scan.py


# Quick Start

For a quick start you need to work in a BASH shell, which is natural
in a Linux or MacOS environment. There is good news for Windows 10++ users
(https://learn.microsoft.com/en-us/windows/wsl/install), which has now also
BASH shell support.

In the repository's root directory (where .git is located) source the
BASH script `go`

```
   $ . go   # which is the short form for `$ source go` (equivalentl)
```

Bash script `go` will:

1) create virtual Python environment (`venv` folder)
2) activate the virtual environment
3) update the Python package installer (`pip`)
4) install Python package `bleak` (a Python BLE library)

After the ´. go´ command has finished, your virtual Python environment is ready
for running the BLE scanner sample:

```
   $ python sample/scan.py
```


# Python Code

```
# scan.py (simple BLE scanner demo using Bleak)

import asyncio
from bleak import BleakScanner as Scanner

async def main():
    print('scan devices ...')
    devices = await Scanner.discover()
    for device in devices:
        print(device)

asyncio.run(main())
```

# Additional Resources

The code is based on a tutorial blog about utilizing the `Bleak` library:
`How to Connect to a Bluetooth Device with a MacBook and Python`.
https://medium.com/@protobioengineering/how-to-connect-to-a-bluetooth-device-with-a-macbook-and-python-7a14ece6a780
