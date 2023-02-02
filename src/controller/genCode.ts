import { Response, Request } from "express";
import { EXPORT_DIR, HttpStatusCode, OVERDUE_TIME } from "../config/constant";
import { readFileSync, writeFileSync } from "fs";
// import { zip } from 'zip-a-folder';
var fs = require("fs-extra");
import { GenCodeService } from "../service";
import { join } from "path";

class GenCodeController {
  //CORE FUNCTION

  async getReactSourceCode(req: Request, res: Response) {
    try {
      // const rootFolderName = new Date().getTime().toString()

      const input = readFileSync(join(__dirname, "example.json"), "utf-8");
      let service = new GenCodeService();
      service.getPage(JSON.parse(input));
    } catch (e: any) {
      console.log(e);
      res.status(HttpStatusCode.OK).json({
        statusCode: HttpStatusCode.INTERNAL_ERROR,
        errorMessage: "Internal Server Error, please try again",
      });
    }
  }

  async getAllFrame(req: Request, res: Response) {
    try {
      const rootFolderName = new Date().getTime().toString();

      // await zip(`${EXPORT_DIR}/${rootFolderName}`, `${EXPORT_DIR}/${rootFolderName}_export/${figmaFolderName}.zip`);

      // const downloadUrl = `http://${HOST_URL}/download-source?time=${rootFolderName}&fileName=${encodeURI(figmaFolderName)}`

      // const fileBaseDir = `${EXPORT_DIR}/${rootFolderName}`;
      // const fileZipDir = `${EXPORT_DIR}/${rootFolderName}_export`;

      // res.status(statusCode).json({
      //     statusCode: HttpStatusCode.OK,
      //     responseData: {
      //         timeCreated: rootFolderName,
      //         fileName: encodeURI(figmaFolderName),
      //         downloadUrl
      //     }
      // })

      // setTimeout(() => {
      //     fs.rmSync(fileBaseDir, { recursive: true, force: true })
      //     fs.rmSync(fileZipDir, { recursive: true, force: true })
      //     console.log("FILE WAS DELETED AFTER 10 MINUTES")
      // }, OVERDUE_TIME)
    } catch (e: any) {
      console.log(e);
      res.status(HttpStatusCode.OK).json({
        statusCode: HttpStatusCode.INTERNAL_ERROR,
        errorMessage: "Internal Server Error, please try again",
      });
    }
  }

  // DOWNLOAD REACT ZIP FOLDER
  async downloadZip(req: Request, res: Response) {
    try {
      const { time = "", fileName = "" } = req.query;
      console.log("filename", decodeURI(<string>fileName));
      const fileDir = `${EXPORT_DIR}/${time}_export/${decodeURI(<string>fileName)}.zip`;
      const fileBaseDir = `${EXPORT_DIR}/${time}`;
      const fileZipDir = `${EXPORT_DIR}/${time}_export`;

      if (!fileName || !fs.existsSync(fileDir)) {
        res.status(HttpStatusCode.OK).json({
          statusCode: HttpStatusCode.NOT_FOUND,
          errorMessage:
            "Sorry, the file you have requested does not exist or was deleted after 10 minutes, please try to generate again!",
        });
      }

      res.download(fileDir, (error: any) => {
        // Remove the two folders after 10 minutes after the file is downloaded
        if (!error) {
          setTimeout(() => {
            fs.rmSync(fileBaseDir, { recursive: true, force: true });
            fs.rmSync(fileZipDir, { recursive: true, force: true });
            console.log("FILE WAS DELETED AFTER 10 MINUTES");
          }, OVERDUE_TIME);
        }
      }); // Set disposition and send it.
    } catch (e) {
      res.status(HttpStatusCode.INTERNAL_ERROR).json({
        statusCode: HttpStatusCode.INTERNAL_ERROR,
        errorMessage: "Internal Server Error, please try again",
      });
    }
  }
}

export default new GenCodeController();
