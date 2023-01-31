import express from "express";
import genCodeController from '../controller/genCode'
const router = express.Router()

router.post('/gen-code', genCodeController.getAllFrame)

export default router
