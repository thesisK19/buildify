import * as http from "http";
import express, { Express, Request, Response, NextFunction } from "express";
import cors, { CorsOptions } from 'cors'
import {
    port,
    allowedOrigins
} from './config/config'

export class Server {
    public readonly _app: Express;
    private readonly corsOption: CorsOptions;
    private _server!: http.Server;

    get app(): Express {
        return this._app;
    }

    get server(): http.Server {
        return this._server;
    }

    constructor() {
        this._app = express();
        this.corsOption = {
            origin: allowedOrigins,
            credentials: true,
            methods: [
                'GET',
                'POST',
            ],

            allowedHeaders: [
                'Content-Type',
            ],
        }

        this._app.set("port", port);

        this.configureMiddleware();
    }

    public configureMiddleware() {
        // Required for POST requests
        this._app.use(express.json());
        this._app.use(express.urlencoded({ extended: true }));
        this._app.use(cors(this.corsOption))

        // CORS
        this._app.use(function (req: Request, res: Response, next: NextFunction) {
            res.setHeader("Access-Control-Allow-Origin", "*");
            res.setHeader("Access-Control-Allow-Credentials", "true");
            res.setHeader("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT");
            res.setHeader("Access-Control-Allow-Headers", "Access-Control-Allow-Origin,Access-Control-Allow-Headers, Origin,Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers,Authorization");
            next();
        });
    }

    public start() {
        this._server = this._app.listen(this._app.get("port"), () => {
            console.log("🚀 Server is running on port " + this._app.get("port"));
        });
    }
}
