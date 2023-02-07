import { Node, Element } from "../config/type";
import BaseService from "./BaseService";
import { readFileSync, writeFileSync } from "fs";
import { join } from "path";
import UtilsService from "./UtilsService";
import { all } from "bluebird";
import { KEY_CHILDREN, PAGES_DIR, EXPORT_DIR, REACT_JS_BASE_DIR, ROUTES_DIR, INDEX } from "../config/constant";
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
  async getPage(rootName: string, input: object) {
    let { components, map } = this.getAllElement(input);
    let code = this.mergeElement("ROOT", map);

    var rootDir = `${EXPORT_DIR}/${rootName}`;
    await this.setUpDirectory(rootName);

    let importReactText = `import React from 'react'`;
    let importComponentsText = Array<string>();

    for (let component of components) {
      importComponentsText.push(`import ${component} from '../components/${component}'`);
    }

    let content = `export default function HomePage () {
                  return (
                    ${code}
                  )}`;

    let text = `${importReactText}\n${importComponentsText.join("\n")}\n\n${content}`;
    console.log(text)

    UtilsService.createAndWriteFile(
      UtilsService.getFileName(`${rootDir}/${PAGES_DIR}/${INDEX}`, "js"),
      text,
      (err: any) => {
        if (err) throw err;
        console.log("File index js is created successfully.");
      }
    );
  }
  setUpDirectory = async (rootName: string) => {
    try {
      var rootDir = `${EXPORT_DIR}/${rootName}`;
      // var routerDir = `${rootDir}/${ROUTE_DIR}`;
      // var pagesDir = `${rootDir}/${PAGES_DIR}`;
      // var componentDir = `${rootDir}/${COMPONENT_DIR}`
      var zipExportDir = `${rootDir}_export`;
      
      await UtilsService.copyFolder(REACT_JS_BASE_DIR, rootDir);

      // UtilsService.createFolder(routerDir);
      // UtilsService.createFolder(pagesDir);
      // UtilsService.createFolder(componentDir)
      // UtilsService.createFolder(zipExportDir);
    } catch (e) {
      console.log(e);
    }
  };
}
