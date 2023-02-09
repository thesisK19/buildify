import React from "react";
import { TextProps } from "./props";
import cx from "classnames";
import Typography from '@material-ui/core/Typography';
import { defaultProps } from "./props";

export const Text = ({
  fontSize,
  textAlign,
  fontWeight,
  color,
  shadow,
  text,
  margin,
  styledClassNames,
  className,
  tagName,
}: TextProps) => {
  const styledClassNamesValues = (
    Object.values(styledClassNames) as string[]
  ).flat();
  return (
    <Typography
      variant={tagName as any}
      className={cx([className, "text", styledClassNamesValues])}
      align={textAlign as any}
      style={{
        width: "100%",
        margin: `${margin[0]}px ${margin[1]}px ${margin[2]}px ${margin[3]}px`,
        color: `rgba(${Object.values(color)})`,
        fontSize: `${fontSize}px`,
        textShadow: `0px 0px 2px rgba(0,0,0,${(shadow || 0) / 100})`,
        fontWeight,
      }}
    >
      {text}
    </Typography>
  );
};

Text.defaultProps = defaultProps;