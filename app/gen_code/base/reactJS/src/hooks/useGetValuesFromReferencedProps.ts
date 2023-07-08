import { useContext } from "react";
import { WithThemeAndDatabase } from "../types";
import { AppContext } from "../context";
import { getValuesFromReferencedPropsObject } from "../utils";

export const useGetValuesFromReferencedProps = <T>(props: WithThemeAndDatabase<T>) => {
  const appContext = useContext(AppContext);
  if (!appContext?.appState) {
    return {};
  }
  const { theme, database } = appContext?.appState;

  const newProps = { ...props };

  getValuesFromReferencedPropsObject(newProps, database, theme);

  return newProps;
};
