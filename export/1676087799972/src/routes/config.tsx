import React from "react"
      import { Home,About } from "../pages"

      export type Route = {
        exact: boolean;
        path: string;
        Page: React.ElementType;
      };

      export const ROUTES: Route[] = [
        {
          exact: true,
          path: "/",
          Page: Home,
        },{
          exact: true,
          path: "/about",
          Page: About,
        }
      ]