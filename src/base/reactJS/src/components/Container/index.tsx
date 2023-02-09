import React from "react";
import cx from "classnames";

import { ContainerProps } from "./props";
import { defaultProps } from "./props";


export const Container = (props: Partial<ContainerProps>) => {
  props = {
    ...defaultProps,
    ...props,
  };
  const {
    flexDirection,
    alignItems,
    justifyContent,
    fillSpace,
    background,
    color,
    padding,
    margin,
    shadow,
    radius,
    children,
    styledClassNames,
    width,
    height,
    className,
  } = props;
  const styledClassNamesValues = (
    Object.values(styledClassNames) as string[]
  ).flat();
  return (
    <div
      className={cx([className, "flex container", styledClassNamesValues])}
      style={{
        justifyContent,
        // @ts-ignore
        flexDirection,
        width,
        height,
        alignItems,
        background: `rgba(${Object.values(background)})`,
        color: `rgba(${Object.values(color)})`,
        padding: `${padding[0]}px ${padding[1]}px ${padding[2]}px ${padding[3]}px`,
        margin: `${margin[0]}px ${margin[1]}px ${margin[2]}px ${margin[3]}px`,
        boxShadow:
          shadow === 0
            ? "none"
            : `0px 3px 100px ${shadow}px rgba(0, 0, 0, 0.13)`,
        borderRadius: `${radius}px`,
        flex: fillSpace === "yes" ? 1 : "unset",
      }}
    >
      {children}
    </div>
  );
};

Container.defaultProps = defaultProps;
