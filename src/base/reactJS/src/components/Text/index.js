import React from "react";
import ContentEditable from "react-contenteditable";
import cx from "classnames";

export const Text = (props) => {
  const { fontSize, textAlign, fontWeight, color, shadow, text, margin, styledClassNames } = props;
  // const styledClassNamesValues = Object.values(styledClassNames).flat();
  return (
    <ContentEditable
      html={text} // innerHTML of the editable div
      tagName="h2" // Use a custom HTML tag (uses a div by default)
      // className={cx(styledClassNamesValues)}
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
