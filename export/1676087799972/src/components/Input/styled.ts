import styled from "styled-components";
import { InputProps } from "./props";

export const StyledInput = styled.input<InputProps>`
  background: ${(props) =>
    props.inputStyle === "full"
      ? `rgba(${Object.values(props.background)})`
      : "transparent"};
  border: 2px solid transparent;
  border-color: ${(props) => `rgba(${Object.values(props.borderColor)})`};
  color: ${(props) => `rgba(${Object.values(props.color)})`};
  font-size: ${(props) => `${props.fontSize}px`};
  text-align: ${(props) => `${props.textAlign}`};
  font-weight: ${(props) => `${props.fontWeight}`};
  padding: ${({ padding }) =>
    `${padding[0]}px ${padding[1]}px ${padding[2]}px ${padding[3]}px`};
  margin: ${({ margin }) =>
    `${margin[0]}px ${margin[1]}px ${margin[2]}px ${margin[3]}px`};
  width: ${(props) => `${props.width}`};
  height: ${(props) => `${props.height}`};
  &:focus {
    border-color: ${(props) => `rgba(${Object.values(props.borderColorFocus)})`};
    outline: none;
  }
`;
