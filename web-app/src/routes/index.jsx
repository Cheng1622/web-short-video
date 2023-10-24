import { Routes, Route } from 'react-router-dom'
import Home from '../pages/Home'


const RoutesConfig = () => (
    <Routes>
        <Route path='/' element={<Home />}></Route>
        <Route path='/home' element={<Home />}></Route>
    </Routes>
)

export default RoutesConfig