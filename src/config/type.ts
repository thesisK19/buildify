export type RGBA = {
  r: number;
  g: number;
  b: number;
  a: number;
};
export type Node = {
  id: string;
  type: string;
  props: object;
  children: Array<string>;
  page: string;
  // parent: null | string,
  // events: any,
};

export type Element = {
  id: string;
  props: string;
  component: string;
  elementString: string;
  children: Array<string>;
};

export type SerializedData = {
  nodes: object;
  pages: PageData[];
};

export type PageData = {
  path: string;
  name: string;
};

export type PageFullInfo = {
  rootId: string;
  path: string;
  name: string;
  nodes: Array<Node>;
};
