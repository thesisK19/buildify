import {RGBA} from "../config/type";
// import {ANCHOR_HORIZOLTAL, ANCHOR_VERTICAL, COMPONENT, LIST_ELEMENT, SPECIAL_COMPONENT} from "../config/constant";
import BaseService from "./BaseService";
import ComponentService from "./ComponentService";

export default class StyleService extends BaseService {
    static convertRGBA = (color: RGBA, opacity?: number): RGBA => {
        const {r, g, b, a} = color
        return {
            r: r * 255,
            g: g * 255,
            b: b * 255,
            a: opacity ?? a
        }
    }
}
