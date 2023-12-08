# scan.py (simple BLE scanner demo using Bleak)

import asyncio
from bleak import BleakScanner as Scanner

async def main():
    print('scan devices ...')
    devices = await Scanner.discover()
    for device in devices:
        print(device)

asyncio.run(main())
