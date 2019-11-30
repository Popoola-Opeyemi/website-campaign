export const getLocalStorage = key => {
  const json = window.localStorage[key];
  if (json == "undefined" || json == undefined) {
    return false;
  }
  return json ? JSON.parse(json) : false;
};

export const saveToLocalStorage = (key, value) => {
  window.localStorage[key] = JSON.stringify(value);
};

export const removeFromLocalStorage = key => {
  delete window.localStorage[key];
};

export const getSession = key => {
  const json = window.sessionStorage[key];
  if (json == "undefined" || json == undefined) {
    return false;
  }
  try {
    return json ? JSON.parse(json) : false;
  } catch (error) {
    return "";
  }
};

export const saveSession = (key, value) => {
  window.sessionStorage[key] = JSON.stringify(value);
};

export const removesession = key => {
  delete window.sessionStorage[key];
};
