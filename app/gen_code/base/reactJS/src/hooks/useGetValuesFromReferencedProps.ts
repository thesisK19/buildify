import { useContext } from "react";
import { AppContext } from "../context";
import { getValuesFromReferencedPropsObject } from "../utils";

export const useGetValuesFromReferencedProps = (props) => {
  const appContext = useContext(AppContext);
  if (!appContext?.appState) {
    return {};
  }
  const { theme, database } = appContext?.appState;

  const newProps = { ...props };

  getValuesFromReferencedPropsObject(newProps, database, theme);

  return newProps;
};
