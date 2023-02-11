import React from "react";
import { HomePage, AboutPage } from "../pages";

export type Route = {
  exact: boolean;
  path: string;
  Page: React.ElementType;
};

export const ROUTES: Route[] = [
  {
    exact: true,
    path: "/",
    Page: HomePage,
  },
  {
    exact: true,
    path: "/about",
    Page: AboutPage,
  },
];
