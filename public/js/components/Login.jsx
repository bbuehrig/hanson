import React, {Component} from 'react';

// Login-Form for User-Logins
class LoginForm extends Component{
  constructor(props) {
    super(props);
    this.state = {
      email: '',
      password: ''
    };

    this.onChangeEmail = this.onChangeEmail.bind(this);
    this.onChangePassword = this.onChangePassword.bind(this);
  }

  // Set states for email and password
  onChangeEmail(e) {
      this.setState({email: e.target.value});
  }

  onChangePassword(e) {
      this.setState({password: e.target.value});
  }

  // Submit Login-Form to server
  handleSubmit(e){
    e.preventDefault();

    var email = this.state.email;
    var password = this.state.password;

    $.ajax({
       url: "/login",
       dataType: 'json',
       type: 'POST',
       data: JSON.stringify({
          "email": email,
          "password": password,
          "username": ""
       }),
       success: function(data) {
         window.location.replace("/dashboard");
       }.bind(this),
       error: function(xhr, status, err) {
         console.error(this.props.url, status, err.toString());
         alert("ERROR");
       }.bind(this)
     });
  }

  // Render...
  render(){
    return(
      <form onSubmit={this.handleSubmit.bind(this)}>
          <div className="form-group">
              <label className="form-label" for="login-email">E-Mail</label>
              <input className="form-input" type="email" id="login-email" placeholder="E-Mail" value={this.state.email} onChange={this.onChangeEmail}/>
              <br />
              <label className="form-label" for="login-password">Password</label>
              <input className="form-input" type="password" id="login-password" placeholder="Password" value={this.state.password} onChange={this.onChangePassword}/>

          </div>

          <div className="form-group">
            <button className="btn btn-primary float-right" type="submit">Login</button>
          </div>
      </form>
    )
  }
}


// Login-Card
class Login extends Component{

  render(){
    return (
      <div className="container">
        <div className="columns">
          <div className="column col-xs-2"></div>

          <div className="column col-xs-8">
            <div className="card">
              <div className="card-header">
                <h4 className="card-title">Login</h4>
              </div>
              <div className="card-body">
                <LoginForm />
              </div>

              <div className="card-footer">Please login...</div>
            </div>
          </div>

          <div className="column col-xs-2"></div>
        </div>
      </div>
    )
  }
}

export default Login
