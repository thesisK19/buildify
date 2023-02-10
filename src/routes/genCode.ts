import express from "express";
import genCodeController from '../controller/genCode'
const router = express.Router()

// router.post('/gen-code', genCodeController.getReactSourceCode)
router.get('/gen-code', genCodeController.getReactSourceCode)

export default router
