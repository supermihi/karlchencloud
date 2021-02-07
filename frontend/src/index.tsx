import React from "react";
import ReactDOM from "react-dom";
import LoginView from "./session/LoginView";

const App = () => (
  <LoginView
    name="Nils"
    loading={false}
    login={() => 1}
    forgetLogin={() => 2}
  />
);

ReactDOM.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>,
  document.getElementById("root")
);
