import React from 'react'
import { useNavigate } from 'react-router-dom'

export const Pro = () => {
    const { REACT_APP_VENDOR, REACT_APP_TOKEN, REACT_APP_SUCCESS_URL, REACT_APP_ERROR_URL } = process.env
    const navigate = useNavigate()
    const payPlan = async () => {
        let amount = 149
        const resp = await fetch(
            "https://sandbox.paymee.tn/api/v1/payments/create",
            {
                method: "POST",
                body: JSON.stringify({ vendor: REACT_APP_VENDOR, amount: amount, note: "test" }),
                headers: {
                    "Content-Type": "application/json",
                    Authorization: `Token ${REACT_APP_TOKEN}`,
                },
            }
        );
        // get the response as json
        const data = await resp.json();

        if (data.message !== "Success") {
            navigate('/paymenterror')
        }

        const token = data.data.token;

        // submit a form with the token as a hidden field
        const form = document.createElement("form");
        form.action = "https://sandbox.paymee.tn/gateway/";
        form.method = "POST";
        form.innerHTML = `<input type="hidden" name="payment_token" value="${token}">`;
        form.innerHTML += `<input type="hidden" name="url_ok" value="${REACT_APP_SUCCESS_URL}">`;
        form.innerHTML += `<input type="hidden" name="url_ko" value="${REACT_APP_ERROR_URL}">`;

        document.body.appendChild(form);
        form.submit();

    }
    return (
        <div className='flex justify-center items-center h-screen'>
            <div className='mr-2 bg-yellow-500 text-gray-100 rounded-md py-2 px-6 text-lg font-bold cursor-pointer hover:bg-yellow-400'
                onClick={() => payPlan()}
            >Get The Pro Plan</div>
        </div>
    )
}
