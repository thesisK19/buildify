import { Route as RouteType } from "./config";
import { Route } from "react-router-dom";

const renderRoutes = (routes: RouteType[]) => {
  return routes.map((route) => {
    const { exact, path, Page } = route;
    return (
      <Route exact={exact} path={path} key={path}>
        <Page />
      </Route>
    );
  });
};

export default renderRoutes;