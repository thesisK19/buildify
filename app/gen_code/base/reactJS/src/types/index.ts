export type ThemeProp = {
    type: string;
    name: string;
    value?: any;
}

export type DatabaseProp = {
    type: string;
    name: string;
    collectionId: number;
    documentId: number;
    key: string;
    value?: any;
}

export type WithThemeAndDatabase<T> = {
    [P in keyof T]: T[P] | ThemeProp | DatabaseProp;
}