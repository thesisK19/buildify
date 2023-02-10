import { Text } from "src/components";

export const PROPS_BY_ID = {
  Button1_xhkddjkd: {
    fontSize: "14",
    textAlign: "center",
    fontWeight: "500",
    background: { r: 92, g: 255, b: 255, a: 0.5 },
    color: { r: 92, g: 90, b: 90, a: 1 },
    buttonStyle: "full",
    text: "Button",
    padding: ["10", "10", "10", "10"],
    margin: ["5", "0", "5", "0"],
    textComponent: {
      ...Text.defaultProps,
      textAlign: "center",
    },
    width: "100%",
    height: "auto",
    styledClassNames: {},
  },
  Container3_xkhddo: {
    flexDirection: "column",
    alignItems: "flex-start",
    justifyContent: "flex-start",
    fillSpace: "no",
    padding: ["0", "0", "0", "0"],
    margin: ["0", "0", "0", "0"],
    background: { r: 0, g: 255, b: 0, a: 0.5 },
    color: { r: 0, g: 0, b: 0, a: 1 },
    shadow: 0,
    radius: 0,
    width: "100%",
    height: "auto",
    styledClassNames: {},
  }
};
