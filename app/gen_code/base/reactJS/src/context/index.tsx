import React, { createContext, useState } from "react";
import { Database, MyDatabase } from "../database";
import MyTheme, { Theme } from "../theme";

type AppContextProps = {
  children: React.ReactNode;
};

type AppContextType = {
  appState: AppState;
  setAppState: React.Dispatch<React.SetStateAction<AppState>>;
};

type AppState = {
  theme: Theme;
  database: Database;
};

export const AppContext = createContext<AppContextType | null>(null);

export const AppProvider: React.FC<AppContextProps> = ({ children }) => {
  const [appState, setAppState] = useState<AppState>({
    theme: MyTheme,
    database: MyDatabase,
  });
  return <AppContext.Provider value={{ appState, setAppState }}> {children} </AppContext.Provider>;
};
