import React, { Component } from 'react';
import axios from "axios";

class App extends Component {
  constructor(){
    super()
    this.state = {
      boards: {}
    }
  }

  componentDidMount(){
    axios.get("http://localhost:8081/get_boards")
    .then(res => {
      console.log("am intrat")
      console.log(res)
      this.setState({boards: res})
    })
    .then(res => {
      console.log(res)
    })
  }
  

  render() {
    return (
      <h1>works</h1>
    );
  }
}

export default App;
