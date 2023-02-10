import React from 'react';

import {Button} from '../Button';
import {Container} from '../Container';

export const OnlyButtons = ({ children, ...props }) => {
  return (
    <div title="only-buttons" className="w-full mt-5" {...props}>
      {children}
    </div>
  );
};

export const Custom1 = (props) => {
  return (
    <Container {...props}>
      <h2 className="text-lg px-10 py-5 text-white">
        I'm a component that only accepts
        <br /> buttons.
      </h2>
      <OnlyButtons>
        <Button />
        <Button
          buttonStyle="outline"
          color={{ r: 255, g: 255, b: 255, a: 1 }}
        />
      </OnlyButtons>
    </Container>
  );
};

