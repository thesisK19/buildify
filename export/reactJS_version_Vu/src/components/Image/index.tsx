import React from "react";
import { defaultProps, ImageProps } from "./props";
import { StyledImage, StyledDiv } from "./styled";
import cx from "classnames";

export const Image = (props: ImageProps) => {
  const { width, height, src, alt, className, ...otherProps } = props;

  return (
    <StyledDiv style={{width, height}}>
      <StyledImage
        className={cx([className, "image"])}
        src={src}
        alt={alt}
        {...otherProps}
      />
    </StyledDiv>
  );
};

Image.defaultProps = defaultProps;
