import { createResource, createSignal } from "solid-js";
import { SensorState } from "../pb/sensorstate";

const fetchSensorState = async () => {
  const res = await fetch("http://localhost:8000/api/SensorState");
  const ss: SensorState = await res.json();
  return ss;
};

export const [sensorState, sensorStateReq] = createResource(fetchSensorState);
