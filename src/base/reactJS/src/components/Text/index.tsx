import { useNode, useEditor } from "libs/core/src";
import React from "react";
import ContentEditable from "react-contenteditable";
import { craftConfig } from "./craft";
import { TextProps } from "./props";
import cx from "classnames";

export const Text = ({
  fontSize,
  textAlign,
  fontWeight,
  color,
  shadow,
  text,
  margin,
  styledClassNames,
}: Partial<TextProps>) => {
  const {
    connectors: { connect },
    setProp,
  } = useNode();
  const { enabled } = useEditor((state) => ({
    enabled: state.options.enabled,
  }));
  const styledClassNamesValues = (
    Object.values(styledClassNames) as string[]
  ).flat();
  return (
    <ContentEditable
      innerRef={connect}
      html={text} // innerHTML of the editable div
      disabled={!enabled}
      onChange={(e) => {
        setProp((prop) => (prop.text = e.target.value), 500);
      }} // use true to disable editing
      tagName="h2" // Use a custom HTML tag (uses a div by default)
      className={cx(styledClassNamesValues)}
      style={{
        width: "100%",
        margin: `${margin[0]}px ${margin[1]}px ${margin[2]}px ${margin[3]}px`,
        color: `rgba(${Object.values(color)})`,
        fontSize: `${fontSize}px`,
        textShadow: `0px 0px 2px rgba(0,0,0,${(shadow || 0) / 100})`,
        fontWeight,
        textAlign,
      }}
    />
  );
};

Text.craft = craftConfig;
