#include <WiFi.h>
#include <OneWire.h>
#include <DallasTemperature.h>

// WiFi credentials
const char* ssid = "Olbring";
const char* password = "PASSWORD";
const char* hostname = "esp32-temperatur-sensor"; // Change this to your desired hostname

// Data wire is connected to GPIO 4
#define ONE_WIRE_BUS 4

OneWire oneWire(ONE_WIRE_BUS);
DallasTemperature sensors(&oneWire);

WiFiServer server(80);

void setup() {
  Serial.begin(115200);
  delay(10);

  // Enable internal pull-up resistor for GPIO 4
  pinMode(ONE_WIRE_BUS, INPUT_PULLUP);

  // Set the hostname
  WiFi.hostname(hostname);

  // Connect to WiFi
  Serial.println();
  Serial.println();
  Serial.print("Connecting to ");
  Serial.println(ssid);

  WiFi.begin(ssid, password);

  while (WiFi.status() != WL_CONNECTED) {
    delay(500);
    Serial.print(".");
  }

  Serial.println("");
  Serial.println("WiFi connected.");
  Serial.println("IP address: ");
  Serial.println(WiFi.localIP());

  server.begin();
}

void loop() {
  WiFiClient client = server.available();

  if (client) {
    Serial.println("New Client.");
    String currentLine = "";

    while (client.connected()) {
      if (client.available()) {
        char c = client.read();
        Serial.write(c);
        if (c == '\n') {
          if (currentLine.length() == 0) {
            client.println("HTTP/1.1 200 OK");
            client.println("Content-type:text/html");
            client.println("Connection: close");
            client.println();

            // Read temperature from DS18B20
            sensors.requestTemperatures();
            float temperatureC = sensors.getTempCByIndex(0);

            // Return temperature as JSON
            client.print("{\"temperature\": ");
            client.print(temperatureC);
            client.println("}");

            break;
          } else {
            currentLine = "";
          }
        } else if (c != '\r') {
          currentLine += c;
        }
      }
    }

    client.stop();
    Serial.println("Client disconnected.");
  }
}
