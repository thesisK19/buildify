import React from "react";
import { BrowserRouter as Router, Switch } from "react-router-dom";
import { ROUTES } from "./config";
import renderRoutes from "./renderRoutes";

const RouterDOM = () => {
  return (
    <Router>
      <Switch>
        {
          renderRoutes(ROUTES)
        }
      </Switch>
    </Router>
  );
};

export default RouterDOM;