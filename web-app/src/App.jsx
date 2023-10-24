import React, { Suspense } from 'react';
import './App.css';
import RoutesConfig from './routes/index'

function App() {
  return (
    <div className="App">
       <Suspense fallback={<div>loading...</div>}>
        <RoutesConfig />
      </Suspense>

    </div>
  );
}

export default App;
