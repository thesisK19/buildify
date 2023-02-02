import { Node, Element } from "../config/type";
import { KEY_CHILDREN } from "../config/constant";
import BaseService from "./BaseService";
import { readFileSync, writeFileSync } from "fs";
import { join } from "path";
export default class GenCodeService extends BaseService {
  // TODO: Classify according to page
  getAllElement(input: object): Map<string, Element> {
    let nodes: Array<Node> = [];
    Object.entries(input).forEach(([key, value]) => {
      nodes.push({
        id: key,
        type: value?.type?.resolvedName,
        props: value?.props,
        children: value?.nodes,
        page: value?.page,
      });
    });
    // console.log(nodes);

    let elements = new Map<string, Element>();
    for (let node of nodes) {
      let element = this.getElement(node);
      elements.set(node.id, element);
    }
    return elements;
  }
  getElement(node: Node): Element {
    // TODO: page
    let component = node.type;
    let props = JSON.stringify(node.props);
    let element = `<${component} {...${props}}>${KEY_CHILDREN}</${component}>`;

    return { components: [component], elementString: element, children: node.children };
  }
  mergeElement(id: string, map: Map<string, Element>): string {
    // TODO: Component
    let element = map.get(id);
    if (element == undefined) return "";
    if (element.children.length == 0) {
      element.elementString = element.elementString.replace(KEY_CHILDREN, "");
      return element.elementString;
    }
    let children = "";
    for (let id of element.children) {
      children += this.mergeElement(id, map);
    }
    element.elementString = element.elementString.replace(KEY_CHILDREN, children);
    return element.elementString;
  }
  getPage(input: object) {
    let map = this.getAllElement(input);
    let code = this.mergeElement("ROOT", map);
    // console.log(code)
    writeFileSync(join(__dirname, "export-test.txt"), code, {
      flag: "w",
    });
  }
}
