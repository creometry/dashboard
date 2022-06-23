import axios from "axios";

export const exchange = async (code, url) => {
  try {
    const resp = await axios.get(`${url}/api/v1/github/exchange/${code}`);
    console.log("resp : ", resp.data);
    const access_token = resp.data.access_token;
    return {
      access_token: access_token,
      error: null,
    };
  } catch (err) {
    return {
      access_token: null,
      error: "get acces_token err : " + err,
    };
  }
};
