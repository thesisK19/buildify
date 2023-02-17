import cx from "classnames";
import React from "react";

import { StyledInput } from "./styled";

import { InputProps } from "./props";
import { defaultProps } from "./props";

export const Input = (props: any) => {
  const { styledClassNames, type, placeholder, inputOptions, className, ...otherProps } =
    props;
  const styledClassNamesValues = (
    Object.values(styledClassNames) as string[]
  ).flat();
  const { required, readonly } = inputOptions;
  return (
    <StyledInput
      className={cx([
        className,
        "input rounded w-full px-4 py-2 mt-4",
        {
          "shadow-lg": props.inputStyle === "full",
        },
        styledClassNamesValues,
      ])}
      type={type}
      required={required}
      placeholder={placeholder}
      readOnly={readonly}
      {...otherProps}
    />
  );
};

Input.defaultProps = defaultProps;
