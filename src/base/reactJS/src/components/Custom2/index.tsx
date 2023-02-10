import React from 'react';

import {Container} from '../Container';
import {Video} from '../Video';

export const Custom2VideoDrop = ({ children }) => {
  return (
    <div className="flex-1 ml-5 h-full">
      {children}
    </div>
  );
};

export const Custom2 = (props) => {
  return (
    <Container {...props} className="overflow-hidden">
      <div className="w-24">
        <h2 className="text-xs text-white">
          You can only drop
          <br />
          one video here.
        </h2>
      </div>
      <Custom2VideoDrop>
        <Video />
      </Custom2VideoDrop>
    </Container>
  );
};
