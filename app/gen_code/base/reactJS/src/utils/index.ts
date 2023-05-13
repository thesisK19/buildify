export const getValuesFromReferencedPropsObject = (propsObject, database, theme) => {
  Object.keys(propsObject).forEach((key) => {
    const currentValue = propsObject[key];
    if (typeof currentValue === "object" && currentValue !== null) {
      // is not null object
      if (currentValue.type === "dynamic-data") {
        propsObject[key] = database?.[currentValue.collectionId]?.documents?.[currentValue.documentId]?.[currentValue.key];
      } else if (currentValue.type === "theme") {
        // get theme value
        propsObject[key] = theme[currentValue.name];
      } else {
        getValuesFromReferencedPropsObject(propsObject[key], database, theme);
      }
    }
  });
};
