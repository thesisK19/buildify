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
  component: string;
  elementString: string;
  children: Array<string>;
};

