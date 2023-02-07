import React from "react";
import {Container} from "../components/Container";
import {Custom1} from "../components/Custom1";
import {Button} from "../components/Button";

export default function HomePage() {
  return (
    <Container
      {...{
        flexDirection: "column",
        alignItems: "flex-start",
        justifyContent: "flex-start",
        fillSpace: "no",
        padding: ["40", "40", "40", "40"],
        margin: ["0", "0", "0", "0"],
        background: { r: 255, g: 255, b: 255, a: 1 },
        color: { r: 0, g: 0, b: 0, a: 1 },
        shadow: 0,
        radius: 0,
        width: "800px",
        height: "auto",
        styledClassNames: {},
      }}
      
    >
      <Container
        {...{
          flexDirection: "column",
          alignItems: "flex-start",
          justifyContent: "flex-start",
          fillSpace: "no",
          padding: ["40", "40", "40", "40"],
          margin: ["0", "0", "40", "0"],
          background: { r: 234, g: 245, b: 245, a: 1 },
          color: { r: 0, g: 0, b: 0, a: 1 },
          shadow: 0,
          radius: 0,
          width: "100%",
          height: "auto",
          styledClassNames: {},
        }}
      >
        <Button
          {...{
            fontSize: "14",
            textAlign: "center",
            fontWeight: "500",
            background: { r: 255, g: 255, b: 255, a: 0.5 },
            color: { r: 92, g: 90, b: 90, a: 1 },
            buttonStyle: "full",
            text: "Button",
            padding: ["10", "10", "10", "10"],
            margin: ["5", "0", "5", "0"],
            textComponent: {
              fontSize: "15",
              textAlign: "center",
              fontWeight: "500",
              color: { r: 92, g: 90, b: 90, a: 1 },
              margin: [0, 0, 0, 0],
              shadow: 0,
              text: "Text",
              styledClassNames: {},
            },
            width: "100%",
            height: "auto",
            styledClassNames: {},
          }}
        ></Button>
        <Container
          {...{
            flexDirection: "column",
            alignItems: "flex-start",
            justifyContent: "flex-start",
            fillSpace: "no",
            padding: ["0", "0", "0", "0"],
            margin: ["0,", "0", "20", "0"],
            background: { r: 76, g: 78, b: 78, a: 0 },
            color: { r: 0, g: 0, b: 0, a: 1 },
            shadow: 0,
            radius: 0,
            width: "100%",
            height: "auto",
            styledClassNames: {},
          }}
        ></Container>
        <Container
          {...{
            flexDirection: "row",
            alignItems: "flex-start",
            justifyContent: "flex-start",
            fillSpace: "no",
            padding: ["0", "0", "0", "0"],
            margin: ["30", "0", "0", "0"],
            background: { r: 76, g: 78, b: 78, a: 0 },
            color: { r: 0, g: 0, b: 0, a: 1 },
            shadow: 0,
            radius: 0,
            width: "100%",
            height: "auto",
            styledClassNames: {},
          }}
        >
          <Container
            {...{
              flexDirection: "row",
              alignItems: "flex-start",
              justifyContent: "flex-start",
              fillSpace: "no",
              padding: ["0", "20", "0", "0"],
              margin: ["0", "0", "0", "0"],
              background: { r: 0, g: 0, b: 0, a: 0 },
              color: { r: 0, g: 0, b: 0, a: 1 },
              shadow: 0,
              radius: 0,
              width: "45%",
              height: "auto",
              styledClassNames: {},
            }}
          >
            <Custom1
              {...{
                flexDirection: "column",
                alignItems: "flex-start",
                justifyContent: "flex-start",
                fillSpace: "no",
                padding: ["20", "20", "20", "20"],
                margin: ["0", "0", "0", "0"],
                background: { r: 119, g: 219, b: 165, a: 1 },
                color: { r: 0, g: 0, b: 0, a: 1 },
                shadow: 40,
                radius: 0,
                width: "100%",
                height: "auto",
                styledClassNames: {},
              }}
            ></Custom1>
          </Container>
          <Container
            {...{
              flexDirection: "column",
              alignItems: "flex-start",
              justifyContent: "flex-start",
              fillSpace: "no",
              padding: ["0", "0", "0", "20"],
              margin: ["0", "0", "0", "0"],
              background: { r: 0, g: 0, b: 0, a: 0 },
              color: { r: 0, g: 0, b: 0, a: 1 },
              shadow: 0,
              radius: 0,
              width: "55%",
              height: "auto",
              styledClassNames: {},
            }}
          ></Container>
        </Container>
      </Container>
    </Container>
  );
}
