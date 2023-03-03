import React, { Component } from 'react';
import Board from './Board';
// import axios from "axios";

class App extends Component {
  constructor(){
    super()
    this.state = {
      boards: []
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
    let boards;
    if (typeof this.state.boards !== 'undefined') {
      boards = this.state.boards?.map(b => {
        return <Board data={b} />
      })
    }
    return (
      <div>
        <h1>works</h1>
        {boards}
      </div>
    );
  }
}

export default App;
