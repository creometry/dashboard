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

export const getUserData = async (access_token) => {
  try {
    const resp = await axios.get(`https://api.github.com/user`, {
      headers: {
        Authorization: `token ${access_token}`,
      },
    });
    return {
      user: resp.data,
      error: null,
    };
  } catch (err) {
    return {
      user: null,
      error: "get user data err : " + err,
    };
  }
};
