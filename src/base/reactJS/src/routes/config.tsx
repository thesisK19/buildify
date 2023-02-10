import React from "react";
import HomePage from "../pages/home";
import AboutPage from "../pages/home";

export type Route = {
    exact: boolean;
    path: string;
    Page: React.ElementType;
}

export const ROUTES: Route[] = [
    {
        exact: true,
        path: "/",
        Page: HomePage,
    },
]