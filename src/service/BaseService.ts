import axios from 'axios'
import { HttpStatusCode } from '../config/constant'

export default class BaseService {
    static get = async (url: string = '', customHeader: object = {}, customConfig: object = {}) => {
        try {
            const data = await axios.get(url, {
                headers: {
                    ...customHeader
                }
                ,
                ...customConfig
            })
            return data
        }
        catch (e: any) {
            if (!e.response) {
                return {
                    status: HttpStatusCode.INTERNAL_ERROR,
                    errorMessage: e.message
                }
            }
            else {
                return e.response.data
            }
        }
    }

    protected get = async (url: string = '', customHeader: object = {}, customConfig: object = {}) => {
        try {
            const data = await axios.get(url, {
                headers: {
                    ...customHeader
                }
                ,
                ...customConfig
            })
            return data
        }
        catch (e: any) {
            if (!e.response) {
                return {
                    status: HttpStatusCode.INTERNAL_ERROR,
                    errorMessage: "Internal Server Error, please try again"
                }
            }
            else {
                return e.response.data
            }
        }
    }
}
