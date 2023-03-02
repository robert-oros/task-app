import React, { Component } from 'react';
// import axios from "axios";

class App extends Component {
  constructor(){
    super()
    this.state = {
      boards: {}
    }
  }

  componentDidMount(){
    fetch("http://localhost:8081/get_boards", {
      method: 'GET'
    })
    .then(respose => respose.json())
    .then(data => {
      this.setState({ boards : data})
    })
  }
  

  render() {
    console.log(this.state.boards[0])
    return (
      <h1>works</h1>
    );
  }
}

export default App;
