import './App.css'
import { lazy, Suspense } from 'react'
import { Routes, Route } from 'react-router-dom'
const Index = lazy(() => import('./Pages/index'))
function App() {
  return (
    <>
    <Suspense>
        <Routes>
            <Route path='/' element={<Index/>}/>
        </Routes>
    </Suspense>
    </>
  )
}

export default App
