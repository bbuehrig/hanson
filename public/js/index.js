import React from 'react';
import ReactDOM from 'react-dom';

import Navigation from './components/Navigation.jsx';
import Login from './components/Login.jsx';
import Dashboard from './components/Dashboard.jsx';


// Get which site to render?
var site = document.getElementById('root').getAttribute("data-site");

// Render sites...
switch(site) {
    // Login Page
    case 'login':
        var components =
          <div>
            <Navigation />
            <Login />
          </div>;

        break;


    // Dashboard
    case 'dashboard':
        var components =
          <div>
            <Navigation />
            <Dashboard />
          </div>;

        break;
}


ReactDOM.render(
    components,
    document.getElementById('root')
  )
