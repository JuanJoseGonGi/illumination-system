const STATES = {
  on: "on",
  off: "off",
}

const API_DOMAIN = "http://localhost:3001";

/**
 * Realiza una petición a la API.
 *
 * @param {string} path
 * @param {string} method
 * @param {Record<string, any> | null} body
 * @returns
 */
async function apiRequest(path, method = "GET", body = null) {
  const options = {
    method,
    headers: {
      "Content-Type": "application/json",
    },
  }

  if (body) {
    options.body = JSON.stringify(body);
  }

  const response = await fetch(`${API_DOMAIN}${path}`, options);
  if (response.status >= 400) {
    if (!response.body) {
      throw new Error(response.statusText);
    }

    const respText = await response.text();
    throw new Error(respText);
  }

  if (!response.body) {
    return;
  }

  if (response.headers.get("Content-Type") === "application/json") {
    return await response.json();
  }
}

async function getSensorsState() {
  const lightsState = await apiRequest("/lights");
  const curtainsState = await apiRequest("/curtains");

  return {
    lights: lightsState,
    curtains: curtainsState,
  }
}

async function changeLightModeState(state) {
  if (!Object.prototype.hasOwnProperty.call(STATES, state)) {
    throw new Error("Estado inválido.");
  }

  const body = {
    state,
  }

  await apiRequest("/lights/mode", "PUT", body);
}

function lightModeControlHandler(event) {
  const lightControlLabel = document.querySelector(".toggle[for='lightControl']");

  const lightModeControl = event.target;
  if (lightModeControl.checked) {
    changeLightModeState(STATES.on).catch(handleError);
    if (lightControlLabel) {
      // @ts-ignore
      lightControlLabel.style.display = "none";
    }
    return;
  }

  changeLightModeState(STATES.off).catch(handleError);

  if (lightControlLabel) {
    // @ts-ignore
    lightControlLabel.style.display = "inline-flex";
  }
}

async function changeLightState(state) {
  if (!Object.prototype.hasOwnProperty.call(STATES, state)) {
    throw new Error("Estado inválido.");
  }

  const body = {
    state,
  }

  await apiRequest("/lights", "PUT", body);
}

function lightControlHandler(event) {
  const lightControl = event.target;
  if (lightControl.checked) {
    changeLightState(STATES.on).catch(handleError);
    return;
  }

  changeLightState(STATES.off).catch(handleError);
}

async function changeCurtainState(state) {
  console.log(state);
  if (!Object.prototype.hasOwnProperty.call(STATES, state)) {
    throw new Error("Estado inválido.");
  }

  const body = {
    state,
  }

  await apiRequest("/curtains", "PUT", body);
}

function curtainControlHandler(event) {
  const curtainControl = event.target;
  if (curtainControl.checked) {
    changeCurtainState(STATES.on).catch(handleError);
    return;
  }

  changeCurtainState(STATES.off).catch(handleError);
}

async function changeBlindState(state) {
  if (!Object.prototype.hasOwnProperty.call(STATES, state)) {
    throw new Error("Estado inválido.");
  }

  const body = {
    state,
  }

  await apiRequest("/blinds", "PUT", body);
}

function blindControlHandler(event) {
  const blindControl = event.target;
  if (blindControl.checked) {
    changeBlindState(STATES.on).catch(handleError);
    return;
  }

  changeBlindState(STATES.off).catch(handleError);
}

function handleError(error) {
  alert("Ocurrió un error: " + error.message);
}

function main() {
  const lightModeControl = document.getElementById("lightModeControl");
  if (!lightModeControl) {
    handleError(new Error("No se encontró el control de modo de luces."));
    return;
  }

  const lightControl = document.getElementById("lightControl");
  if (!lightControl) {
    handleError(new Error("No se encontró el control de luces."));
    return;
  }

  const curtainControl = document.getElementById("curtainControl");
  if (!curtainControl) {
    handleError(new Error("No se encontró el control de cortinas."));
    return;
  }

  const blindControl = document.getElementById("blindControl");
  if (!blindControl) {
    handleError(new Error("No se encontró el control de persianas."));
    return;
  }

  lightModeControl.addEventListener("change", lightModeControlHandler);
  lightControl.addEventListener("change", lightControlHandler);
  curtainControl.addEventListener("change", curtainControlHandler);
  blindControl.addEventListener("change", blindControlHandler);
}

document.addEventListener("DOMContentLoaded", main);