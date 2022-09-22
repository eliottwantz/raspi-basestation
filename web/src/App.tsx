import { Component, createEffect, createSignal, For, Show } from "solid-js";
import {
  fetchSensorData,
  sensorData,
  sensorState,
  sensorStateReq,
} from "./store";

const App: Component = () => {
  const [sensorId, setSensorId] = createSignal<string>("");
  createEffect(() => console.log(sensorState()));
  createEffect(() => console.log(sensorData()));
  return (
    <div>
      <div>
        <button onClick={() => sensorStateReq.refetch()}>
          Fetch Sensor State
        </button>
        <form>
          <input
            type="text"
            value={sensorId()}
            onInput={(e) => setSensorId(e.currentTarget.value)}
          />
          <button
            onClick={(e) => {
              e.preventDefault();
              fetchSensorData(sensorId());
              setSensorId("");
            }}
          >
            Fetch Sensor Data
          </button>
        </form>
      </div>
      <Show when={sensorState()}>
        <For each={Object.keys(sensorState()!)}>
          {(bm, i) => (
            <div>
              <div>{bm}</div>
              <div>{JSON.stringify(Object.values(sensorState()!)[i()])}</div>
            </div>
          )}
        </For>
      </Show>
    </div>
  );
};

export default App;
