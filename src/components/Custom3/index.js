import { Element, useNode } from "libs/core/src";
import React from "react";

import { Button } from "../Button";
import { Container } from "../Container";

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
        *
      </Container>
      {/* ------------------ */}
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
          >
            <Text
              {...{
                fontSize: "14",
                textAlign: "left",
                fontWeight: "400",
                color: { r: 92, g: 90, b: 90, a: 1 },
                margin: [0, 0, 0, 0],
                shadow: 0,
                text: "Govern what goes in and out of your components",
                styledClassNames: {},
              }}
            ></Text>
          </Container>
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
    </Container>
  );
};
