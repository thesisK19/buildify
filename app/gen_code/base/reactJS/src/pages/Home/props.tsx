export const props = {
  ROOT: {
    fontSize: {
      type: "theme",
      name: "fontSizeNormal",
    },
    textAlign: "center",
    fontWeight: {
      type: "theme",
      name: "fontWeightBold",
    },
    background: { r: 0, g: 0, b: 0, a: 0 },
    color: { r: "0", g: "0", b: "0", a: "1" },
    anchorStyle: "outline",
    text: {
      type: "dynamic-data",
      collectionId: 2,
      documentId: 2,
      key: 'name',
    }, padding: ["10", "10", "10", "10"],
    margin: ["5", "0", "5", "0"],
    textComponent: {
      fontSize: "15",
      textAlign: "center",
      fontWeight: {
        type: "theme",
        name: "fontWeightNormal",
      },
      color: { r: "92", g: "90", b: "90", a: "1" },
      margin: ["0", "0", "0", "0"],
      shadow: 0,
      text: {
        type: "dynamic-data",
        collectionId: 1,
        documentId: 1,
        key: 'pKey2',
      },
      styledClassNames: {},
      tagName: "h3",
    },
    width: "auto",
    height: "auto",
    styledClassNames: {},
    events: { pageNavigate: "", absoluteUrlNavigate: "", href: "", clickType: "href" },
    children: [],
  },
};

/// Expect 
// {
//   "fontSize": "14px",
//   "textAlign": "center",
//   "fontWeight": 600,
//   "background": {
//       "r": 0,
//       "g": 0,
//       "b": 0,
//       "a": 0
//   },
//   "color": {
//       "r": "0",
//       "g": "0",
//       "b": "0",
//       "a": "1"
//   },
//   "anchorStyle": "outline",
//   "text": "Document 2 Button Text",
//   "padding": [
//       "10",
//       "10",
//       "10",
//       "10"
//   ],
//   "margin": [
//       "5",
//       "0",
//       "5",
//       "0"
//   ],
//   "textComponent": {
//       "fontSize": "15",
//       "textAlign": "center",
//       "fontWeight": 400,
//       "color": {
//           "r": "92",
//           "g": "90",
//           "b": "90",
//           "a": "1"
//       },
//       "margin": [
//           "0",
//           "0",
//           "0",
//           "0"
//       ],
//       "shadow": 0,
//       "text": "Key 2",
//       "styledClassNames": {},
//       "tagName": "h3"
//   },
//   "width": "auto",
//   "height": "auto",
//   "styledClassNames": {},
//   "events": {
//       "pageNavigate": "",
//       "absoluteUrlNavigate": "",
//       "href": "",
//       "clickType": "href"
//   },
//   "children": []
// }
