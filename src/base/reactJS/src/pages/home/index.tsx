import React, { FC, ReactElement } from "react";
import { Button, Container } from "src/components";
import { PROPS_BY_ID } from "./props";

const HomePage: FC = (): ReactElement => (
  <Container {...PROPS_BY_ID['Container3_xkhddo']}>
    <h1 className="custom-mt-2">Home page</h1>
    <Button {...PROPS_BY_ID["Button1_xhkddjkd"]} />
  </Container>
);

export default HomePage;
