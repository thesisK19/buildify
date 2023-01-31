import UtilsService from "./UtilsService";
import ComponentService from './ComponentService'
import fs from "fs";
import {ROUTE_DIR} from "../config/constant";

export default class RouterService {
    static createRouterDomFile = (frameNameList: string[], rootDir: string) => {
        let contentFile = ""
        let rootName = "RouterDOM"
        let pageDir = "../Components"
        let header = "import React from 'react';\nimport { BrowserRouter as Router, Route, Switch } from 'react-router-dom';"
        let importPage = ""
        let routeComponent = ""
        let funcHeader = `\nconst ${rootName} = () => {\n\treturn (\n\t\t<Router>\n\t\t\t<Switch>`
        let homePageName = "HomePage"
        let homePageDir = "index"
        let funcFooter = "\n\t\t\t</Switch>\n\t\t</Router>\n\t);\n}"
        let footer = `\nexport default ${rootName};`

        if (frameNameList.indexOf(homePageDir) !== -1) {
            homePageDir += "__"
        }

        if (frameNameList.indexOf(homePageName) !== -1){
            homePageName = `__${homePageName}__`
        }
        importPage = `\nimport ${homePageName} from '${pageDir}/${homePageDir}';`
        routeComponent = `\n\t\t\t\t<Route exact path="/"><${homePageName} /></Route>`

        frameNameList.forEach((name) => {
            importPage += `\nimport ${ComponentService.formatComponentName(name)} from '${pageDir}/${name}';`
            routeComponent += RouterService.createRouterLink(name)
        })
        contentFile += `${header}${importPage}${funcHeader}${routeComponent}${funcFooter}${footer}`
        fs.writeFile(UtilsService.getFileName(`${rootDir}/${ROUTE_DIR}/index`, "js"), contentFile, (err: any) => {
            if (err) throw err;
            console.log('File router is created successfully.');
        });
    }

    static createRouterLink = (name: string): string => {
        return `\n\t\t\t\t<Route exact path="/${name.toLowerCase()}"><${ComponentService.formatComponentName(name)} /></Route>`
    }

    static createUrlList = (nameList: string[], idList: string[]): any => {
        let urlList: any = {}
        for (let i = 0; i < nameList.length; i++) {
            urlList[idList[i]] = nameList[i].toLowerCase()
        }
        return urlList
    }
}
