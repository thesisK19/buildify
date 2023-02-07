import styled from "styled-components";

export const StyledButton = styled.button`
  background: ${(props) => (props.buttonStyle === "full" ? `rgba(${Object.values(props.background)})` : "transparent")};
  border: 2px solid transparent;
  border-color: ${(props) =>
    props.buttonStyle === "outline" ? `rgba(${Object.values(props.background)})` : "transparent"};
  padding: ${({ padding }) => `${padding[0]}px ${padding[1]}px ${padding[2]}px ${padding[3]}px`};
  margin: ${({ margin }) => `${margin[0]}px ${margin[1]}px ${margin[2]}px ${margin[3]}px`};
  width: ${(props) => `${props.width}`};
  height: ${(props) => `${props.height}`};
`;
