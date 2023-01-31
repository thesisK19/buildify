import { Response, Request } from 'express'
import { EXPORT_DIR, HttpStatusCode , OVERDUE_TIME} from '../config/constant'
var fs = require("fs-extra");
import { zip } from 'zip-a-folder';
import { RequestBodyType } from '../config/type';
import {
    GenCodeService,
} from "../service";

class GenCodeController {
    //CORE FUNCTION
    async getAllFrame(req: Request, res: Response) {
        try {
            const rootFolderName = new Date().getTime().toString()
            let {
                figmaAccessToken,
                figmaFileId,
                enableDataIntegration,
                isAirtableHost,
                dataSourceEndpoint,
                dataSourceAuthToken,
            }: RequestBodyType = req.body

            const figmaService: GenCodeService = new GenCodeService(figmaAccessToken, figmaFileId)
            const componentService: ComponentService = new ComponentService(figmaService, airtableService, dataSourceService, enableDataIntegration, isAirtableHost)

            // const svgService: SvgService = new SvgService(figmaToken, fileId)

            /// FILTER DATA FROM FIGMA TO KNOW WHICH ONE IS FRAME
            const { frameData, statusCode, figmaFolderName, figmaComponentsData = []} = await figmaService.getAllFrames()


            /// GET ALL IMAGE ASSET FROM FIGMA
            const figmaImageAsset = await figmaService.getAllImageHashMap()

            /// GETTING FRAME DATA with REACT CODE GEN AND CSS STYLE
            const {imgAsset} = await componentService.convertFrame(frameData, figmaImageAsset.data, rootFolderName, figmaComponentsData)

            // DOWNLOAD ALL RELEVANT SVG JSX CONTENT
            // await getSvgJsxAsset(figmaToken, fileId, svgAsset, timeCreated)

            // DOWNLOAD ALL IMAGE FILE
            const imageService: ImageService = new ImageService(figmaAccessToken, figmaFileId, imgAsset, figmaImageAsset.data)
            await imageService.downloadAllImageAssets(rootFolderName)

            await zip(`${EXPORT_DIR}/${rootFolderName}`, `${EXPORT_DIR}/${rootFolderName}_export/${figmaFolderName}.zip`);

            const HOST_URL = (process.env.FIGACT_ENV === 'PRODUCTION') ? req.headers.host : 'localhost:5050'

            const downloadUrl = `http://${HOST_URL}/download-source?time=${rootFolderName}&fileName=${encodeURI(figmaFolderName)}`

            const fileBaseDir = `${EXPORT_DIR}/${rootFolderName}`;
            const fileZipDir = `${EXPORT_DIR}/${rootFolderName}_export`;

            res.status(statusCode).json({
                statusCode: HttpStatusCode.OK,
                responseData: {
                    timeCreated: rootFolderName,
                    fileName: encodeURI(figmaFolderName),
                    downloadUrl
                }
            })

            setTimeout(() => {
                fs.rmSync(fileBaseDir, { recursive: true, force: true })
                fs.rmSync(fileZipDir, { recursive: true, force: true })
                console.log("FILE WAS DELETED AFTER 10 MINUTES")
            }, OVERDUE_TIME)
        }
        catch (e: any) {
            console.log(e)
            res
                .status(HttpStatusCode.OK)
                .json({
                    statusCode: HttpStatusCode.INTERNAL_ERROR,
                    errorMessage: "Internal Server Error, please try again",
                })
        }
    }

    // DOWNLOAD REACT ZIP FOLDER
    async downloadZip(req: Request, res: Response) {
        try {
            const { time = '', fileName = '' } = req.query
            console.log("filename", decodeURI(<string>fileName))
            const fileDir = `${EXPORT_DIR}/${time}_export/${decodeURI(<string>fileName)}.zip`;
            const fileBaseDir = `${EXPORT_DIR}/${time}`;
            const fileZipDir = `${EXPORT_DIR}/${time}_export`;

            if (!fileName || !fs.existsSync(fileDir)) {
                res.status(HttpStatusCode.OK).json({
                    statusCode: HttpStatusCode.NOT_FOUND,
                    errorMessage: "Sorry, the file you have requested does not exist or was deleted after 10 minutes, please try to generate again!"
                })
            }

            res.download(fileDir, (error: any) => {
                // Remove the two folders after 10 minutes after the file is downloaded
                if (!error) {
                    setTimeout(() => {
                        fs.rmSync(fileBaseDir, { recursive: true, force: true })
                        fs.rmSync(fileZipDir, { recursive: true, force: true })
                        console.log("FILE WAS DELETED AFTER 10 MINUTES")
                    }, OVERDUE_TIME)
                }
            }); // Set disposition and send it.
        }
        catch (e) {
            res
                .status(HttpStatusCode.INTERNAL_ERROR)
                .json({
                    statusCode: HttpStatusCode.INTERNAL_ERROR,
                    errorMessage: "Internal Server Error, please try again",
                })
        }
    }
}

export default new GenCodeController
