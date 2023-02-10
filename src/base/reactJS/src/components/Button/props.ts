import { Text } from "../Text";

export type ButtonProps = Partial<{
  fontSize: string;
  textAlign: string;
  fontWeight: string;
  width: string;
  height: string;
  background: Record<"r" | "g" | "b" | "a", number>;
  color: Record<"r" | "g" | "b" | "a", number>;
  buttonStyle: string;
  padding: any[];
  margin: any[];
  text: string;
  textComponent: any;
  styledClassNames: Record<string, any>;
  className: string;
}>;

export const defaultProps: ButtonProps = {
  fontSize: "14",
  textAlign: "center",
  fontWeight: "500",
  background: { r: 255, g: 255, b: 255, a: 0.5 },
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
};
