import { Node, Element } from "../config/type";
import BaseService from "./BaseService";
import { readFileSync, writeFileSync } from "fs";
import { join } from "path";
import UtilsService from "./UtilsService";
import { all } from "bluebird";
import { KEY_CHILDREN, PAGES_DIR, EXPORT_DIR, REACT_JS_BASE_DIR, ROUTES_DIR, INDEX } from "../config/constant";
export default class GenCodeService extends BaseService {
  // TODO: Classify according to page
  getAllElement(input: object): { components: Array<string>; map: Map<string, Element>; propsByIds: Array<string> } {
    let nodes: Array<Node> = [];
    Object.entries(input).forEach(([key, value]) => {
      nodes.push({
        id: key.replace("Craft", ""),
        type: value?.type?.resolvedName.replace("Craft", ""),
        props: value?.props,
        children: value?.nodes,
        page: value?.page,
        pageName: "Home",
      });
    });
    // console.log(nodes);

    let elements = new Map<string, Element>();
    let components = new Set<string>();
    let propsByIds = new Array<string>();
    for (let node of nodes) {
      let element = this.getElement(node);
      elements.set(node.id, element);
      components.add(element.component);
      propsByIds.push(`"${element.id}": ${element.props}`);
    }
    return {
      components: Array.from(components),
      map: elements,
      propsByIds,
    };
  }
  getElement(node: Node): Element {
    // TODO: page
    let component = node.type;
    let id = node.id;
    let props = JSON.stringify(node.props);
    let element = `<${component} {...PROPS_BY_ID["${id}"]}>${KEY_CHILDREN}</${component}>`;

    return { id, props, component: component, elementString: element, children: node.children };
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

  // index
  async getPage(rootName: string, input: object) {
    let { components, map, propsByIds } = this.getAllElement(input);
    let code = this.mergeElement("ROOT", map);

    await this.setUpDirectory(rootName);

    // create folder
    // create props
    console.log("ddddddddddddd");
    let propsText = `export const PROPS_BY_ID = {${propsByIds.join(",")}}`;
    console.log(propsText);
    UtilsService.createAndWriteFile(
      UtilsService.getFileName(`${rootName}/${PAGES_DIR}/home/props`, "tsx"),
      propsText,
      (err: any) => {
        if (err) throw err;
        console.log("File index js is created successfully.");
      }
    );

    // create page
    let importReactText = `import React, { FC, ReactElement } from "react"`;
    let importComponentsText = `import { ${components.join(",")} } from "src/components";`;
    let importPropsById = `import { PROPS_BY_ID } from "./props";`;
    let exportDefaultPageText = `export default HomePage;`;

    let content = `const HomePage: FC = (): ReactElement => (
                    ${code}
                  )`;

    let text = `${importReactText}\n${importComponentsText}\n${importPropsById}\n\n${content}\n\n${exportDefaultPageText}`;
    console.log(text);

    UtilsService.createAndWriteFile(
      UtilsService.getFileName(`${rootName}/${PAGES_DIR}/home/${INDEX}`, "tsx"),
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
      // var zipExportDir = `${rootDir}_export`;
      await UtilsService.copyFolder(REACT_JS_BASE_DIR, rootDir);
      // TODO: create all page folder

      // UtilsService.createFolder(zipExportDir);
    } catch (e) {
      console.log(e);
    }
  };
}
