#include "BluetoothSerial.h"
#include <Adafruit_SSD1306.h>

#define SCREEN_WIDTH 128
#define SCREEN_HEIGHT 64
#define OLED_RESET    -1

BluetoothSerial SerialBT;
Adafruit_SSD1306 display(SCREEN_WIDTH, SCREEN_HEIGHT, &Wire, OLED_RESET);

constexpr const char* BT_PIN = "1234";

void sppCallback(esp_spp_cb_event_t evt, esp_spp_cb_param_t *) {
  if (evt == ESP_SPP_SRV_OPEN_EVT) {
    display.clearDisplay();
    display.setCursor(0, 0);
    display.println("PC connected");
    display.println("Waiting data…");
    display.display();
  }
}


void setup() {
  Serial.begin(115200);
  delay(100);
  if (!SerialBT.begin("ESP32TrafficDisplay", /*isMaster=*/false)) {
    Serial.println("Bluetooth init failed");
    while (true) delay(1000);
  }

  SerialBT.setPin(BT_PIN, 4);
  SerialBT.enableSSP();
  SerialBT.register_callback(sppCallback);

  Serial.println("Bluetooth started. Pair with code 1234");

  if (!display.begin(SSD1306_SWITCHCAPVCC, 0x3C)) {
    Serial.println("SSD1306 allocation failed");
    while (true) delay(1000);
  }
  display.clearDisplay();
  display.setTextSize(1);
  display.setTextColor(SSD1306_WHITE);
  display.setCursor(0, 10);
  display.println("Waiting for PC...");
  display.display();
}

bool shownConnected = false;

void loop() {
  if (!shownConnected && SerialBT.hasClient()) {
    display.clearDisplay();
    display.setCursor(0, 0);
    display.println("PC connected");
    display.println("Waiting data…");
    display.display();
    shownConnected = true;
  }

  if (SerialBT.available()) {
  String line = SerialBT.readStringUntil('\n');
  int c1 = line.indexOf(',');
  int c2 = line.indexOf(',', c1 + 1);
  if (c1 > 0 && c2 > c1) {
    String up   = line.substring(0, c1);
    String down = line.substring(c1 + 1, c2);
    String run  = line.substring(c2 + 1);

    display.clearDisplay();
    display.setCursor(0, 0);
    display.printf("Upl : %s kB/s\n", up.c_str());
    display.printf("Dwl : %s kB/s\n", down.c_str());
    display.printf("Rtm : %s s\n",   run.c_str());
    display.display();
  }
}

}
