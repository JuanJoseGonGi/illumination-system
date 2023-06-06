String command;

void setup()
{
  Serial.begin(9600);
}

void loop()
{
  if (Serial.available())
  {
    command = Serial.readStringUntil('\n');

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

  if (state == "ON")
  {
    // Turn lights on
    // Add your specific code here
    Serial.println("Lights turned on");
    return;
  }

  if (state == "OFF")
  {
    // Turn lights off
    // Add your specific code here
    Serial.println("Lights turned off");
    return;
  }

  Serial.println("Unknown command for lights");
}

void handleCurtains(String command)
{
  String state = command.substring(9); // get the state part
  state.trim();                        // remove trailing newline and spaces

  if (state == "OPEN")
  {
    // Open curtains
    // Add your specific code here
    Serial.println("Curtains opened");
    return;
  }

  if (state == "CLOSED")
  {
    // Close curtains
    // Add your specific code here
    Serial.println("Curtains closed");
    return;
  }

  Serial.println("Unknown command for curtains");
}
