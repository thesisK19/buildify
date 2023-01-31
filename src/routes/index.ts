import { Application } from "express";
import GenCodeRouter from './genCode'

// SPECIFY HOW MANY ENDPOINTS WE SHOULD HAVE
// METHOD: GET, POST
export default function route(app: Application) {
    app.use('/', GenCodeRouter)
}
