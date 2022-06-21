import React from 'react'
import { Link } from 'react-router-dom'

export const PaymentError = () => {
    return (
        <div className='flex flex-col justify-center items-center h-screen'>
            <div className='text-2xl text-red-500 font-bold'>Payment Error</div>
            <Link to="/steps">
                <div className='text-lg underline'>
                    Retry
                </div>
            </Link>
        </div>
    )
}
