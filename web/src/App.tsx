import { Component, createEffect, For, Show } from "solid-js";
import { sensorState, sensorStateReq } from "./store";

const App: Component = () => {
  createEffect(() => console.log(sensorState()));
  return (
    <div>
      <button onClick={() => sensorStateReq.refetch()}>
        Fetch Sensor State
      </button>
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
