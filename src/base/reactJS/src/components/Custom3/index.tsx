import React from "react";

import {Button} from "../Button";
import {Container} from "../Container";
import { Custom1 } from "../Custom1";
import {Text} from "../Text";

export const Custom3BtnDrop = ({ children }) => {
  return <div className="w-full h-full">{children}</div>;
};

export const Custom3 = (props) => {
  return (
    <Container {...props} className="overflow-hidden">
      <div className="w-full mb-4">
        <h2 className="text-center text-xs text-white">I must have at least 1 button</h2>
      </div>
      <Custom3BtnDrop>
        <Button background={{ r: 184, g: 247, b: 247, a: 1 }} />
      </Custom3BtnDrop>
    </Container>
  );
};
