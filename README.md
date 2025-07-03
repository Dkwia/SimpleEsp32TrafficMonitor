
# 📶 ESP32 Bluetooth PC Traffic Display

This project displays your PC's current network upload/download speed on an ESP32's OLED screen using a Bluetooth Serial connection. Ideal for real-time bandwidth monitoring without cluttering your desktop.

![ESP32 OLED showing upload/download](https://via.placeholder.com/600x300.png?text=ESP32+Bluetooth+Traffic+Monitor)

---

## 🔧 Features

- 📟 Real-time display on 128x64 I2C OLED (SSD1306)
- 🔗 Connects via **Bluetooth Classic (SPP)** — no Wi-Fi needed
- 📊 Shows **upload** and **download** speeds in kB/s
- 🖥️ Cross-platform Go-based sender app (or send manually via terminal)
- 🔒 Optional pairing code for secure connections (default: `1234`)

---

## 🧰 Requirements

### 🖥 PC Side
- Windows 10/11 with Bluetooth SPP support
- Go installed (for sender app) or any serial terminal (PuTTY, etc.)

### 🛠 ESP32 Side
- ESP32 board (with Bluetooth)
- SSD1306 128x64 I2C OLED screen
- Arduino IDE
- Libraries:
  - `BluetoothSerial` (built into ESP32 core)
  - `Adafruit_SSD1306`
  - `Adafruit_GFX`

---

## 🚀 Getting Started

### 1. ⚙️ Flash the ESP32

1. Install the required libraries in the Arduino IDE.
2. Load and upload the sketch from [`sketch.ino`](sketch.ino).
3. OLED will display: `Waiting for PC...`.

### 2. 🔵 Pair via Bluetooth

1. Open **Bluetooth Settings** in Windows.
2. Scan and pair with: **ESP32TrafficDisplay**.
3. Enter PIN code: `1234`.

> ℹ️ If Windows shows “Paired” but not “Connected”, proceed to the next step.

### 3. 🧪 Send Data from PC

#### Option A: Use Go app

```sh
go run traffic_bt.go -port=COM7
