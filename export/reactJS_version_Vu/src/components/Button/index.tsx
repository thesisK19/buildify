import cx from "classnames";
import React from "react";
import { Text } from "../Text";
import { ButtonProps, defaultProps } from "./props";

import { StyledButton } from "./styled";

export const Button = React.forwardRef((props: ButtonProps, ref: any) => {
  const {
    text,
    textComponent,
    color,
    styledClassNames,
    fontSize,
    fontWeight,
    textAlign,
    className,
    ...otherProps
  } = props;
  const styledClassNamesValues = (
    Object.values(styledClassNames) as string[]
  ).flat();
  return (
    <StyledButton
      ref={ref}
      className={cx([
        className,
        "button rounded w-full px-4 py-2 mt-4",
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
});

Button.defaultProps = defaultProps;
