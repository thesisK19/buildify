export type RGBA = {
  r: number;
  g: number;
  b: number;
  a: number;
};
export type Node = {
  id: string;
  hidden: boolean,
  linkedNodes: Array<string>, //?
  name: string, // ??
  nodes: Array<string>,
  parent: null | string,
  props: object,
  events: any,
};

export type Component = {
  name: string;
};
