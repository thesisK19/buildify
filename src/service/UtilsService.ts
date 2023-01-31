import { EXPORT_DIR } from "../config/constant";
var fs = require("fs-extra");

export default class UtilsService {
    
    static getFileName = (name: string, extention: string): string => {
        return `${EXPORT_DIR}/${name}.${extention}`;
    };
    
    static createFolder = (dir: string) => {
        fs.mkdirSync(dir, { recursive: true });
    }
    
    static createAndWriteFile = (dir: string, content: string, callback: any) => {
        fs.writeFile(dir, content, callback);
    }
    
    static copyFolder = async (baseDir: string, endDir: string) => {
        await fs.copy(baseDir, endDir);
    }
}
