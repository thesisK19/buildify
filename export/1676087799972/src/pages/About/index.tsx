import React, { FC, ReactElement } from "react";
import { Container, Text, Image, Button, Video, Input } from "src/components";
import { PROPS_BY_ID } from "./props";

const About: FC = (): ReactElement => (
  <Container {...PROPS_BY_ID["ROOT_about"]}>
    <Text {...PROPS_BY_ID["Text_SS0mt8wxyk"]} />
    <Button {...PROPS_BY_ID["Button_l44yKBgA_4"]} />
    <Button {...PROPS_BY_ID["Button_HMwJyOTH_o"]} />
    <Button {...PROPS_BY_ID["Button_z_N1TpNJH8"]} />
    <Image {...PROPS_BY_ID["Image_x-HJcZ61k6"]} />
    <Video {...PROPS_BY_ID["Video_s75Y21HcQK"]} />
    <Input {...PROPS_BY_ID["Input_b7E1UXnuvX"]} />
    <Container {...PROPS_BY_ID["Container_dmKjodcYHo"]} />
  </Container>
);

export default About;
