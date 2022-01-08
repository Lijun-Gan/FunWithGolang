
import Homepage from './components/Homepage';
import { Route, Routes } from 'react-router-dom';
import GoFiberAPI from './components/GoFiberAPI';
import GqlgenAPI from './components/GqlgenAPI';
import GraphQLAPI from './components/GraphQLAPI';

function App() {


  return (

    <div className="App">
      <Homepage /> 

      <Routes>
        <Route exact path="/goFiber" element={ <GoFiberAPI /> }/>
        <Route exact path="/gqlgen" element={ <GqlgenAPI />}/>
        <Route exact path="/graphql" element={ < GraphQLAPI /> }/>
        <Route exact path="/" element={ < Homepage /> } />
      </Routes>

    </div>
  );
}

export default App;
