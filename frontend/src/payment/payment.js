import axios from "axios";

export const checkPayment = async (payment_token, auth_token) => {
  // make axios request to check if the payment is valid or not with authorization header
  try {
    const resp = await axios.get(
      `https://sandbox.paymee.tn/api/v1/payments/${payment_token}/check`,
      {
        headers: {
          Authorization: `Token ${auth_token}`,
        },
      }
    );
    return resp.data.message;
  } catch (err) {
    return "error";
  }
};

export const createPayment = async (url, vendor, price, auth_token) => {
  try {
    const resp = await fetch(url, {
      method: "POST",
      body: JSON.stringify({ vendor: vendor, amount: price, note: "test" }),
      headers: {
        "Content-Type": "application/json",
        Authorization: `Token ${auth_token}`,
      },
    });
    // get the response as json
    const data = await resp.json();
    return data;
  } catch (err) {
    return {
      message: "error",
    };
  }
};
