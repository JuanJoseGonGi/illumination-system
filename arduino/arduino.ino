#include <Servo.h>

bool isAutomatic = false;
int fotoresistor = A2;
int valueLDR, valuePIRLED;
int ledLDR = 22;
int pinPIR = 2;

// servo
Servo servo;
int MOTORSERVO = 8;
int minPulse = 500;  // 0
int maxPulse = 2500; // 180

// r2r
int pin9 = 9;
int pin10 = 10;
int pin11 = 11;
int pin12 = 12;

unsigned long previousMillis = 0;
const unsigned long interval = 4000; // 10 segundos

bool isEngineOn;
unsigned long engineOnMillis;
const unsigned long engineOnMaxDuration = 300;

void setup()
{
  pinMode(fotoresistor, INPUT);
  pinMode(pinPIR, INPUT);
  pinMode(ledLDR, OUTPUT);
  digitalWrite(ledLDR, HIGH);
  servo.attach(MOTORSERVO, minPulse, minPulse); // servo
  pinMode(pin9, OUTPUT);
  pinMode(pin10, OUTPUT);
  pinMode(pin11, OUTPUT);
  pinMode(pin12, OUTPUT);

  isEngineOn = false;

  Serial.begin(9600);
}

void loop()
{
  delay(2);
  if (isAutomatic)
  {
    ctrlLDR();
    ctrlPIR();
  }

  handleEngineStop();

  readSerialController();
}

void handleEngineStop()
{
  if (!isEngineOn)
  {
    return;
  }

  unsigned long currentMillis = millis();
  if (currentMillis - engineOnMillis >= engineOnMaxDuration)
  {
    stopEngine();
  }
}

void ctrlLDR()
{
  if (valuePIRLED == HIGH)
  {
    // Reiniciar el temporizador si se detecta movimiento
    previousMillis = millis();
    valueLDR = map(analogRead(fotoresistor), 0, 1023, 0, 255);
    // analogWrite(ledLDR, HIGH); // control del motor
    digitalWrite(ledLDR, LOW);
    Serial.println("Bombillo ON");
    // Serial.println(valueLDR);
  }
  else
  {
    unsigned long currentMillis = millis();
    // Verificar si han pasado 10 segundos desde la última detección de movimiento
    if (currentMillis - previousMillis >= interval)
    {
      // analogWrite(ledLDR, 0);
      digitalWrite(ledLDR, HIGH);
      Serial.println("Bombillo OFF");
    }
  }
  // Serial.println(valuePIRLED);
}

void crtlPuenteH()
{
  analogWrite(6, 255);
  analogWrite(7, 255);
  analogWrite(5, 0);
  analogWrite(4, 0);
  delay(5000); // Pausa de 5 segundos
  analogWrite(6, 0);
  analogWrite(7, 0);
}

void ctrlPIR()
{
  valuePIRLED = digitalRead(pinPIR);
  Serial.print("PIR: ");
  Serial.println(valuePIRLED);
}

void turnOnLed()
{
  analogWrite(ledLDR, 255);
}

void turnOffLed()
{
  analogWrite(ledLDR, 0);
}

void runEngine(int pin6Value, int pin7Value, int pin5Value, int pin4Value)
{
  analogWrite(6, pin6Value);
  analogWrite(7, pin7Value);
  analogWrite(5, pin5Value);
  analogWrite(4, pin4Value);
  playAlarm();
  stopAlarm();
  isEngineOn = true;
}

void runEngineForward()
{
  runEngine(255, 255, 0, 0);
}

void runEngineBackward()
{
  runEngine(0, 0, 255, 255);
}

void openCurtain()
{
  runEngineForward();
}

void closeCurtain()
{
  runEngineBackward();
}

void openBlinds()
{
  servo.write(90);
  playAlarm();
  stopAlarm();
}

void cloneBlinds()
{
  servo.write(20);
  playAlarm();
  stopAlarm();
}

void playAlarm()
{
  digitalWrite(pin9, HIGH);
  digitalWrite(pin10, HIGH);
  digitalWrite(pin11, HIGH);
  digitalWrite(pin12, HIGH);
  delay(1000);
}

void stopAlarm()
{
  digitalWrite(pin9, LOW);
  digitalWrite(pin10, LOW);
  digitalWrite(pin11, LOW);
  digitalWrite(pin12, LOW);
}

void readSerialController()
{
  if (Serial.available())
  {
    String command = Serial.readStringUntil('\n');

    if (command.startsWith("lights"))
    {
      handleLights(command);
    }

    else if (command.startsWith("curtains"))
    {
      handleCurtains(command);
    }
  }
}

void handleLights(String command)
{
  String state = command.substring(7); // get the state part
  state.trim();                        // remove trailing newline and spaces

  if (state == "on")
  {
    turnOffLed();
    Serial.println("Lights turned off");
    return;
  }

  if (state == "off")
  {
    turnOnLed();
    Serial.println("Lights turned on");
    return;
  }

  Serial.println("Unknown command for lights");
}

void handleCurtains(String command)
{
  String state = command.substring(9); // get the state part
  state.trim();                        // remove trailing newline and spaces

  if (state == "on")
  {
    openCurtain();
    Serial.println("Curtains opened");
    return;
  }

  if (state == "off")
  {
    closeCurtain();
    Serial.println("Curtains closed");
    return;
  }

  Serial.println("Unknown command for curtains");
}

void handleBlinds(String command)
{
  String state = command.substring(9); // get the state part
  state.trim();                        // remove trailing newline and spaces

  if (state == "on")
  {
    openBlinds();
    Serial.println("Blinds opened");
    return;
  }

  if (state == "off")
  {
    cloneBlinds();
    Serial.println("Blinds closed");
    return;
  }

  Serial.println("Unknown command for curtains");
}

void handleLightsMode(String command)
{
  String state = command.substring(9); // get the state part
  state.trim();                        // remove trailing newline and spaces

  if (state == "on")
  {
    isAutomatic = true;
    Serial.println("Automatic");
    return;
  }

  if (state == "off")
  {
    isAutomatic = false;
    Serial.println("Manual");
    return;
  }

  Serial.println("Unknown command for curtains");
}

void stopEngine()
{
  analogWrite(6, 0);
  analogWrite(7, 0);
  analogWrite(5, 0);
  analogWrite(4, 0);
  isEngineOn = false;
  Serial.println("Engine stopped");
}