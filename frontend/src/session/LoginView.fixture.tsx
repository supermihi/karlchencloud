import React from "react";
import LoginView from "./LoginView";

export default (
  <LoginView
    name="Nils"
    loading={false}
    login={() => console.log("login")}
    forgetLogin={() => 2}
  />
);
