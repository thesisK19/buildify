import React from "react";
import cx from "classnames";
import { Text } from "../Text";
import { StyledButton } from "./styled";

export default Button = (props) => {
  const {
    text,
    textComponent,
    color,
    styledClassNames,
    fontSize,
    fontWeight,
    textAlign,
    ...otherProps
  } = props;

  const styledClassNamesValues = (
    Object.values(styledClassNames)
  ).flat();
  return (
    <StyledButton
      className={cx([
        "rounded w-full px-4 py-2 mt-4",
        {
          "shadow-lg": props.buttonStyle === "full",
        },
        styledClassNamesValues,
      ])}
      {...otherProps}
    >
      <Text
        {...textComponent}
        text={text}
        color={color}
        fontSize={fontSize}
        fontWeight={fontWeight}
        textAlign={textAlign}
      />
    </StyledButton>
  );
};

