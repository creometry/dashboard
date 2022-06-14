import React from 'react'
import { Link } from 'react-router-dom'

export const Plan = () => {
    return (
        <div className='flex flex-col justify-center items-center h-screen'>
            <div className='text-2xl font-bold text-black'>
                Choose a plan
            </div>
            <div className='flex mt-6'>
                <Link to='/plans/starter'>
                    <div className='mr-2  rounded-md py-2 px-6 text-lg font-bold cursor-pointer underline'>Starter</div>
                </Link>
                <Link to='/plans/dev'>
                    <div className='mr-2  rounded-md py-2 px-6 text-lg font-bold cursor-pointer underline'>Dev</div>
                </Link>
                <Link to='/plans/pro'>
                    <div className='mr-2  rounded-md py-2 px-6 text-lg font-bold cursor-pointer underline'>Pro</div>
                </Link>

            </div>




        </div>
    )
}
