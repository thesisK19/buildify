import React, { FC, ReactElement } from "react";
import {
  Button,
  Video,
  Anchor,
  Image,
  Container,
  Custom1,
  Custom3,
  OnlyButtons,
  Text,
  Custom2,
  Custom2VideoDrop,
  Custom3BtnDrop,
} from "src/components";
import { useGetValuesFromReferencedProps } from "../../hooks/useGetValuesFromReferencedProps";
import { props } from "./props";
const Home: FC = (): ReactElement => {
  const pageProps = useGetValuesFromReferencedProps(props);
  return (
    <Container {...pageProps["ROOT"]}>
      Hello
    </Container>
  );
}

export default Home;
