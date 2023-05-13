import React from "react";
import "./index.css";
import "./App.css";
import "./styles/index.scss";
import RouterDOM from "./routes";
import { AppProvider } from "./context";

function App() {
  return (
    <div className="App" id="app-container">
      <AppProvider>
        <RouterDOM />
      </AppProvider>
    </div>
  );
}

export default App;
