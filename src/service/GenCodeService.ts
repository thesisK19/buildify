import { Node, Element, SerializedData, PageData, PageFullInfo } from "../config/type";
import BaseService from "./BaseService";
import { readFileSync, writeFileSync } from "fs";
import { join } from "path";
import UtilsService from "./UtilsService";
import { all } from "bluebird";
import { KEY_CHILDREN, PAGES_DIR, EXPORT_DIR, REACT_JS_BASE_DIR, ROUTES_DIR, INDEX } from "../config/constant";
export default class GenCodeService extends BaseService {
  convertData(nodesData: object, pagesData: PageData[]): Map<string, PageFullInfo> {
    // map pagePath => pageName
    let mapPagesData = new Map<string, string>();
    pagesData.forEach((page) => {
      mapPagesData.set(page.path, page.name);
    });

    // object map: page => node
    let map = new Map<string, PageFullInfo>();

    Object.entries(nodesData).forEach(([key, value]) => {
      let pagePath: string = value?.page;
      if (!map.has(pagePath)) {
        let pageName = mapPagesData.get(pagePath);
        map.set(pagePath, {
          rootId: "",
          path: pagePath,
          name: pageName ? pageName : "Unkown",
          nodes: new Array<Node>(),
        });
      }
      let id = key.replace("Craft", "");
      let page = map.get(pagePath);
      if (page == undefined) return;

      if (id.startsWith("ROOT")) {
        page.rootId = id;
      }
      page.nodes.push({
        id,
        type: value?.type?.resolvedName.replace("Craft", ""),
        props: value?.props,
        children: value?.nodes,
        page: pagePath,
      });
    });
    return map;
  }

  getAllElementOfPage(nodes: Array<Node>): {
    components: Array<string>;
    map: Map<string, Element>;
    propsByIds: Array<string>;
  } {
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
      element.elementString = `<${element.component} {...PROPS_BY_ID["${id}"]}/>`;
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
  async genPages(rootFolderName: string, data: SerializedData) {
    let { nodes, pages } = data;
    let mapPageToNodes = this.convertData(nodes, pages);

    // copy base reactJS folder + create page folders
    await this.setUpDirectory(rootFolderName, pages);

    // gen all pages
    for (const [pagePath, data] of mapPageToNodes) {
      await this.getPage(rootFolderName, data.name, data.rootId, data.nodes);
    }
    // index page
    this.getIndexPage(rootFolderName, pages);

    // routing
    await this.getConfigRoutes(rootFolderName, pages);
  }

  getIndexPage(rootFolderName: string, pages: PageData[]) {
    let importPages = Array<string>();
    let exportPages = Array<string>();
    for (let page of pages) {
      importPages.push(`import ${page.name} from "./${page.name}";`);
      exportPages.push(page.name);
    }
    let content = `${importPages.join("\n")}\n\nexport { ${exportPages.join(",")} }`;
    UtilsService.createAndWriteFile(
      UtilsService.getFileName(`${rootFolderName}/${PAGES_DIR}/${INDEX}`, "ts"),
      content,
      (err: any) => {
        if (err) throw err;
        console.log("File index js is created successfully.");
      }
    );
  }

  async getConfigRoutes(rootFolderName: string, pages: PageData[]) {
    let importReactText = `import React from "react"`;
    let importPages = Array<string>();
    let routes = Array<string>();
    for (let page of pages) {
      importPages.push(page.name);
      routes.push(
        `{
          exact: true,
          path: "${page.path}",
          Page: ${page.name},
        }`
      );
    }
    let content = `${importReactText}
      import { ${importPages.join(",")} } from "../pages"

      export type Route = {
        exact: boolean;
        path: string;
        Page: React.ElementType;
      };

      export const ROUTES: Route[] = [
        ${routes.join(",")}
      ]`;
    UtilsService.createAndWriteFile(
      UtilsService.getFileName(`${rootFolderName}/${ROUTES_DIR}/config`, "tsx"),
      content,
      (err: any) => {
        if (err) throw err;
        console.log("File index js is created successfully.");
      }
    );
  }

  async getPage(rootFolderName: string, pageName: string, ROOTid: string, nodesData: Array<Node>) {
    let { components, map, propsByIds } = this.getAllElementOfPage(nodesData);
    let code = this.mergeElement(ROOTid, map);

    // create props
    let propsText = `export const PROPS_BY_ID = {${propsByIds.join(",")}}`;
    // console.log(propsText);
    UtilsService.createAndWriteFile(
      UtilsService.getFileName(`${rootFolderName}/${PAGES_DIR}/${pageName}/props`, "tsx"),
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
    let exportDefaultPageText = `export default ${pageName};`;

    let content = `const ${pageName}: FC = (): ReactElement => (
                    ${code}
                  )`;

    let text = `${importReactText}\n${importComponentsText}\n${importPropsById}\n\n${content}\n\n${exportDefaultPageText}`;
    // console.log(text);

    UtilsService.createAndWriteFile(
      UtilsService.getFileName(`${rootFolderName}/${PAGES_DIR}/${pageName}/${INDEX}`, "tsx"),
      text,
      (err: any) => {
        if (err) throw err;
        console.log("File index js is created successfully.");
      }
    );
  }
  setUpDirectory = async (rootFolderName: string, pages: PageData[]) => {
    try {
      var rootDir = `${EXPORT_DIR}/${rootFolderName}`;
      // var zipExportDir = `${rootDir}_export`;
      await UtilsService.copyFolder(REACT_JS_BASE_DIR, rootDir);
      // create all page folder
      for (let page of pages) {
        UtilsService.createFolder(`${rootDir}/${PAGES_DIR}/${page.name}`);
      }

      // UtilsService.createFolder(zipExportDir);
    } catch (e) {
      console.log(e);
    }
  };
}
