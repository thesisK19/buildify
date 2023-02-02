import { Node, Element } from "../config/type";
import { KEY_CHILDREN } from "../config/constant";
import BaseService from "./BaseService";
import { readFileSync, writeFileSync } from "fs";
import { join } from "path";
import { all } from "bluebird";
export default class GenCodeService extends BaseService {
  // TODO: Classify according to page
  getAllElement(input: object): { components: Set<string>; map: Map<string, Element> } {
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
    let components = new Set<string>();
    for (let node of nodes) {
      let element = this.getElement(node);
      elements.set(node.id, element);
      components.add(element.component);
    }
    return {
      components,
      map: elements,
    };
  }
  getElement(node: Node): Element {
    // TODO: page
    let component = node.type;
    let props = JSON.stringify(node.props);
    let element = `<${component} {...${props}}>${KEY_CHILDREN}</${component}>`;

    return { component: component, elementString: element, children: node.children };
  }
  mergeElement(id: string, map: Map<string, Element>): string {
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
    let allElement = this.getAllElement(input);
    let code = this.mergeElement("ROOT", allElement.map);
    // console.log(code)
    writeFileSync(join(__dirname, "export-test.txt"), JSON.stringify([...allElement.components]) + "\n\n" + code, {
      flag: "w",
    });
  }
}
