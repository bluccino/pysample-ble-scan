![Python BLE Scan Sample](https://github.com/bluccino/pysample-ble-scan/assets/17394277/9214608c-f4e5-455c-b0e9-39da8abddd68)


# Python BLE Scan Sample (v1.0.0)

Simple Python sample demonstrating BLE scanning based on the Python Bleak
library. The sample code can be found in sample/scan.py


# Quick Start

For a quick start you need to work in a BASH shell with `git` support and access
to the system's Bluetooth (`Bluez`) stack, which is natural in a Linux or MacOS environment. This tutorial does not include support for Windows platforms, even
users with some Python/Windows experience should still be able to get the
sample running.

## Step 1: Download the Repository

```
   $ git clone https://github.com/bluccino/pysample-ble-scan
```

## Step 2: Change to the Repository's Root Directory

The repository's root directory can be identified as the directory
where `.git` is located.
```
   $ cd pysample-ble-scan   # change to repository's root directory
   $ ls
   .		.git		README.md	sample
   ..		.gitignore	go
```

## Step 3: Get Going

In the repository's root directory source the BASH script `go` (`sourcing`
the script means, to enabled the script for a change of current environment variables)

```
   $ . go      # `dot-space-go`, the short form for `$ source go` (equivalent)
   (venv) $    # prompt `(venv) $` indicates an activated virtual environment 
```

Bash script `go` will:

* create a virtual Python environment (`venv` folder)
* activate the virtual environment (BASH prompt changes to `(venv) $`)
* update the Python package installer (`pip`)
* install Python package `bleak` (a Python BLE library)

## Step 4: Run Sample

After the ´. go´ command has finished, your virtual Python environment is ready
for running the BLE scanner sample:

```
   (venv) $ python sample/scan.py
```

The scanner sample should output lines beginning with Bluetooth addresses of
BLE devices in the vicinity, similar to the text below.

```
   scan devices ...
   8570AC27-4992-7D89-88C1-5C5BCAC07247: None
   7483473F-2962-4927-619E-BDF4628C8930: None
   AF5BD257-F282-9448-037C-DD17E1C16DCB: None
   8E10B2B0-0EDB-1B2C-930C-53513C790A01: Azzurro
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

## [1] How to Connect to a Bluetooth Device with a MacBook and Python

A helpful blog about jump starting Python BLE access on base of Bleak.

https://medium.com/@protobioengineering/how-to-connect-to-a-bluetooth-device-with-a-macbook-and-python-7a14ece6a780


## [2] Reading your phone's battery level over Bluetooth BLE with Python Bleak

Another useful resource demonstrating how to write a Python BLE client for reading
the heart rate of a BLE heart rate server emulated on a smart phone.

https://getwavecake.com/blog/reading-your-phones-battery-level-over-bluetooth-ble-with-python-bleak/
