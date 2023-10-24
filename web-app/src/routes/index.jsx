import { lazy } from 'react'
import { Routes, Route } from 'react-router-dom'
import Home from '../pages/Home'

const RoutesConfig=()=>(
    <Routes>
        <Route path='/' element={<Home/>}></Route>
    </Routes>
)

export default RoutesConfig