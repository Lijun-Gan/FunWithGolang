import React from 'react';
import  { Link } from 'react-router-dom';


function Homepage() {


  return (
    <div className="Homepage">
      <h1>Welcome to a messey WEB!</h1>

      <nav className="displayFlex">
          <Link className="linkBoder" to="/goFiber"> Go with Fiber</Link>
          <Link className="linkBoder" to="/gqlgen"> Go with GraphQL</Link>
          <Link className="linkBoder" to="/graphql"> GraphQL</Link>
      </nav>
    </div>
  );
}

export default Homepage;
