import React from 'react'
import { Link } from 'react-router-dom'

export const NotFound = () => {
    return (
        <div
            className='bg-zinc-900 w-full h-screen flex justify-center items-center'
        >
            <div className='flex flex-col items-center'>
                <div className='text-creo font-bold text-9xl'>404</div>
                <Link to="/">
                    <div className='text-lg underline underline-offset-4 text-gray-200 font-bold'>
                        Go Home
                    </div>
                </Link>
            </div>
        </div>
    )
}
