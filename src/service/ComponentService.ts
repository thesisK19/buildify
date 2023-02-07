import {Promise as BPromise} from "bluebird";
import {
    COMPONENT_DIR,
    PAGES_DIR,
    EXPORT_DIR,
    REACT_JS_BASE_DIR,
    ROUTES_DIR,
} from "../config/constant";

import UtilsService from "./UtilsService";
import StyleService from "./StyleService";
import RouterService from "./RouterService";
import BaseService from "./BaseService";
var fs = require("fs-extra");

export default class ComponentService extends BaseService{

    static wrapText = (parent: string[], child: string): string => {
        child = child.split("\n").join("\n\t");
        return parent[0] + child + parent[1];
    };

    private formatFont = (font: string): string => {
        return font.replace(/[\s]/gi, "+");
    }

    static formatName = (name: string): string => {
        name = name.trim()
        name = name.replace(/[^\w^\|]/gi, "");
        if (name[0] >= "0" && name[0] <= "9") name = "_" + name;
        return name.length > 0 ? name : "_"
    };

    static formatComponentName = (name: string): string => {
        let newName = ComponentService.formatName(name)
        return newName[0].toUpperCase() + newName.slice(1)
    };

    private formatNameSaveClassNameList = (cptName: string, classNameList: any): string => {
        let name = ComponentService.formatName(cptName)
            if(classNameList.hasOwnProperty(name)){
                classNameList[name] += 1
                name += `_${classNameList[name]}`
            }
            else{
                classNameList[name] = 0
            }
        return name
    }



    private getFontUrl = (fontList: string[]) => {
        let fontText = ""
        fontList.forEach((font) => {
            fontText += `${this.formatFont(font)}:ital,wght@0,100;0,300;0,400;0,500;0,700;0,900;1,100;1,300;1,400;1,500;1,700;1,900|`
        })
        fontText = fontText.substr(0, fontText.length - 1)
        return `\t\t<link href="https://fonts.googleapis.com/css?family=${fontText}" rel="stylesheet">`
    }

    private createHTMLFile = async (rootDir: string, fontList: string[]) => {
        let header = ""
        let footer = ""
        await fs.readFile('src/base/html/header.txt', 'utf8' , (err: any, data: string) => {
            if (err) throw err;
            header = data
            fs.readFile('src/base/html/footer.txt', 'utf8' , (err: any, data: string) => {
                if (err) throw err;
                footer = data
                let fontUrl = this.getFontUrl(fontList)
                let result = `${header}\n${fontUrl}\n${footer}`
                UtilsService.createAndWriteFile(
                    UtilsService.getFileName(`${rootDir}/public/index`, "html"),
                    `${result}`,
                    (err: any) => {
                        if (err) throw err;
                        console.log("File html is created successfully.");
                    }
                )
                // fs.writeFile(
                //     getFileName(`${rootDir}/public/index`, "html"),
                //     `${result}`,
                //     (err: any) => {
                //         if (err) throw err;
                //         console.log("File html is created successfully.");
                //     }
                // );
            })
        })
    }

    private createIndexComponent = (frameNameList: string[], rootDir: string) => {
        let indexName = "index"
        if (frameNameList.indexOf(indexName) !== -1) {
            indexName += "__"
        }

        let homePageName = "HomePage"
        if (frameNameList.indexOf(homePageName) !== -1) {
            homePageName = `__${homePageName}__`
        }

        var importReactText = "import React from 'react'";
        var importRouterDom = `\nimport {Link} from 'react-router-dom'`;
        var defaultFuncText = [
            `\nexport default function ${homePageName} () {\n    return (`,
            "\n\t)\n}",
        ];
        var mainClassText = [
            "\n<div>",
            "\n</div>",
        ];
        var content = ""
        frameNameList.forEach((frameName) => {
            content += `\n<Link to='/${frameName}'><div>${frameName}</div></Link>`;
        })
        content = ComponentService.wrapText(mainClassText, content)
        content = ComponentService.wrapText(defaultFuncText, content)
        let text = `${importReactText}${importRouterDom}${content}`

        UtilsService.createAndWriteFile(
            UtilsService.getFileName(`${rootDir}/${COMPONENT_DIR}/${indexName}`, "js"),
            text,
            (err: any) => {
                if (err) throw err;
                console.log("File index js is created successfully.");
            }
        );
    }
}
